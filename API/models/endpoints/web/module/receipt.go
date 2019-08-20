package module

import (
	"accountBook/models/beans"
	"accountBook/models/beans/customer"
	"accountBook/models/beans/dbBeans"
	"accountBook/models/endpoints"
	"accountBook/models/endpoints/web"
	"accountBook/models/service"
	"strings"

	"github.com/kinwyb/go/db"
	"github.com/kinwyb/go/err1"
)

func init() {
	endpoints.RegisterLogContentParseFun("Receipt", logContentParse)
}

func logContentParse(db *customer.LogAddReq) {
	switch db.CallFunc {
	case "Receipt.Add":
		if len(db.Args) > 0 {
			arg := db.Args[0].(*dbBeans.ReceiptDB)
			if arg.Money > 0 {
				db.Content = strings.Split(db.Content, "【")[0] + "【收入单据】"
			} else {
				db.Content = strings.Split(db.Content, "【")[0] + "【支出单据】"
			}
		}
	}
}

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
	ctx.RequestArgs = []interface{}{req} //请求参数
	if err := endpoints.CheckPower("Receipt.Add", ctx.Child()); err != nil {
		return err
	}
	return service.ReceiptAdd(req, ctx.Child())
}
