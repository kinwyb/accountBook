package service

import (
	"accountBook/models/beans"
	"accountBook/models/beans/customer"
	"accountBook/models/beans/dbBeans"
	"accountBook/models/dataBase"

	"github.com/kinwyb/go/err1"
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

// 收支类型列表按等级
func ReceiptListByLevel(level int64, ctx *beans.Context) []*dbBeans.ReceiptType {
	defer ctx.Start("sev.ReceiptListByLevel").Finish()
	if level < 1 {
		level = 0
	}
	return dataBase.ReceiptTypeListByLevel(level, ctx.Child())
}

// 收支类型树形结构
func ReceiptTypeTree(ctx *beans.Context) []*customer.ReceiptTypeTree {
	defer ctx.Start("sev.ReceiptTypeTree").Finish()
	all := dataBase.ReceiptTypeListAll(ctx.Child())
	receiptTreeMap := map[int64]*customer.ReceiptTypeTree{}
	var ret []*customer.ReceiptTypeTree
	for _, v := range all {
		if x, ok := receiptTreeMap[v.Id]; ok {
			x.ID = v.Id
			x.Name = v.Name
		} else {
			r := &customer.ReceiptTypeTree{
				ID:   v.Id,
				Name: v.Name,
			}
			receiptTreeMap[v.Id] = r
		}
		if v.ParentId < 1 {
			ret = append(ret, receiptTreeMap[v.Id])
		}
		if v.ParentId > 0 {
			if x, ok := receiptTreeMap[v.ParentId]; ok {
				x.Children = append(x.Children, receiptTreeMap[v.Id])
			} else {
				receiptTreeMap[v.ParentId] = &customer.ReceiptTypeTree{
					Children: []*customer.ReceiptTypeTree{
						receiptTreeMap[v.Id],
					},
				}
			}
		}
	}
	return ret
}

// 收支类型列表
func ReceiptTypeAdd(req *dbBeans.ReceiptTypeDB, ctx *beans.Context) err1.Error {
	defer ctx.Start("sev.ReceiptTypeAdd").Finish()
	return dataBase.ReceiptTypeAdd(req, ctx.Child())
}
