package routers

import (
	"chapter08/controllers"
	"github.com/astaxie/beego" // 导入beego包
)

func init() {
	// 注解路由
	beego.Include(&controllers.HomeController{})
}
