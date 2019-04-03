package web

import (
	"accountBook/models/beans"
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
}
