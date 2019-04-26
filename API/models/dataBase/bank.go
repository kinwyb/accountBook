package dataBase

import (
	"accountBook/models/beans"
	"accountBook/models/beans/dbBeans"

	"github.com/kinwyb/go/err1"
)

// 银行列表
func BankList(ctx *beans.Context) []*dbBeans.Bank {
	defer ctx.Start("db.BankList").Finish()
	return dbBeans.BankGetList("", ctx.Query)
}

// 新增银行
func BankAdd(req *dbBeans.BankDB, ctx *beans.Context) err1.Error {
	defer ctx.Start("db.BankAdd").Finish()
	return Insert(req, dbBeans.TableBank, ctx.Query)
}
