package routers

import (
	"bee-web/controllers"
	api1 "bee-web/controllers/api/v1"
	api2 "bee-web/controllers/api/v2"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/satori/go.uuid"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// 用于为每个请求，添加 request uuid ，用于请求追踪
	var FilterRequestID = func(ctx *context.Context) {
		p, _ := uuid.NewV4()
		requestId := p.String()
		ctx.Input.SetData("requestId", requestId)
	}
	beego.InsertFilter("/api/*", beego.BeforeRouter, FilterRequestID)

	ns :=
		beego.NewNamespace("/api",
			beego.NSGet("/health", func(ctx *context.Context) {
				ctx.Output.Body([]byte("API is alived."))
			}),
			beego.NSNamespace("/v1",
				beego.NSRouter("/version", &api1.VersionController{}),
				beego.NSRouter("/node", &api1.NodeController{}, "get:ListNode;post:CreateNode"),
			),
			beego.NSNamespace("/v2",
				beego.NSRouter("/version", &api2.VersionController{}),
			),
		)
	beego.AddNamespace(ns)
}
