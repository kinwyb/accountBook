package service

import (
	"accountBook/models/beans"
	"accountBook/models/beans/customer"
	"accountBook/models/beans/dbBeans"
	"accountBook/models/dataBase"

	"github.com/kinwyb/go/db"
	"github.com/kinwyb/go/err1"
)

// 新增日志
func LogAdd(req *dbBeans.LogDB, ctx *beans.Context) err1.Error {
	return dataBase.LogAdd(req, ctx)
}

// 日志列表
func LogList(req *customer.LogListReq, pg *db.PageObj, ctx *beans.Context) []*dbBeans.Log {
	if req.StartTime != "" {
		req.StartTime = req.StartTime + " 00:00:00"
	}
	if req.EndTime != "" {
		req.EndTime = req.EndTime + " 23:59:59"
	}
	return dataBase.LogList(req, pg, ctx)
}
