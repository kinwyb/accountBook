package module

import (
	"accountBook/models/beans"
	"accountBook/models/beans/customer"
	"accountBook/models/endpoints"
	"accountBook/models/endpoints/web"
	"accountBook/models/service"

	"github.com/kinwyb/go/db"
	"github.com/kinwyb/go/err1"
)

var Receipt web.IReceiptEndpoint = &receiptEp{}

type receiptEp struct{}

func (receiptEp) List(req *customer.ReceiptListReq, pg *db.PageObj, ctx *beans.Context) (*customer.ReceiptListResp, err1.Error) {
	defer ctx.Start("ep.ReceiptList").Finish()
	if err := endpoints.CheckPower("ReceiptList", ctx.Child()); err != nil {
		return nil, err
	}
	return service.ReceiptList(req, pg, ctx.Child()), nil
}
