package module

import (
	"accountBook/models/beans"
	"accountBook/models/beans/dbBeans"
	"accountBook/models/endpoints"
	"accountBook/models/endpoints/web"
	"accountBook/models/service"

	"github.com/kinwyb/go/err1"
)

var ReceiptType web.IReceiptTypeEndpoint = &receiptTypeEp{}

type receiptTypeEp struct{}

func (receiptTypeEp) ListByLevel(level int64, ctx *beans.Context) ([]*dbBeans.ReceiptType, err1.Error) {
	defer ctx.Start("ep.ReceiptTypeListByLevel").Finish()
	if err := endpoints.CheckPower("ReceiptTypeListByLevel", ctx.Child()); err != nil {
		return nil, err
	}
	return service.ReceiptListByLevel(level, ctx.Child()), nil
}

func (receiptTypeEp) List(parentID int64, ctx *beans.Context) ([]*dbBeans.ReceiptType, err1.Error) {
	defer ctx.Start("ep.ReceiptTypeList").Finish()
	if err := endpoints.CheckPower("ReceiptTypeList", ctx.Child()); err != nil {
		return nil, err
	}
	return service.ReceiptTypeList(parentID, ctx.Child()), nil
}
