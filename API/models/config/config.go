package config

import (
	"github.com/astaxie/beego/config"
	"github.com/kinwyb/go/db/mysql"
)

func InitConfig(conf config.Configer) {
	//数据库
	initDB(conf)
	//todo: 加载其他初始化配置
}

//数据库
func initDB(conf config.Configer) {
	//设置数据库配置信息
	mysql.InitializeDB(
		conf.DefaultString("mysql.host", "127.0.0.1:3306"),
		conf.DefaultString("mysql.username", "root"),
		conf.DefaultString("mysql.password", ""),
		conf.DefaultString("mysql.name", ""),
	)
}
