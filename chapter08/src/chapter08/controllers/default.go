package controllers

import (
	"time"

	"github.com/astaxie/beego" // 导入beego包
)

// 定义控制器
type HomeController struct {
	beego.Controller
}

func (c *HomeController) URLMapping() {
	c.Mapping("Index", c.Index)
	c.Mapping("Login", c.Login)
	c.Mapping("Logout", c.Logout)
}

// @router /home/index/ [get]
func (c *HomeController) Index() {
	c.Ctx.Output.Body([]byte("Index:" + time.Now().Format("2006-01-02 15:04:05")))
}

// @router /home/login/ [post]
func (c *HomeController) Login() {
	c.Ctx.Output.Body([]byte("Login:" + time.Now().Format("2006-01-02 15:04:05")))
}

// @router /home/logout/ [get]
func (c *HomeController) Logout() {
	c.Ctx.Output.Body([]byte("Logout:" + time.Now().Format("2006-01-02 15:04:05")))
}
