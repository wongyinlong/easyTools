package routers

import (
	. "dc/controllers/cust"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "User-Agent", "Authorization", "Token", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	// 后台管理路由
	nsAdmin := beego.NewNamespace("/admin")

	RegisterStaticRouter()
	RegisterAdminRouter(nsAdmin)

	beego.AddNamespace(nsAdmin)
}

func RegisterStaticRouter() {
	beego.Get("/favicon.ico", func(ctx *context.Context) {
		ctx.Output.Body([]byte(""))
	})
}

func RegisterAdminRouter(ns *beego.Namespace) {
	ns.Namespace(new(ExanpleAuthController).Router()) // 这里注册子路由
	ns.Namespace(new(AuthorizeController).Router())
	ns.Namespace(new(RoleController).Router())

}
