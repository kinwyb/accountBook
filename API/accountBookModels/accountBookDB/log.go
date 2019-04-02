package accountBookDB

import (
	"accountBook/accountBookModels/accountBookBeans"

	"code.aliyun.com/zhizaofang/zfgoutil"

	"github.com/kinwyb/go/logs"
)

var log = zfgoutil.GetLogger(accountBookBeans.LogName)

func init() {
	zfgoutil.RegisterLog(setLog)
}

func setLog(lg *logs.LogFiles) {
	log = lg.GetLog(accountBookBeans.LogName)
}
