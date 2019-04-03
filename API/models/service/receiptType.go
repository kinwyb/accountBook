package service

import (
	"accountBook/models/beans"
	"accountBook/models/beans/dbBeans"
	"accountBook/models/dataBase"
)

// 收支类型

// 收支类型列表
func ReceiptTypeList(parentID int64, ctx *beans.Context) []*dbBeans.ReceiptType {
	defer ctx.Start("sev.ReceiptTypeList").Finish()
	if parentID < 1 {
		parentID = 0
	}
	return dataBase.ReceiptTypeList(parentID, ctx.Child())
}
