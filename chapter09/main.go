package main

import (
	"time"
	"github.com/astaxie/beego" // 导入beego包
	"github.com/astaxie/beego/context" // 导入beego包
)

// 定义控制器
type HomeController struct {
	beego.Controller
}

// 处理/v1/home/index/路径请求
func (c *HomeController) Index() {
	c.Ctx.Output.Body([]byte("Index:" + time.Now().Format("2006-01-02 15:04:05")))
}

type UserController struct {
	beego.Controller
}

// 处理/v1/login/路径请求
func (c *UserController) Login() {
	c.Ctx.Output.Body([]byte("Login:" + time.Now().Format("2006-01-02 15:04:05")))
}

// 处理/v1/logout/路径请求
func (c *UserController) Logout() {
	c.Ctx.Output.Body([]byte("Logout:" + time.Now().Format("2006-01-02 15:04:05")))
}

func main() {
	// 定义命名空间
	ns := beego.NewNamespace("/v1",
		beego.NSGet("/", func(ctx *context.Context) {
			ctx.Output.Body([]byte("ns:" + time.Now().Format("2006-01-02 15:04:05")))
		}),
		beego.NSRouter("/login", &UserController{}, "get,post:Login"),
		beego.NSRouter("/logout", &UserController{}, "*:Logout"),
		beego.NSAutoRouter(&HomeController{}),
	)

	// 添加命名空间
	beego.AddNamespace(ns)
	// 启动beego服务
	beego.Run()
}