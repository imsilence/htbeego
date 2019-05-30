package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	c.Controller.Init(ctx, controllerName, actionName, app)
	fmt.Println("Init")
}

func (c *HomeController) Prepare() {
	fmt.Println("Prepare")
}

func (c *HomeController) Get() {
	fmt.Println("Get")
	c.Ctx.Output.Body([]byte("Get"))
}

func (c *HomeController) Post() {
	fmt.Println("Post")
	c.Ctx.Output.Body([]byte("Post"))
}

func (c *HomeController) Render() error {
	fmt.Println("Render")
	return c.Controller.Render()
}

func (c *HomeController) Finish() {
	fmt.Println("Finish")
}

func main() {
	beego.Router("/", &HomeController{})
	beego.Run()
}