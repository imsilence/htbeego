package main

import (
	_ "chapter08/routers"
	"github.com/astaxie/beego" // 导入beego包
)

func main() {
	// 启动beego服务
	beego.Run()
}