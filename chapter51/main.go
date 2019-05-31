package main

import (
	"strings"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {
	c.Data["name"] = "kk"
	c.Data["age"] = 30
	c.Data["boy"] = true
	c.Data["teachers"] = []string{"kk", "woniu"}
	c.Data["scores"] = map[string]float64{"kk" : 8, "woniu" : 9}
	c.Data["students"] = []struct{
		Name string
		Addr string
	}{{"小明", "北京"}, {"小李", "西安"}, {"小黑", "北京"}}

	c.TplName = "index.html"
}

func main() {
	beego.AddFuncMap("upper", func(in string) string {
		return strings.ToUpper(in)
	})
	beego.BConfig.RunMode = "dev"
	beego.AutoRouter(&HomeController{})
	beego.Run()
}