// @APIVersion 2.0.0
// @Title accountBook
// @Description accountBook
// @Contact wangyb@zhifangw.cn
// @TermsOfServiceUrl http://www.zhifangw.cn
package routers

import (
	"accountBook/application/web/controllers/web"
	"accountBook/models/endpoints/web/module"

	"github.com/astaxie/beego"
)

func init() {
	webRouter()
	idx := &indexCtl{}
	beego.Router("/", idx, "*:Index")
}

//webb后台
func webRouter() {
	bank := &web.BankController{
		Serv: module.Bank,
	}
	// 银行
	beego.Router("/v1/web/bank/list", bank, "GET:List")
	// 收支类型
	receiptType := &web.ReceiptTypeController{
		Serv: module.ReceiptType,
	}
	beego.Router("/v1/web/receiptType/list", receiptType, "GET:List")
	beego.Router("/v1/web/receiptType/list/level", receiptType, "GET:ListByLevel")
	// 收支明细
	receipt := &web.ReceiptController{
		Serv: module.Receipt,
	}
	beego.Router("/v1/web/receipt/list", receipt, "POST:List")
}

//首页控制器
type indexCtl struct {
	beego.Controller
}

func (i *indexCtl) Index() {
	i.Redirect("/swagger", 301)
}
