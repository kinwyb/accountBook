package web

import (
	"accountBook/application/web/controllers"
	"accountBook/models/beans/dbBeans"
	"accountBook/models/endpoints/web"
	"encoding/json"
)

// 收支类型相关接口
type ReceiptTypeController struct {
	controllers.RestController
	Serv web.IReceiptTypeEndpoint
}

// @Title 收支类型列表
// @Description 收支类型列表
// @Param token header string true Token
// @Param parentID query string false 父级ID
// @Success 200 {array} dbBeans.ReceiptType
// @router /list [get]
func (b *ReceiptTypeController) List() {
	parent, _ := b.GetInt64("parentID", -1)
	ret, err := b.Serv.List(parent, b.OCtx)
	if err != nil {
		b.RespError(err)
		return
	}
	b.ResponseSUCC(ret)
}

// @Title 收支类型列表
// @Description 收支类型列表
// @Param token header string true Token
// @Param level query int64 false 等级
// @Success 200 {array} dbBeans.ReceiptType
// @router /list/level [get]
func (b *ReceiptTypeController) ListByLevel() {
	level, _ := b.GetInt64("level", -1)
	ret, err := b.Serv.ListByLevel(level, b.OCtx)
	if err != nil {
		b.RespError(err)
		return
	}
	b.ResponseSUCC(ret)
}

// @Title 收支类型树形结构
// @Description 收支类型树形结构
// @Param token header string true Token
// @Success 200 {array} customer.ReceiptTypeTree
// @router /tree [get]
func (b *ReceiptTypeController) Tree() {
	ret, err := b.Serv.Tree(b.OCtx)
	if err != nil {
		b.RespError(err)
		return
	}
	b.ResponseSUCC(ret)
}

// @Title 新增
func (b *ReceiptTypeController) Add() {
	var req dbBeans.ReceiptTypeDB
	e := json.Unmarshal(b.Ctx.Input.RequestBody, &req)
	if e != nil {
		b.LogError("dbBeans.ReceiptTypeDB请求内容:%s\n请求类型:%s",
			b.Ctx.Input.RequestBody, b.Ctx.Request.Header.Get("content-type"))
		b.RespError(controllers.ParamDecodeFail, e)
		return
	}
	err := b.Serv.Add(&req, b.OCtx)
	if err != nil {
		b.RespError(err)
		return
	}
	b.ResponseSUCC(controllers.Success)
}
