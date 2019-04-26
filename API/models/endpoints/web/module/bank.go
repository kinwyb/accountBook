package module

import (
	"accountBook/models/beans"
	"accountBook/models/beans/customer"
	"accountBook/models/beans/dbBeans"
	"accountBook/models/endpoints"
	"accountBook/models/endpoints/web"
	"accountBook/models/service"
	"fmt"
	"time"

	"github.com/kinwyb/go"

	"github.com/kinwyb/go/err1"
)

var Bank web.IBankEndpoint = &bankEp{}

type bankEp struct{}

func (bankEp) Add(req *dbBeans.BankDB, ctx *beans.Context) err1.Error {
	defer ctx.Start("ep.BankAdd").Finish()
	if err := endpoints.CheckPower("BankAdd", ctx.Child()); err != nil {
		return err
	}
	return service.BankAdd(req, ctx.Child())
}

func (bankEp) ListComputeWithDay(startTime string, endTime string, ctx *beans.Context) ([]*customer.BankListCompateResp, err1.Error) {
	defer ctx.Start("ep.BankListComputeWithDay").Finish()
	if err := endpoints.CheckPower("ListComputeWithDay", ctx.Child()); err != nil {
		return nil, err
	}
	bankList := service.BankList(ctx.Child())
	startMoneyMap := map[string]float64{}
	endMoneyMap := map[string]float64{}
	if startTime != "" {
		startData := service.ReceiptEndTimeMoneyCount(startTime, ctx.Child())
		for _, v := range startData {
			startMoneyMap[fmt.Sprintf("%d-%d", v.BankId, v.MoneyType)] = v.Money
		}
	} else {
		for _, v := range bankList {
			startMoneyMap[fmt.Sprintf("%d-%d", v.Id, beans.CNY)] = v.BankMoney
			startMoneyMap[fmt.Sprintf("%d-%d", v.Id, beans.USD)] = v.BankMoneyUsa
		}
	}
	if endTime == "" {
		endTime = time.Now().Format(heldiamgo.DateTimeFormat)
	}
	endData := service.ReceiptEndTimeMoneyCount(endTime, ctx.Child())
	for _, v := range endData {
		endMoneyMap[fmt.Sprintf("%d-%d", v.BankId, v.MoneyType)] = v.Money
	}
	ret := make([]*customer.BankListCompateResp, len(bankList))
	for i, v := range bankList {
		r := &customer.BankListCompateResp{
			BankID:       v.Id,
			BankName:     v.BankName,
			SCNY:         startMoneyMap[fmt.Sprintf("%d-%d", v.Id, beans.CNY)],
			ECNY:         endMoneyMap[fmt.Sprintf("%d-%d", v.Id, beans.CNY)],
			SUSD:         startMoneyMap[fmt.Sprintf("%d-%d", v.Id, beans.USD)],
			EUSD:         endMoneyMap[fmt.Sprintf("%d-%d", v.Id, beans.USD)],
			BankAccount:  v.BankAccount,
			Contacts:     v.BankPeople,
			ContactPhone: v.BankPhone,
		}
		r.ECNY = r.SCNY + r.ECNY
		r.EUSD = r.SUSD + r.EUSD
		ret[i] = r
	}
	return ret, nil
}

func (bankEp) List(ctx *beans.Context) ([]*dbBeans.Bank, err1.Error) {
	defer ctx.Start("ep.BankList").Finish()
	if err := endpoints.CheckPower("BankList", ctx.Child()); err != nil {
		return nil, err
	}
	return service.BankList(ctx.Child()), nil
}
