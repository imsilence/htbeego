package main

import (
	"fmt"
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:881019@tcp(127.0.0.1:3306)/htbeego?charset=utf8mb4&loc=Asia%2FShanghai")

	ormer := orm.NewOrm()
	rawset := ormer.Raw("insert into user(name, password, gender, birthday, `desc`, created, updated) values(?, ?, ?, ?, ?, ?, ?)")
	preparer, err := rawset.Prepare()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer preparer.Close()

	count := 10
	birthday, _ := time.Parse("2006-01-02", "1988-10-19")
	now := time.Now()

	for i := 0; i < count; i++ {
		preparer.Exec(fmt.Sprintf("kk%d", i), "123qwe!@#", i%2 == 0, birthday, "", now, now)
	}
}