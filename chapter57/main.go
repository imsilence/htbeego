package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	beego.BConfig.Listen.EnableAdmin = true
	beego.Get("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hi"))
	})
	beego.Run()
}