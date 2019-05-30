package main

import (
	"fmt"
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int
	Name string `orm:"size(64);"`
	Password string `orm:"size(1024)"`
	Gender bool `orm:""`
	Birthday time.Time `orm:"type(date)"`
	Integral float64 `orm:"digits(12);decimals(3)"`
	Desc string `orm:"type(text)"`
	Status int `orm:""`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func main() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:881019@tcp(127.0.0.1:3306)/htbeego?charset=utf8mb4&loc=Asia%2FShanghai")

	orm.RegisterModel(&User{})

	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&User{})
	inserter, _ := queryset.PrepareInsert()

	defer inserter.Close()

	count := 10
	birthday, _ := time.Parse("2006-01-02", "1988-10-19")

	for i := 0; i < count; i++ {
		fmt.Println(inserter.Insert(&User{
			Name: fmt.Sprintf("kk%d", i),
			Password: "123abc!@#",
			Gender: i % 2 == 0,
			Birthday: birthday,
		}))
	}
}