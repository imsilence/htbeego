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

	var users []*User
	num, err := queryset.OrderBy("-name", "id").All(&users, "Id", "Name")
	fmt.Println(num, err)
	for _, user := range users {
		fmt.Println(user.Id, user.Name)
	}

	num, err = queryset.OrderBy("name", "-id").All(&users, "Id", "Name")
	fmt.Println(num, err)
	for _, user := range users {
		fmt.Println(user.Id, user.Name)
	}
}