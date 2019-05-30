package main

import (
	"time"
	"github.com/astaxie/beego" // 导入beego包
)

// 定义控制器
type HomeController struct {
	beego.Controller
}

// 处理GET和HEAD请求
func (c *HomeController) GetFunc() {
	c.Ctx.Output.Body([]byte("GET:" + time.Now().Format("2006-01-02 15:04:05")))
}

// 处理POST请求
func (c *HomeController) PostFunc() {
	c.Ctx.Output.Body([]byte("POST:" + time.Now().Format("2006-01-02 15:04:05")))
}

// 处理其他类型请求
func (c *HomeController) OtherFunc() {
	c.Ctx.Output.Body([]byte("Other:" + time.Now().Format("2006-01-02 15:04:05")))
}

func main() {
	// 定义处理路径/的控制器路由
	/**
	* HEAD和GET请求由函数GetFunc处理
	* POST请求由PostFunc处理
	* 其他请求由OtherFunc处理
	*/
	beego.Router("/", &HomeController{}, "head,get:GetFunc;post:PostFunc;*:OtherFunc")

	// 启动beego服务
	beego.Run()
}