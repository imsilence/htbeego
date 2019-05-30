package controllers

import (
	"os"
	"io"
	"fmt"
	"net/http"
	"html/template"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Open() {
	this.Data["xsrf_token"] = this.XSRFToken()
	this.Data["xsrf_input"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "user/login.html"
}

func (this *UserController) Login() {
	var (
		username string
		password string
	)

	this.Ctx.Input.Bind(&username, "username")
	this.Ctx.Input.Bind(&password, "password")
	fmt.Println(username, password)


	fmt.Println(this.GetString("username"), this.GetString("password"))
	fmt.Println(this.Input().Get("username"), this.Input().Get("password"))

	var user struct {
		UserName string `form:"username"`
		Password string `form:"password"`
	}

	if err := this.ParseForm(&user); err == nil {
		fmt.Println(user)
	} else {
		fmt.Println(user)
	}

	fmt.Println(string(this.Ctx.Input.RequestBody))

	if file, header, err := this.GetFile("file"); err == nil {
		defer file.Close()
		fmt.Println(header.Filename)
		io.Copy(os.Stdout, file)
	} else {
		fmt.Println(err)
	}


	this.Data["username"] = username
	this.Data["password"] = password
	this.Data["xsrf_token"] = this.XSRFToken()
	this.Data["xsrf_input"] = template.HTML(this.XSRFFormHTML())

	var count int
	if vv, ok := this.GetSession("count").(int); ok {
		count = vv
	}
	this.SetSession("count", count + 1)
	this.SessionRegenerateID()

	this.Data["count"] = count
	this.TplName = "user/login.html"
}

func (this *UserController) Logout() {
	this.DestroySession()
	this.Redirect(beego.URLFor(".Open"), http.StatusMovedPermanently)
}
