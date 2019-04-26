package web

import (
	"accountBook/models/beans"
	"accountBook/models/beans/customer"
	"accountBook/models/beans/dbBeans"

	"github.com/kinwyb/go/err1"
)

type IBankEndpoint interface {
	// @Title 银行列表
	// @Description 银行列表
	// @Param token header string true Token
	// @Success 200 {array} dbBeans.Bank
	// @router /list [get]
	List(ctx *beans.Context) ([]*dbBeans.Bank, err1.Error)

	// @Title 银行按日期进行汇总计算
	// @Description 银行按日期进行汇总计算
	// @Param token header string true Token
	// @Param startTime query string false 起始时间
	// @Param endTime query string false 结束时间
	// @Success 200 {array} customer.BankListCompateResp
	// @router /list/compute/day [get]
	ListComputeWithDay(startTime string, endTime string, ctx *beans.Context) ([]*customer.BankListCompateResp, err1.Error)

	// @Title 新增银行信息
	// @Description 新增银行信息
	// @Param token header string true Token
	// @Param req body dbBeans.BankDB true 参数详情
	// @router /add [post]
	Add(req *dbBeans.BankDB, ctx *beans.Context) err1.Error
}
