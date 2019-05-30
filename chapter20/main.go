package main

import (
	"net/http"
	"fmt"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Error() {
	key := c.GetString("key", "404")
	c.Abort(key)
}

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error500() {
	c.TplName = "error/500.html"
}

func (c *ErrorController) ErrorDb() {
	c.TplName= "error/db.html"
}

func main() {
	beego.ErrorHandler("404", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ERROR 404")
	})
	beego.ErrorController(&ErrorController{})
	beego.AutoRouter(&HomeController{})
	beego.Run()
}