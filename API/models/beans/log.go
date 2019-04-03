package beans

import (
	"time"

	"github.com/kinwyb/go/logs"
)

const LogName = "service.log"

//注册一个日志获取函数
type RegisterLogFunc func(log *logs.LogFiles)

var logFactory = logs.NewLogFiles("", 24*time.Hour)

var log = GetLogger(LogName)

var logmap []RegisterLogFunc

//设置日志路径
func SetLogPath(filepath string, level ...logs.Level) {
	if filepath == "" {
		return
	} else if len(level) < 1 {
		level = []logs.Level{logs.Debug}
	}
	logFactory = logs.NewLogFiles(filepath, 24*time.Hour, level[0])
	log = logFactory.GetLog(LogName)
	for _, v := range logmap {
		if v != nil {
			v(logFactory)
		}
	}
}

func RegisterLog(fun RegisterLogFunc) {
	if fun != nil {
		logmap = append(logmap, fun)
	}
}

//获取一个日志
func GetLogger(logname string) logs.Logger {
	return logFactory.GetLog(logname)
}
