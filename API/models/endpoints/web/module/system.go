package module

import (
	"accountBook/models/beans"
	"accountBook/models/beans/customer"
	"accountBook/models/beans/dbBeans"
	"accountBook/models/endpoints"
	"accountBook/models/endpoints/web"
	"accountBook/models/service"

	"github.com/kinwyb/go/db"
	"github.com/kinwyb/go/err1"
)

var System web.ISystemEp = &systemEp{}

type systemEp struct{}

func (systemEp) LogList(req *customer.LogListReq, pg *db.PageObj, ctx *beans.Context) ([]*dbBeans.Log, err1.Error) {
	defer ctx.Start("ep.system.LogList").Finish()
	if err := endpoints.CheckPower("system.LogList", ctx.Child()); err != nil {
		return nil, err
	}
	return service.LogList(req, pg, ctx.Child()), nil
}
