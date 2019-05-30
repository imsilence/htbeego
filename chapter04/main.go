package main

import (
	"time"
	"github.com/astaxie/beego" // 导入beego包
)

// 定义控制器
type HomeController struct {
	beego.Controller
}

// 定义GET请求处理方法
func (c *HomeController) Get() {
	c.Ctx.Output.Body([]byte("GET:" + time.Now().Format("2006-01-02 15:04:05")))
}


// 定义POST请求处理方法
func (c *HomeController) Post() {
	c.Ctx.Output.Body([]byte("POST:" + time.Now().Format("2006-01-02 15:04:05")))
}

func main() {
	// 定义处理路径/的控制器路由
	beego.Router("/", &HomeController{})

	// 启动beego服务
	beego.Run()
}