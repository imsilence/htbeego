package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {
	k := "key"
	v := c.GetSession(k)
	fmt.Println(v)
	value := 0
	if vv, ok := v.(int); ok {
		value = vv
	}
	fmt.Println(value)
	c.SetSession(k, value + 1)
	c.Data["json"] = value
	c.ServeJSON()
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.AutoRouter(&HomeController{})
	beego.Run()
}