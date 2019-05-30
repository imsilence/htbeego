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
	user := &User{Id: 10}
	if err := ormer.Read(user); err == nil {
		user.Name = "silence"
		num, err := ormer.Update(user)
		fmt.Println(num, err)
	} else {
		fmt.Println(err)
	}

	birthday, _ := time.Parse("2006-01-02", "1988-10-19")
	user = &User{Id:1, Name:"silence", Birthday:birthday}
	fmt.Println(ormer.Update(user))
	fmt.Println(user)

	user = &User{Id:11, Name:"silence", Birthday:birthday}
	fmt.Println(ormer.Update(user))
	fmt.Println(user)

	user = &User{Id:12, Name:"silence"}
	fmt.Println(ormer.Update(user, "Name"))
	fmt.Println(user)
}