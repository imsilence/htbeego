package main

import (
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int
	Name string `orm:"size(64)"`
	Tel string `orm:"size(32)"`
	Addr string `orm:"size(1024)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Update time.Time `orm:"auto_now;type(datetime)"`
}

type Log struct {
	Id int
	Log string `orm:"type(text)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func main() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:881019@tcp(127.0.0.1:3306)/htbeego?charset=utf8mb4&loc=Asia%2FShanghai")
	orm.RegisterDataBase("log", "mysql", "root:881019@tcp(127.0.0.1:3306)/htlog?charset=utf8mb4&loc=Asia%2FShanghai")
	orm.RegisterModel(&User{}, &Log{})
	orm.RunSyncdb("default", true, true); orm.RunSyncdb("log", true, true)

	ormer := orm.NewOrm()
	// 写入默认库
	ormer.Insert(&User{Name: "kk", Tel: "15200000000", Addr: "西安市"})
	// 写入log库
	ormer.Using("log")
	ormer.Insert(&Log{Log:"测试多数据库写入"})
}