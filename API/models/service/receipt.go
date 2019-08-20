package service

import (
	"accountBook/models/beans"
	"accountBook/models/beans/customer"
	"accountBook/models/beans/dbBeans"
	"accountBook/models/dataBase"
	"fmt"
	"strings"
	"time"

	heldiamgo "github.com/kinwyb/go"

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
		obj := &customer.Receipt{
			Id:          fmt.Sprintf("S%09d", v.Id),
			Money:       v.Money,
			Bank:        bankMap[v.BankId],
			Description: v.Description,
			Createtime:  v.Createtime,
			//Operator:    "丁丽丽",
			//MoneyType:   "人民币",
			Type: tpValue,
			Shop: shopValue,
		}
		if v.MoneyType == beans.USD {
			obj.MoneyType = "美金"
		} else {
			obj.MoneyType = "人民币"
		}
		ret.Data = append(ret.Data, obj)
	}
	// 计算汇总数据
	if pg.Page < 2 {
		ret.Counts = receiptListCount(req, ctx.Child())
	}
	return ret
}

// 计算汇总数据
func receiptListCount(req *customer.ReceiptListReq, ctx *beans.Context) []*customer.ReceiptCount {
	defer ctx.Start("sev.receiptListCount").Finish()
	//银行期初
	startCNY := 0.0
	startUSD := 0.0
	banks := dataBase.BankList(ctx.Child())
	for _, v := range banks {
		if req.BankID > 0 && req.BankID != v.Id {
			continue
		}
		startCNY = startCNY + v.BankMoney
		startUSD = startUSD + v.BankMoneyUsa
	}
	// 流水金额
	startData := map[string]float64{}
	endData := map[string]float64{}
	if req.StartTime != "" {
		result := dataBase.ReceiptEndTimeMoneyCount(req.StartTime, req.BankID, true, ctx.Child())
		for _, v := range result {
			key := fmt.Sprintf("%d", v.MoneyType)
			if v.Money < 0 {
				key = key + "Out"
			} else {
				key = key + "In"
			}
			startData[key] = startData[key] + v.Money
		}
	}
	if req.EndTime == "" {
		req.EndTime = time.Now().Format(heldiamgo.DateTimeFormat)
	}
	result := dataBase.ReceiptEndTimeMoneyCount(req.EndTime, req.BankID, true, ctx.Child())
	for _, v := range result {
		key := fmt.Sprintf("%d", v.MoneyType)
		if v.Money < 0 {
			key = key + "Out"
		} else {
			key = key + "In"
		}
		endData[key] = endData[key] + v.Money
	}
	// 汇总人民币
	rmb := &customer.ReceiptCount{
		MoneyType:  "人民币",
		AllIn:      endData[fmt.Sprintf("%dIn", beans.CNY)] - startData[fmt.Sprintf("%dIn", beans.CNY)],
		AllOut:     endData[fmt.Sprintf("%dOut", beans.CNY)] - startData[fmt.Sprintf("%dOut", beans.CNY)],
		StartMoney: startData[fmt.Sprintf("%dIn", beans.CNY)] + startData[fmt.Sprintf("%dOut", beans.CNY)],
		EndMoney:   endData[fmt.Sprintf("%dIn", beans.CNY)] + endData[fmt.Sprintf("%dOut", beans.CNY)],
	}
	rmb.EndMoney = rmb.EndMoney + startCNY
	rmb.StartMoney = rmb.StartMoney + startCNY
	usd := &customer.ReceiptCount{
		MoneyType:  "美金",
		AllIn:      endData[fmt.Sprintf("%dIn", beans.USD)] - startData[fmt.Sprintf("%dIn", beans.USD)],
		AllOut:     endData[fmt.Sprintf("%dOut", beans.USD)] - startData[fmt.Sprintf("%dOut", beans.USD)],
		StartMoney: startData[fmt.Sprintf("%dIn", beans.USD)] + startData[fmt.Sprintf("%dOut", beans.USD)],
		EndMoney:   endData[fmt.Sprintf("%dIn", beans.USD)] + endData[fmt.Sprintf("%dOut", beans.USD)],
	}
	usd.EndMoney = usd.EndMoney + startUSD
	usd.StartMoney = usd.StartMoney + startUSD
	return []*customer.ReceiptCount{rmb, usd}
}

// 时间范围内容金额统计
func ReceiptEndTimeMoneyCount(endTime string, ctx *beans.Context) []*dbBeans.Receipt {
	defer ctx.Start("sev.ReceiptTimeRangeMoneyCount").Finish()
	return dataBase.ReceiptEndTimeMoneyCount(endTime, 0, false, ctx.Child())
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
