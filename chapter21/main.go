package main

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {
	beego.Emergency("this is emergency")
	beego.Alert("this is alert")
	beego.Critical("this is critical")
	beego.Error("this is error")
	beego.Warning("this is warning")
	beego.Notice("this is notice")
	beego.Informational("this is informational")
	beego.Debug("this is debug")
	c.Ctx.Output.Body([]byte("Index"))
}

func main() {
	beego.SetLogger("file",
		`{
			"filename" : "web.log",
			"maxlines" : 0,
			"maxsize" : 0,
			"maxfiles" : 0,
			"daily" : true,
			"maxdays": 7,
			"rotate" : true,
			"level" : 7,
			"perm" : "0777"
		}`,
	)
	beego.SetLogFuncCall(true)
	beego.BConfig.Log.AccessLogs = true
	beego.AutoRouter(&HomeController{})
	beego.Run()
}