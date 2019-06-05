package web

import (
	"accountBook/application/web/controllers"
	"accountBook/models/beans/customer"
	"accountBook/models/beans/dbBeans"
	"accountBook/models/endpoints/web"
	"encoding/json"

	"github.com/kinwyb/go/db"
)

// 收支信息相关接口
type ReceiptController struct {
	controllers.RestController
	Serv web.IReceiptEndpoint
}

// @Title 收支列表
// @Description 收支列表
// @Param token header string true Token
// @Param req body customer.ReceiptListReq true 请求参数
// @Param page query int false 当前页数
// @Param pageSize query int false 每页条数
// @Success 200 {object} customer.ReceiptListResp
// @router /list [post]
func (i *ReceiptController) List() {
	var req customer.ReceiptListReq
	e := json.Unmarshal(i.Ctx.Input.RequestBody, &req)
	if e != nil {
		i.LogError("customer.ReceiptListReq请求内容:%s\n请求类型:%s", i.Ctx.Input.RequestBody, i.Ctx.Request.Header.Get("content-type"))
		i.RespError(controllers.ParamDecodeFail, e)
		return
	}
	page, _ := i.GetInt("page", 1)
	pageSize, _ := i.GetInt("pageSize", 20)
	pg := &db.PageObj{Page: page, Rows: pageSize}
	ret, err := i.Serv.List(&req, pg, i.OCtx)
	if err != nil {
		i.RespError(err)
		return
	}
	i.Page(pg)
	i.ResponseSUCC(ret)
}

// @Title 下一个单据号
// @router /nextNo [get]
func (i *ReceiptController) NextNo() {
	ret := i.Serv.NextNo(i.OCtx)
	i.ResponseSUCC(ret)
}

// @Title 新增收支单据
func (i *ReceiptController) Add() {
	var req dbBeans.ReceiptDB
	e := json.Unmarshal(i.Ctx.Input.RequestBody, &req)
	if e != nil {
		i.LogError("dbBeans.ReceiptDB请求内容:%s\n请求类型:%s", i.Ctx.Input.RequestBody, i.Ctx.Request.Header.Get("content-type"))
		i.RespError(controllers.ParamDecodeFail, e)
		return
	}
	err := i.Serv.Add(&req, i.OCtx)
	if err != nil {
		i.RespError(err)
		return
	}
	i.ResponseSUCC(controllers.Success)
}
