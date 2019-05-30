package routers

import (
	"quickstart/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"fmt"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Head("/method/", func(ctx *context.Context) {
		ctx.WriteString("HEAD")
	})
	beego.Get("/method/", func(ctx *context.Context) {
		ctx.WriteString("GET")
	})
	beego.Post("/method/", func(ctx *context.Context) {
		ctx.WriteString("POST")
	})
	beego.Put("/method/", func(ctx *context.Context) {
		ctx.WriteString("PUT")
	})
	beego.Delete("/method/", func(ctx *context.Context) {
		ctx.WriteString("DELETE")
	})

	beego.Any("/any/:id:int/", func(ctx *context.Context) {
		fmt.Println(ctx.Input.Param(":id"))

		ctx.WriteString(ctx.Input.Method())
	})

	beego.Router("/user/login/", &controllers.UserController{}, "get:Open")
	beego.Router("/user/login/", &controllers.UserController{}, "post:Login")
	beego.Router("/user/logout/", &controllers.UserController{}, "*:Logout")

	beego.InsertFilter("/*", beego.BeforeStatic, func(ctx *context.Context) {
		fmt.Println("BeforeStatic")
		fmt.Println(ctx.Input.RunController)
		fmt.Println(ctx.Input.RunMethod)
	}, true)

	beego.InsertFilter("/*", beego.BeforeRouter, func(ctx *context.Context) {
		fmt.Println("BeforeRouter")
		fmt.Println(ctx.Input.RunController)
		fmt.Println(ctx.Input.RunMethod)
	}, true)

	beego.InsertFilter("/*", beego.BeforeExec, func(ctx *context.Context) {
		fmt.Println("BeforeExec")
		fmt.Println(ctx.Input.RunController)
		fmt.Println(ctx.Input.RunMethod)
	}, true)

	beego.InsertFilter("/*", beego.AfterExec, func(ctx *context.Context) {
		fmt.Println("AfterExec")
		fmt.Println(ctx.Input.RunController)
		fmt.Println(ctx.Input.RunMethod)
	}, false)

	beego.InsertFilter("/*", beego.FinishRouter, func(ctx *context.Context) {
		fmt.Println("FinishRouter")
		fmt.Println(ctx.Input.RunController)
		fmt.Println(ctx.Input.RunMethod)
	}, false)
}
