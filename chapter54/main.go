package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

type LoginForm struct {
	Username string `form:"username,text,用户名"`
	Password string `form:"password,password,密码"`
}

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Login() {
	form := &LoginForm{}
	if c.Ctx.Input.IsPost() {
		err := c.ParseForm(form)
		fmt.Println(err, form)
	}
	c.Data["form"] = form
	c.TplName = "login.html"
}

func main() {
	beego.BConfig.RunMode = "dev"
	beego.AutoRouter(&HomeController{})
	beego.Run()
}