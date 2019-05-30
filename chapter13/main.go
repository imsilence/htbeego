package main

import (
	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) Get() {
	c.TplName="login.html"
}

func (c *AuthController) Post() {
	c.Redirect("/home/", 302)
}

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {
	c.TplName = "index.html"
	c.Layout = "layout.html"
}

func main() {
	beego.Router("/auth/", &AuthController{})
	beego.Router("/home/", &HomeController{}, "*:Index")
	beego.Run()
}