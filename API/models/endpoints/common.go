package endpoints

import (
	"accountBook/models/beans"
	"accountBook/models/beans/customer"
	"accountBook/models/beans/dbBeans"
	"accountBook/models/log"
	"accountBook/models/service"
	"fmt"
	"math"
	"strings"

	heldiamgo "github.com/kinwyb/go"

	"github.com/kinwyb/go/err1"
)

func init() {
	go saveLog()
}

//检测权限
//@param string fun 函数点
func CheckPower(fun string, ctx *beans.Context) err1.Error {
	//记录日志
	if v := buildLog(fun, ctx); v != nil {
		logChan <- v
	}
	// todo 检测权限
	return nil
}

// pws所有全值集合
// chkValue当前要校验的权限值
func powerVerification(pws []uint, chkValue uint) bool {
	index := int(math.Ceil(float64(chkValue) / 64)) //获取到权限在第几个值 == 1 => 2 , 0 => 1
	if len(pws) <= int(index) {
		//找不到相关权限对应的值
		return false
	}
	chkValue = chkValue % 64
	// 1左移d位表示权限对应的点位设置成1 .e.g  1<<3 => 00...00100
	// &将值进行确定权限power对应值的相应位置是否也是1来确定是否有权限
	return 1<<uint(chkValue)&pws[index] != 0
}

var funNameMap = map[string]string{
	"Bank.Add":        "新增银行",
	"Receipt.Add":     "新增收支单据",
	"ReceiptType.Add": "新增收支类型",
}

type LogContentParseFun func(db *customer.LogAddReq)

// 日志内容处理函数
var logContentParseFunMap = map[string]LogContentParseFun{}

func buildLog(fun string, ctx *beans.Context) *customer.LogAddReq {
	content := funNameMap[fun]
	if content == "" {
		return nil
	}
	arg := heldiamgo.JsonString(ctx.RequestArgs)
	return &customer.LogAddReq{
		LogDB: dbBeans.LogDB{
			CallFunc: fun,
			Args:     &arg,
			Action:   content,
			Content:  fmt.Sprintf("%s:执行了【%s】", "丁丽丽", content),
			Userid:   1,
		},
		Args: ctx.RequestArgs,
	}
}

var logChan = make(chan *customer.LogAddReq, 100)

func saveLog() {
	defer func() {
		if err := recover(); err != nil {
			log.Error(log.ServiceTag, "保存日志线程崩溃:%s", err)
			go saveLog()
		}
	}()
	for {
		lg := <-logChan
		if cfun, ok := logContentParseFunMap[strings.Split(lg.CallFunc, ".")[0]]; ok {
			cfun(lg)
		}
		ctx := beans.NewContext("saveLog")
		err := service.LogAdd(&lg.LogDB, ctx)
		if err != nil {
			log.Error(log.ServiceTag, "日志保存失败:%s", err)
		}
	}
}

// 注册日志内容处理函数
func RegisterLogContentParseFun(funString string, fun LogContentParseFun) {
	logContentParseFunMap[funString] = fun
}
