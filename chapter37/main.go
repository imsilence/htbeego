package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:881019@tcp(127.0.0.1:3306)/htbeego?charset=utf8mb4&loc=Asia%2FShanghai")


	ormer := orm.NewOrm()
	rawset := ormer.Raw("select name, count(*) as cnt from user group by name")

	var values []orm.Params
	rawset.Values(&values)
	for _, v := range values {
		fmt.Println(v["name"], v["cnt"])
	}
}