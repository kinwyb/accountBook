package log

import (
	"github.com/kinwyb/go/logs"
)

//设置日志路径
func SetLogPath(filepath string, level ...logs.Level) {
	if filepath == "" {
		return
	} else if len(level) < 1 {
		level = []logs.Level{logs.Debug}
	}
	logs.SetLogPath(filepath, level...)
}

func GetLog(tag Tag) logs.Logger {
	return logs.GetLogger(tag.Value())
}

//调试
func Debug(tag Tag, format string, args ...interface{}) {
	logs.GetLogger(tag.Value()).Debug(format, args...)
}

//输出
func Info(tag Tag, format string, args ...interface{}) {
	logs.GetLogger(tag.Value()).Info(format, args...)
}

//警告
func Warning(tag Tag, format string, args ...interface{}) {
	logs.GetLogger(tag.Value()).Warning(format, args...)
}

//错误
func Error(tag Tag, format string, args ...interface{}) {
	logs.GetLogger(tag.Value()).Error(format, args...)
}

//关键
func Critical(tag Tag, format string, args ...interface{}) {
	logs.GetLogger(tag.Value()).Critical(format, args...)
}

//警报
func Alert(tag Tag, format string, args ...interface{}) {
	logs.GetLogger(tag.Value()).Alert(format, args...)
}

//紧急
func Emergency(tag Tag, format string, args ...interface{}) {
	logs.GetLogger(tag.Value()).Emergency(format, args...)
}
