package web

import (
	"accountBook/application/web/controllers"
	"accountBook/models/beans/dbBeans"
	"accountBook/models/endpoints/web"
	"encoding/json"
)

// 银行相关接口
type BankController struct {
	controllers.RestController
	Serv web.IBankEndpoint
}

// @Title 银行列表
// @Description 银行列表
// @Param token header string true Token
// @Success 200 {array} dbBeans.Bank
// @router /list [get]
func (b *BankController) List() {
	ret, err := b.Serv.List(b.OCtx)
	if err != nil {
		b.RespError(err)
		return
	}
	b.ResponseSUCC(ret)
}

// @Title 银行按日期进行汇总计算
// @Description 银行按日期进行汇总计算
// @Param token header string true Token
// @Param startTime query string false 起始时间
// @Param endTime query string false 结束时间
// @Success 200 {array} customer.BankListCompateResp
// @router /list/compute/day [get]
func (b *BankController) ListComputeWithDay() {
	startTime := b.GetString("startTime")
	endTime := b.GetString("endTime")
	ret, err := b.Serv.ListComputeWithDay(startTime, endTime, b.OCtx)
	if err != nil {
		b.RespError(err)
		return
	}
	b.ResponseSUCC(ret)
}

// @Title 新增银行信息
// @Description 新增银行信息
// @Param token header string true Token
// @Param req body dbBeans.BankDB true 参数详情
// @router /add [post]
func (b *BankController) Add() {
	var req dbBeans.BankDB
	if err := json.Unmarshal(b.Ctx.Input.RequestBody, &req); err != nil {
		b.ResponseError(-1, err.Error())
		return
	}
	err := b.Serv.Add(&req, b.OCtx)
	if err != nil {
		b.RespError(err)
		return
	}
	b.ResponseSUCC(controllers.Success)
}
