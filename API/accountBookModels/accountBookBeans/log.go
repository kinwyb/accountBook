package accountBookBeans

import (
	"code.aliyun.com/zhizaofang/zfgoutil"

	"github.com/kinwyb/go/logs"
)

const LogName = "service.log"

var log = zfgoutil.GetLogger(LogName)

func init() {
	zfgoutil.RegisterLog(setLog)
}

func setLog(lg *logs.LogFiles) {
	log = lg.GetLog(LogName)
}
