package cust

import (
	"app/controllers/basic"

	"github.com/astaxie/beego"
)

// 这里组合了 basic文件夹中的 baseController。组合了这个controller的controller可以直接访问。
// 其原理是因为beego给controller提供了 prepare 这个钩子。会自动执行，重新实现它让它具有了检查功能。
type ExampleBaseController struct {
	basic.BaseController
}

func (c *ExampleBaseController) Router() *beego.Namespace {
	ns := beego.NewNamespace("/base",
		beego.NSRouter("/get", c, "get:GetFunc"),
		beego.NSRouter("/post", c, "post:PostFunc"),
		// 这里的路由 就是路由的用法嘛。
	)
	return ns
}

func (c *ExampleBaseController) GetFunc() {

	c.StopRun()
}

func (c *ExampleBaseController) PostFunc() {

	c.StopRun()
}
