package service

import (
	"accountBook/models/beans"
	"accountBook/models/beans/dbBeans"
	"accountBook/models/dataBase"

	"github.com/kinwyb/go/err1"
)

// 银行列表
func BankList(ctx *beans.Context) []*dbBeans.Bank {
	defer ctx.Start("sev.BankList").Finish()
	return dataBase.BankList(ctx.Child())
}

// 新增银行
func BankAdd(req *dbBeans.BankDB, ctx *beans.Context) err1.Error {
	defer ctx.Start("sev.BankAdd").Finish()
	return dataBase.BankAdd(req, ctx.Child())
}
