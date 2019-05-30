package main

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {
	c.Redirect(beego.URLFor("UserController.Login"), 302)
}

type UserController struct {
	beego.Controller
}

func (c *UserController) Login() {
	c.TplName = "login.html"
}

func main() {
	beego.AutoRouter(&HomeController{})
	beego.AutoRouter(&UserController{})
	beego.Run()
}