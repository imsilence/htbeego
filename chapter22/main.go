package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type LoginForm struct {
	Username string `form:"username" valid:"Required"`
	Password string `form:"password" valid:"Required"`
}

func (f *LoginForm) Valid(v *validation.Validation) {
	if f.Username != "kk" || f.Password != "123456" {
		v.SetError("username", "用户名或密码错误")
	}
}

type UserController struct {
	beego.Controller
}

func (c *UserController) Login() {
	username := c.GetString("username", "")
	password := c.GetString("password", "")
	var valid validation.Validation

	valid.Required(username, "username").Message("用户名或密码为空")
	valid.Required(password, "password").Message("用户名或密码为空")

	if valid.HasErrors() {
		for _, e := range valid.Errors {
			fmt.Println(e.Key, e.Message)
		}
	}

	var lf LoginForm
	var valid2 validation.Validation

	if err := c.ParseForm(&lf); err == nil {
		correct, err := valid2.Valid(&lf)
		if err != nil {
			fmt.Println(err)
		}
		if !correct {
			for _, e := range valid2.Errors {
				fmt.Println(e.Key, e.Message)
			}
		}
	} else {
		fmt.Println(err)
	}
	c.ServeJSON()
}

func main() {
	beego.AutoRouter(&UserController{})
	beego.Run()
}