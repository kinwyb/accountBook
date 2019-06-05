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

var Receipt web.IReceiptEndpoint = &receiptEp{}

type receiptEp struct{}

func (receiptEp) List(req *customer.ReceiptListReq, pg *db.PageObj, ctx *beans.Context) (*customer.ReceiptListResp, err1.Error) {
	defer ctx.Start("ep.ReceiptList").Finish()
	if err := endpoints.CheckPower("ReceiptList", ctx.Child()); err != nil {
		return nil, err
	}
	return service.ReceiptList(req, pg, ctx.Child()), nil
}

func (receiptEp) NextNo(ctx *beans.Context) string {
	defer ctx.Start("ep.ReceiptNextNo").Finish()
	if err := endpoints.CheckPower("ReceiptNextNo", ctx.Child()); err != nil {
		return ""
	}
	return service.ReceiptNextNo(ctx.Child())
}

func (receiptEp) Add(req *dbBeans.ReceiptDB, ctx *beans.Context) err1.Error {
	defer ctx.Start("ep.ReceiptAdd").Finish()
	if err := endpoints.CheckPower("ReceiptAdd", ctx.Child()); err != nil {
		return err
	}
	return service.ReceiptAdd(req, ctx.Child())
}
