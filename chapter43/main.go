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

	if err := ormer.Begin(); err == nil {
		ormer.Raw("update user set gender = ?", 0).Exec()
		_, err := ormer.Raw("insert into user(name) values(?)", "kk").Exec()
		if err == nil {
			ormer.Commit()
		} else {
			fmt.Println(err)
			ormer.Rollback()
		}
	}
}