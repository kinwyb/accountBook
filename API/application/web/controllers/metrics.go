package controllers

import (
	"accountBook/accountBookModels/accountBookBeans"
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/kinwyb/go"
	"github.com/shirou/gopsutil/process"

	"github.com/astaxie/beego"
	"github.com/rcrowley/go-metrics/exp"

	"github.com/rcrowley/go-metrics"
	"github.com/vrischmann/go-metrics-influxdb"
)

var enableMetrics = false //是否开启性能统计
var metricsRegistry metrics.Registry
var allRequestMetrics = "all.request"
var metricsCancel context.CancelFunc

//开启性能统计
func StartMetrics() {
	if accountBookBeans.MetricsInfluxdb == "" ||
		accountBookBeans.MachineNo == "" { //时序数据库地址错误不启动
		return
	}
	enableMetrics = true
	metricsRegistry = metrics.NewRegistry()
	beego.Handler("/debug/metrics", exp.ExpHandler(metricsRegistry))
	beego.Handler(heldiamgo.PprofHttpHandler())
	metrics.RegisterDebugGCStats(metricsRegistry)
	metrics.RegisterRuntimeMemStats(metricsRegistry)
	go metrics.CaptureDebugGCStats(metricsRegistry, time.Second*5)
	go metrics.CaptureRuntimeMemStats(metricsRegistry, time.Second*5)
	go influxdb.InfluxDB(metricsRegistry, time.Second, accountBookBeans.MetricsInfluxdb,
		"factoryshop_"+accountBookBeans.MachineNo+
			"_"+accountBookBeans.RunMode, "", "")
	ctx := context.Background()
	ctx, metricsCancel = context.WithCancel(ctx)
	go runInfo(metricsRegistry, time.Second*5, ctx) //运行时执行的数据
}

//运行时的数据
func runInfo(r metrics.Registry, d time.Duration, ctx context.Context) {
	ps, err := process.Processes()
	if err != nil {
		log.Error("运行数据获取失败:%s", err.Error())
		return
	}
	var proc *process.Process
	for _, v := range ps {
		if cmd, err := v.Cmdline(); err == nil {
			if cmd == os.Args[0] {
				proc = v
				break
			}
		}
	}
	if proc == nil {
		log.Error("运行进程ID获取失败")
		return
	}
	log.Info("获取到运行进程ID:%d", proc.Pid)
	memRSS := metrics.NewGauge()
	memVMS := metrics.NewGauge()
	r.Register("process.Mem.RSS", memRSS)
	r.Register("process.Mem.VMS", memVMS)
	infs, err := net.Interfaces()
	if err != nil {
		log.Error("系统网络获取失败")
	}
	netRecvGauge := map[string]metrics.Gauge{}
	netSentGauge := map[string]metrics.Gauge{}
	netLastRecvBytes := map[string]uint64{}
	netLastSentBytes := map[string]uint64{}
	for _, v := range infs {
		addr, _ := v.Addrs()
		if len(addr) < 1 {
			continue
		}
		addrString := addr[0].String()
		netRecvGauge[v.Name] = metrics.NewGauge()
		netSentGauge[v.Name] = metrics.NewGauge()
		r.Register(fmt.Sprintf("process.net.recv.%s", addrString), netRecvGauge[v.Name])
		r.Register(fmt.Sprintf("process.net.sent.%s", addrString), netSentGauge[v.Name])
	}
	dBits := uint64(d / time.Second)
	for {
		select {
		case <-time.Tick(d):
			//内存
			mem, err := proc.MemoryInfo()
			if err != nil {
				log.Error("进程运行内存获取失败:%s", err.Error())
			} else {
				memRSS.Update(int64(mem.RSS))
				memVMS.Update(int64(mem.VMS))
			}
			//网络
			netInfo, err := proc.NetIOCounters(true)
			if err == nil {
				for _, v := range netInfo {
					if gauge, ok := netRecvGauge[v.Name]; ok {
						recv := netLastRecvBytes[v.Name]
						if recv < 1 {
							netLastRecvBytes[v.Name] = v.BytesRecv
						} else {
							bits := v.BytesRecv - recv
							bits = bits / dBits
							gauge.Update(int64(bits))
							netLastRecvBytes[v.Name] = v.BytesRecv
						}
					}
					if gauge, ok := netSentGauge[v.Name]; ok {
						sent := netLastSentBytes[v.Name]
						if sent < 1 {
							netLastSentBytes[v.Name] = v.BytesSent
						} else {
							bits := v.BytesSent - sent
							bits = bits / dBits
							gauge.Update(int64(bits))
							netLastSentBytes[v.Name] = v.BytesSent
						}
					}
				}
			}
		case <-ctx.Done():
			goto end
		}
	}
end:
	return
}

//结束性能统计
func StopMetrics() {
	enableMetrics = false
	metricsRegistry.UnregisterAll()
	metricsRegistry = nil
	metricsCancel() //关闭
}
