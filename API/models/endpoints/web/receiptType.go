package web

import (
	"accountBook/models/beans"
	"accountBook/models/beans/customer"
	"accountBook/models/beans/dbBeans"

	"github.com/kinwyb/go/err1"
)

// 收支类型
type IReceiptTypeEndpoint interface {
	// @Title 收支类型列表
	// @Description 收支类型列表
	// @Param token header string true Token
	// @Param parentID query string false 父级ID
	// @Success 200 {array} dbBeans.ReceiptType
	// @router /list [get]
	List(parentID int64, ctx *beans.Context) ([]*dbBeans.ReceiptType, err1.Error)

	// @Title 收支类型列表
	// @Description 收支类型列表
	// @Param token header string true Token
	// @Param level query int64 false 等级
	// @Success 200 {array} dbBeans.ReceiptType
	// @router /list/level [get]
	ListByLevel(level int64, ctx *beans.Context) ([]*dbBeans.ReceiptType, err1.Error)

	// @Title 收支类型树形结构
	// @Description 收支类型树形结构
	// @Param token header string true Token
	// @Success 200 {array} customer.ReceiptTypeTree
	// @router /tree [get]
	Tree(ctx *beans.Context) ([]*customer.ReceiptTypeTree, err1.Error)
}
