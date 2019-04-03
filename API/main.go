package main

import (
	"accountBook/application/web/controllers"
	_ "accountBook/application/web/routers"
	"accountBook/models/beans"
	"accountBook/models/config"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	beego.LoadAppConfig("ini", "app.conf")
	beego.BConfig.WebConfig.DirectoryIndex = true
	// beego.BConfig.WebConfig.StaticDir["/swagger"] = filepath.Join(beego.AppPath, "swagger")
	//beego.BeeLogger.SetLevel(beego.LevelWarning)
	if beego.BConfig.RunMode != "dev" {
		logpath := beego.AppConfig.DefaultString("log.path", "")
		//设置日志
		beans.SetLogPath(logpath)
		beego.BeeLogger.Reset()
		beego.BConfig.Log.AccessLogs = true
		//beego.BConfig.Log.AccessLogsFormat = "JSON_FORMAT"
		beego.BeeLogger.Async(1000)
		beego.BeeLogger.SetLogger(logs.AdapterFile,
			"{\"filename\":\""+logpath+"/request.log\",\"level\":7,\"maxlines\":0,\"maxsize\":0,\"daily\":true,\"maxdays\":10}")
		beego.BeeLogger.EnableFuncCallDepth(true)
	}
	config.InitConfig(beego.AppConfig)
	controllers.StartMetrics()
	defer controllers.StopMetrics()
	beego.Run()
}
