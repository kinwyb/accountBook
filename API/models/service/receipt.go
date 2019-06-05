package service

import (
	"accountBook/models/beans"
	"accountBook/models/beans/customer"
	"accountBook/models/beans/dbBeans"
	"accountBook/models/dataBase"
	"fmt"
	"strings"

	"github.com/kinwyb/go/err1"

	"github.com/kinwyb/go/db"
)

// 收支列表
func ReceiptList(req *customer.ReceiptListReq, pg *db.PageObj, ctx *beans.Context) *customer.ReceiptListResp {
	defer ctx.Start("sev.ReceiptList").Finish()
	if req.ShopID > 0 {
		rShop := dataBase.ReceiptTypeQueryByID(req.ShopID, ctx.Child())
		if rShop == nil { //没有这个选择
			return nil
		}
		if req.ReceiptType != "" {
			rType := dataBase.ReceiptTypeQueryByParentIDAndName(rShop.Id, req.ReceiptType, ctx.Child())
			if rType == nil {
				return nil
			}
			req.ReceiptType = fmt.Sprintf("%d", rType.Id)
		} else {
			rTypes := dataBase.ReceiptTypeList(rShop.Id, ctx.Child())
			var rTypeArray []string
			for _, v := range rTypes {
				rTypeArray = append(rTypeArray, fmt.Sprintf("%d", v.Id))
			}
			rTypeArray = append(rTypeArray, fmt.Sprintf("%d", rShop.Id))
			req.ReceiptType = strings.Join(rTypeArray, ",")
		}
	} else if req.ReceiptType != "" {
		rTypes := dataBase.ReceiptTypeQueryByName(req.ReceiptType, ctx.Child())
		if len(rTypes) < 1 {
			return nil
		}
		var rTypeArray []string
		for _, v := range rTypes {
			rTypeArray = append(rTypeArray, fmt.Sprintf("%d", v.Id))
		}
		req.ReceiptType = strings.Join(rTypeArray, ",")
	}
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
	ret := &customer.ReceiptListResp{}
	for _, v := range result {
		tp := receiptTypeMap[v.Type]
		tpValue := ""
		shopValue := ""
		if tp != nil {
			tpValue = tp.Name
			tpParent := receiptTypeMap[tp.ParentId]
			if tpParent != nil && tp.ParentId != 0 {
				shopValue = tpParent.Name
			} else { //没有父级,就表示直接是店铺
				shopValue = tpValue
				tpValue = ""
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
			Shop:        shopValue,
		})
	}
	return ret
}

// 时间范围内容金额统计
func ReceiptEndTimeMoneyCount(endTime string, ctx *beans.Context) []*dbBeans.Receipt {
	defer ctx.Start("sev.ReceiptTimeRangeMoneyCount").Finish()
	return dataBase.ReceiptEndTimeMoneyCount(endTime, ctx.Child())
}

// 下一个
func ReceiptNextNo(ctx *beans.Context) string {
	defer ctx.Start("sev.ReceiptLastID").Finish()
	id := dataBase.ReceiptLastID(ctx.Child())
	return fmt.Sprintf("SJ%07d", id+1)
}

// 新增收支单
func ReceiptAdd(req *dbBeans.ReceiptDB, ctx *beans.Context) err1.Error {
	defer ctx.Start("sev.ReceiptAdd").Finish()
	req.Id = 0
	req.Lastmodify = ""
	req.Operator = 1
	return dataBase.ReceiptAdd(req, ctx.Child())
}
