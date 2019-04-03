package service

import (
	"accountBook/models/beans"
	"accountBook/models/beans/customer"
	"accountBook/models/beans/dbBeans"
	"accountBook/models/dataBase"
	"fmt"

	"github.com/kinwyb/go/db"
)

// 收支列表
func ReceiptList(req *customer.ReceiptListReq, pg *db.PageObj, ctx *beans.Context) *customer.ReceiptListResp {
	defer ctx.Start("sev.ReceiptList").Finish()
	ret := &customer.ReceiptListResp{}
	result := dataBase.ReceiptList(req, pg, ctx.Child())
	bankList := dataBase.BankList(ctx.Child())
	bankMap := map[int64]string{}
	for _, v := range bankList {
		bankMap[v.Id] = v.BankName
	}
	receiptTypeList := dataBase.ReceiptTypeListAll(ctx.Child())
	receiptTypeMap := map[int64]*dbBeans.ReceiptType{}
	for _, v := range receiptTypeList {
		receiptTypeMap[v.Id] = v
	}
	for _, v := range result {
		tp := receiptTypeMap[v.Type]
		tpValue := ""
		if tp != nil {
			tpValue = tp.Name
			if tp.ParentId != 0 {
				tpParent := receiptTypeMap[tp.ParentId]
				if tpParent != nil {
					tpValue = fmt.Sprintf("%s/%s", tpParent.Name, tp.Name)
				}
			}
		}
		ret.Data = append(ret.Data, &customer.Receipt{
			Id:          fmt.Sprintf("S%09d", v.Id),
			Money:       v.Money,
			Bank:        bankMap[v.BankId],
			Description: v.Description,
			Createtime:  v.Createtime,
			Operator:    "丁丽丽",
			MoneyType:   "人民币",
			Type:        tpValue,
		})
	}
	return ret
}
