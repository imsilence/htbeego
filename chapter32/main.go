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

	cond01 := orm.NewCondition()
	cond01 = cond01.And("name__contains", "silence")
	cond01 = cond01.And("gender", true)

	cond02 := orm.NewCondition()
	cond02 = cond02.And("name__contains", "kk")
	cond02 = cond02.And("gender", false)

	cond := orm.NewCondition()
	cond = cond.AndCond(cond01).OrCond(cond02)

	num, err := queryset.SetCond(cond).All(&users, "Id", "Name", "gender")
	fmt.Println(num, err)
	for _, user := range users {
		fmt.Println(user.Id, user.Name, user.Gender)
	}
}