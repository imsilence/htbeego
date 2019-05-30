package main

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c HomeController) Head() {
	c.Ctx.Output.Body([]byte("Head"))
}

func (c HomeController) Get() {
	c.Ctx.Output.Body([]byte("Get"))
}

func (c HomeController) Put() {
	c.Ctx.Output.Body([]byte("Put"))
}

func (c HomeController) Patch() {
	c.Ctx.Output.Body([]byte("Patch"))
}

func (c HomeController) Post() {
	c.Ctx.Output.Body([]byte("Post"))
}

func (c HomeController) Delete() {
	c.Ctx.Output.Body([]byte("Delete"))
}

func (c HomeController) Options() {
	c.Ctx.Output.Body([]byte("Options"))
}

func main() {
	beego.Router("/home/", &HomeController{})
	beego.Run()
}