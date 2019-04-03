package module

import (
	"accountBook/models/beans"
	"accountBook/models/beans/dbBeans"
	"accountBook/models/endpoints"
	"accountBook/models/endpoints/web"
	"accountBook/models/service"

	"github.com/kinwyb/go/err1"
)

var Bank web.IBankEndpoint = &bankEp{}

type bankEp struct{}

func (bankEp) List(ctx *beans.Context) ([]*dbBeans.Bank, err1.Error) {
	defer ctx.Start("ep.BankList").Finish()
	if err := endpoints.CheckPower("BankList", ctx.Child()); err != nil {
		return nil, err
	}
	return service.BankList(ctx.Child()), nil
}
