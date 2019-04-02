// @APIVersion 2.0.0
// @Title accountBook
// @Description accountBook
// @Contact wangyb@zhifangw.cn
// @TermsOfServiceUrl http://www.zhifangw.cn
package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1")
	beego.AddNamespace(ns)
	idx := &indexCtl{}
	beego.Router("/", idx, "*:Index")
}

//首页控制器
type indexCtl struct {
	beego.Controller
}

func (i *indexCtl) Index() {
	i.Redirect("/swagger", 301)
}
