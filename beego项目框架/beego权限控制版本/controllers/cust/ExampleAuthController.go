package cust

import (
	"app/controllers/basic"

	"github.com/astaxie/beego"
)

// 这里组合了 basic文件夹中的 AuthorizeController。组合了这个controller的controller不可以直接访问。
// 其原理是因为beego给controller提供了 prepare 这个钩子。会自动执行，重新实现它让它具有了检查功能。
type ExanpleAuthController struct {
	basic.AuthorizeController
}

func (c *ExanpleAuthController) Router() *beego.Namespace {
	ns := beego.NewNamespace("/auth",
		beego.NSRouter("/get", c, "get:GetFunc"),
		beego.NSRouter("/post", c, "post:PostFunc"),
	)
	return ns
}

func (c *ExanpleAuthController) GetFunc() {

	c.StopRun()
}

func (c *ExanpleAuthController) PostFunc() {

	c.StopRun()
}
