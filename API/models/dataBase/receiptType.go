package dataBase

import (
	"accountBook/models/beans"
	"accountBook/models/beans/dbBeans"
)

// 收支类型

// 收支类型列表
func ReceiptTypeList(parentID int64, ctx *beans.Context) []*dbBeans.ReceiptType {
	defer ctx.Start("db.ReceiptTypeList").Finish()
	return dbBeans.ReceiptTypeGetList(" parent_id = ? ", ctx.Query, parentID)
}

// 查询所有收支类型
func ReceiptTypeListAll(ctx *beans.Context) []*dbBeans.ReceiptType {
	defer ctx.Start("db.ReceiptTypeListAll").Finish()
	return dbBeans.ReceiptTypeGetList("", ctx.Query)
}
