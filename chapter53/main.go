package main

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {
	c.Layout = "layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["SectionTitle"] = "title.html"
	c.LayoutSections["SectionStyles"] = "styles.html"
	c.LayoutSections["SectionScripts"] = "scripts.html"
	c.TplName = "index.html"
}

func main() {
	beego.BConfig.RunMode = "dev"
	beego.AutoRouter(&HomeController{})
	beego.Run()
}