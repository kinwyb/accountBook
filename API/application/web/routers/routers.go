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
	beego.Router("/v1/web/bank/list", bank, "GET:List")
}

//首页控制器
type indexCtl struct {
	beego.Controller
}

func (i *indexCtl) Index() {
	i.Redirect("/swagger", 301)
}
