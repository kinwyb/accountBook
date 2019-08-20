package web

import (
	"accountBook/models/beans"
	"accountBook/models/beans/customer"
	"accountBook/models/beans/dbBeans"

	"github.com/kinwyb/go/err1"

	"github.com/kinwyb/go/db"
)

type ISystemEp interface {

	// 日志列表
	LogList(req *customer.LogListReq, pg *db.PageObj, ctx *beans.Context) ([]*dbBeans.Log, err1.Error)
}
