package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	fmt.Println("Base Prepare")
}

type AuthController struct {
	BaseController
}

func (c *AuthController) Authenticate() {
	c.Ctx.Output.Body([]byte("Unauthorized"))
}

func (c *AuthController) Prepare() {
	c.BaseController.Prepare()
	fmt.Println("Auth Prepare")
	c.Authenticate()
}

type UserController struct{
	BaseController
}

func (c *UserController) Login() {
	c.Ctx.Output.Body([]byte("User:Login"))
	fmt.Println("User:Login")
}

func (c *UserController) Logout() {
	c.Ctx.Output.Body([]byte("User:Logout"))
	fmt.Println("User:Logout")
}

type HomeController struct {
	AuthController
}

func (c *HomeController) Index() {
	c.Ctx.Output.Body([]byte("Home:Index"))
	fmt.Println("Home:Index")
}

func main() {
	beego.AutoRouter(&HomeController{})
	beego.AutoRouter(&UserController{})
	beego.Run()
}