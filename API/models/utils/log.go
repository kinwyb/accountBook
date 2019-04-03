package utils

import (
	"accountBook/models/beans"

	"github.com/kinwyb/go/logs"
)

var log = beans.GetLogger(beans.LogName)

func init() {
	beans.RegisterLog(setLog)
}

func setLog(lg *logs.LogFiles) {
	log = lg.GetLog(beans.LogName)
}
