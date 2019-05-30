package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

)

type HomeController struct {
	beego.Controller
}
func (c *HomeController) Index() {
	fmt.Println("index")
	c.Ctx.Output.Body([]byte("Index"))
}

func main() {
	beego.InsertFilter("/*", beego.BeforeStatic, func(ctx *context.Context) {
		fmt.Println("BeforeStatic")
		// 只要有输出流，这后续Filter和Router不执行
	}, false, true)

	beego.InsertFilter("/*", beego.BeforeRouter, func(ctx *context.Context) {
		fmt.Println("BeforeRouter")
		// 有输出流, 则后续Router不执行，Filter取决于returnOnOutput值(false执行，true不执行)
	}, false, true)

	beego.InsertFilter("/*", beego.BeforeExec, func(ctx *context.Context) {
		fmt.Println("BeforeExec")
		// 有输出流, 则后续Router不执行，Filter取决于returnOnOutput值(false执行，true不执行)
	}, false, true)

	beego.InsertFilter("/*", beego.AfterExec, func(ctx *context.Context) {
		fmt.Println("AfterExec")
		// 有输出流, 则后续Filter取决于returnOnOutput值(false执行，true不执行)
	}, false, true)

	beego.InsertFilter("/*", beego.FinishRouter, func(ctx *context.Context) {
		fmt.Println("FinishRouter")
		// 有输出流, 则后续Filter取决于returnOnOutput值(false执行，true不执行)
	}, false, true)

	beego.AutoRouter(&HomeController{})
	beego.Run()
}