package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	fmt.Println(beego.AppConfig.DefaultString("MYSQLHOST", "127.0.0.1"))
	beego.Run()
}