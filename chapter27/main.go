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
	user := &User{Id: 1}
	if err := ormer.Read(user); err == nil {
		fmt.Println(user)
	} else {
		fmt.Println(err)
	}

	user = &User{Name: "kk"}
	if err := ormer.Read(user, "Name"); err == nil {
		fmt.Println(user)
	} else {
		fmt.Println(err)
	}

	user = &User{Gender: true}
	if err := ormer.Read(user, "Gender"); err == nil {
		fmt.Println(user)
	} else {
		fmt.Println(err)
	}

	user = &User{Gender: true, Name : "kk8"}
	if err := ormer.Read(user, "Gender", "Name"); err == nil {
		fmt.Println(user)
	} else {
		fmt.Println(err)
	}
}