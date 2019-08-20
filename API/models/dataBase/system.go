package dataBase

import (
	"accountBook/models/beans"
	"accountBook/models/beans/customer"
	"accountBook/models/beans/dbBeans"
	"strings"

	"github.com/kinwyb/go/db"

	heldiamgo "github.com/kinwyb/go"

	"github.com/kinwyb/go/err1"
)

// 新增日志
func LogAdd(req *dbBeans.LogDB, ctx *beans.Context) err1.Error {
	defer ctx.Start("db.LogAdd").Finish()
	req.InstallTime = heldiamgo.TimeNow()
	return Insert(req, dbBeans.TableLog, ctx.Query)
}

// 日志列表
func LogList(req *customer.LogListReq, pg *db.PageObj, ctx *beans.Context) []*dbBeans.Log {
	defer ctx.Start("db.LogList").Finish()
	whereSQL := strings.Builder{}
	whereSQL.WriteString(" 1 = 1")
	var args []interface{}
	if req.StartTime != "" {
		whereSQL.WriteString(" AND install_time >= ? ")
		args = append(args, req.StartTime)
	}
	if req.EndTime != "" {
		whereSQL.WriteString(" AND install_time <= ? ")
		args = append(args, req.EndTime)
	}
	whereSQL.WriteString(" ORDER BY install_time DESC ")
	return dbBeans.LogGetPageList(whereSQL.String(), ctx.Query, pg, args...)
}
