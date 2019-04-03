package web

import (
	"accountBook/application/web/controllers"
	"accountBook/models/endpoints/web"
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
