package main

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {
	c.Data["json"] = c.XSRFToken()
	c.ServeJSON()
}

func main() {
	beego.BConfig.WebConfig.EnableXSRF = true
	beego.AutoRouter(&HomeController{})
	beego.Run()
}