package main

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {
	c.Layout = "layout.html"
	c.TplName = "index.html"
}

func main() {
	beego.AutoRouter(&HomeController{})
	beego.Run()
}