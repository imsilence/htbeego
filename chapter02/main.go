package main

import (
	"time" //导入time包

	"github.com/astaxie/beego" // 导入beego包
	"github.com/astaxie/beego/context" // 导入beego/context包
)

func main() {
	/**
	* 定义路由函数
	* 处理路径为/ 方法为GET的请求
	* 向客户端响应当前时间
	*/
	beego.Get("/", func(ctx *context.Context) {
		// 写入当前时间字符串到响应流
		ctx.Output.Body([]byte(time.Now().Format("2006-01-02 15:04:05")))
	})

	// 启动beego服务
	beego.Run()
}