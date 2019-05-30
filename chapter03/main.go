package main

import (
	"github.com/astaxie/beego" // 导入beego包
	"github.com/astaxie/beego/context" // 导入beego/context包
)

func main() {
	/**
	* 定义路由函数
	* 处理路径为/ 方法为GET的请求
	* 向客户端响路由参数
	*/

	// 匹配/v1/和/v1/xxx/格式的路径
	beego.Get("/v1/?:id/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("Get id:" + ctx.Input.Param(":id"))) //获取参数
	})

	// 匹配/v2/xxx/格式的路径
	beego.Get("/v2/:id/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("Get id:" + ctx.Input.Param(":id")))
	})

	// 匹配/v3/1/格式的路径
	beego.Get("/v3/:id([0-9]{1,5})/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("Get id:" + ctx.Input.Param(":id")))
	})

	// 匹配/v4/1/格式的路径
	beego.Get(`/v4/:id(\d{1,5})/`, func(ctx *context.Context) {
		ctx.Output.Body([]byte("Get id:" + ctx.Input.Param(":id")))
	})

	// 匹配/v5/1/格式的路径
	beego.Get(`/v5/:id:int/`, func(ctx *context.Context) {
		ctx.Output.Body([]byte("Get id:" + ctx.Input.Param(":id")))
	})

	// 匹配/v6/开头格式的路径
	beego.Get(`/v6/*`, func(ctx *context.Context) {
		ctx.Output.Body([]byte("Get splat:" + ctx.Input.Param(":splat")))
	})

	// 匹配/v7/开头格式路径中的文件路径和文件后缀
	beego.Get(`/v7/*.*`, func(ctx *context.Context) {
		ctx.Output.Body([]byte("Get path:" + ctx.Input.Param(":path") + ", ext:" + ctx.Input.Param(":ext")))
	})

	// 启动beego服务
	beego.Run()
}