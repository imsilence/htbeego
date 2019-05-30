package main

import (
	"fmt"
	"github.com/astaxie/beego"
)


type HomeController struct {
	beego.Controller
}

func (c *HomeController) Login() {
	flash := beego.NewFlash()
	flash.Error("Error")
	flash.Warning("Warning")
	flash.Notice("Notice")
	flash.Success("Success")
	flash.Set("self", "Self")
	flash.Store(&c.Controller)
	c.Redirect(beego.URLFor("HomeController.Index"), 302)
}

func (c *HomeController) Index() {
	flash := beego.ReadFromRequest(&c.Controller)
	fmt.Println(flash.Data)
	c.TplName = "index.html"
}

func main() {
	beego.AutoRouter(&HomeController{})
	beego.Run()
}