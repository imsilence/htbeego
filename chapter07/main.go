package main

import (
	"time"
	"github.com/astaxie/beego" // 导入beego包
)

// 定义控制器
type HomeController struct {
	beego.Controller
}

// 处理/home/index/路径请求
func (c *HomeController) Index() {
	c.Ctx.Output.Body([]byte("Index:" + time.Now().Format("2006-01-02 15:04:05")))
}

// 处理/home/login/路径请求
func (c *HomeController) Login() {
	c.Ctx.Output.Body([]byte("Login:" + time.Now().Format("2006-01-02 15:04:05")))
}

// 处理/home/logout/路径请求
func (c *HomeController) Logout() {
	c.Ctx.Output.Body([]byte("Logout:" + time.Now().Format("2006-01-02 15:04:05")))
}

func main() {
	// 定义自动匹配路由
	beego.AutoRouter(&HomeController{})

	// 启动beego服务
	beego.Run()
}