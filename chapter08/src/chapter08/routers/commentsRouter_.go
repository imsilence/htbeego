package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["chapter08/controllers:HomeController"] = append(beego.GlobalControllerRouter["chapter08/controllers:HomeController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/home/index/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chapter08/controllers:HomeController"] = append(beego.GlobalControllerRouter["chapter08/controllers:HomeController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/home/login/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chapter08/controllers:HomeController"] = append(beego.GlobalControllerRouter["chapter08/controllers:HomeController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/home/logout/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
