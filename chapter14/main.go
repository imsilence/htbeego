package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strconv"
)

type HomeController struct{
	beego.Controller
}

func (c *HomeController) Index() {

	fmt.Println(c.GetControllerAndAction())

	fmt.Println(c.Ctx.Input.Method(), c.Ctx.Input.IsAjax(), c.Ctx.Input.IsWebsocket())

	fmt.Println(c.Ctx.Input.Header("User-Agent"))

	fmt.Println(c.GetString("name", "kk"))
	fmt.Println(c.GetInt("age", 18))
	fmt.Println(c.GetFloat("score", 0.0))
	fmt.Println(c.GetBool("boy", true))
	fmt.Println(c.GetStrings("hobby", []string{}))


	if f, h, err := c.GetFile("txt"); err == nil {
		defer f.Close()
		fmt.Println(h.Filename, h.Size, h.Header)
		bytes, _ := ioutil.ReadAll(f)
		fmt.Println(string(bytes))
	} else {
		fmt.Println(err)
	}

	c.SaveToFile("txt", "txt.txt")

	var u struct {
		Name string `form:"name"`
		Age int `form:"age"`
		Score float64 `form:"score"`
		Boy bool `form:"boy"`
		Hobby []string `form:"hoby"`
	}

	if err := c.ParseForm(&u); err == nil {
		fmt.Println(u)
	} else {
		fmt.Println(err)
	}

	fmt.Println(string(c.Ctx.Input.CopyBody(1024 * 1024 * 10)))
	fmt.Println(string(c.Ctx.Input.RequestBody))

	var (
		name string
		age int
	)
	fmt.Println(c.Ctx.Input.Bind(&name, "name"), name)
	fmt.Println(c.Ctx.Input.Bind(&age, "age"), age)

	secret := "123abc!@#"
	skey, key := "skey", "key"
	svv, vv := 0, 0
	if v, ok := c.GetSecureCookie(secret, skey); ok {
		svv, _ = strconv.Atoi(v)
	}
	fmt.Println(svv)
	c.SetSecureCookie(secret, skey, strconv.Itoa(svv+1))


	if v, err := strconv.Atoi(c.Ctx.Input.Cookie(key)); err == nil {
		vv = v
	}
	fmt.Println(vv)
	c.Ctx.Output.Cookie(key, strconv.Itoa(vv+1))

	c.Ctx.Output.Body([]byte("Index"))
}


func main() {
	beego.AutoRouter(&HomeController{})
	beego.Run()
}