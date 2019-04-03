package service

import (
	"accountBook/models/beans"
	"accountBook/models/beans/dbBeans"
	"accountBook/models/dataBase"
)

// 银行列表
func BankList(ctx *beans.Context) []*dbBeans.Bank {
	defer ctx.Start("sev.BankList").Finish()
	return dataBase.BankList(ctx.Child())
}
