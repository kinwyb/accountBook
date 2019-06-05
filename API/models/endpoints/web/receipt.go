package web

import (
	"accountBook/models/beans"
	"accountBook/models/beans/customer"
	"accountBook/models/beans/dbBeans"

	"github.com/kinwyb/go/db"
	"github.com/kinwyb/go/err1"
)

type IReceiptEndpoint interface {
	// @Title 收支列表
	// @Description 收支列表
	// @Param token header string true Token
	// @Param req body customer.ReceiptListReq true 请求参数
	// @Param page query int false 当前页数
	// @Param pageSize query int false 每页条数
	// @Success 200 {object} customer.ReceiptListResp
	// @router /list [post]
	List(req *customer.ReceiptListReq, pg *db.PageObj, ctx *beans.Context) (*customer.ReceiptListResp, err1.Error)

	// @Title 下一个单据
	NextNo(ctx *beans.Context) string

	// 新增收支单据
	Add(req *dbBeans.ReceiptDB, ctx *beans.Context) err1.Error
}
