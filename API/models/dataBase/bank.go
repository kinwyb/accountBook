package dataBase

import (
	"accountBook/models/beans"
	"accountBook/models/beans/dbBeans"
)

// 银行列表
func BankList(ctx *beans.Context) []*dbBeans.Bank {
	defer ctx.Start("db.BankList").Finish()
	return dbBeans.BankGetList("", ctx.Query)
}
