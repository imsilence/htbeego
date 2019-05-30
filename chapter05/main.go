package main

import (
	"github.com/astaxie/beego" // 导入beego包
)

// 定义控制器
type HomeController struct {
	beego.Controller
}

// 定义GET请求处理方法
func (c *HomeController) Get() {
	c.Ctx.Output.Body([]byte("GET:" + c.Ctx.Input.Param(":id")))
}


func main() {
	// 定义处理路径/1/的控制器路由
	beego.Router("/:id:int/", &HomeController{})

	// 启动beego服务
	beego.Run()
}