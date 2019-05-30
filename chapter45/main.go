package main

import (
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/*
User 1:1 Profile
User 1:n Post
Post m:n Tag
*/

type User struct {
	Id int
	Name string `orm:"size(64)"`
	Tel string `orm:"size(32)"`
	Addr string `orm:"size(1024)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Update time.Time `orm:"auto_now;type(datetime)"`

	Profile *Profile `orm:"reverse(one)"`
	Posts []*Post `orm:"reverse(many)"`
}

type Profile struct {
	Id int
	Password string `orm:"size(1024)"`
	Salt string `orm:"size(32)"`
	Algo string `orm:"size(16)"`

	User *User `orm:"rel(one)"`
}

type Post struct {
	Id int
	Title string `orm:"size(128)"`
	Content string `orm:"type(text)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`

	User *User `orm:"rel(fk)"`
	Tags []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id int
	Name string `orm:"size(64)"`

	Posts []*Post `orm:"reverse(many)"`
}

func main() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:881019@tcp(127.0.0.1:3306)/htbeego?charset=utf8mb4&loc=Asia%2FShanghai")

	orm.RegisterModel(&User{}, &Profile{}, &Post{}, &Tag{})

	orm.RunSyncdb("default", true, true)

	ormer := orm.NewOrm()

	u := &User {
		Name: "kk",
		Tel: "15200000000",
		Addr: "西安市",
	}
	ormer.Insert(u)

	p := &Profile{
		Password : "123qwe!@#",
		Salt : "",
		Algo : "none",
		User: u,
	}
	ormer.Insert(p)

	tg := &Tag{Name:"Golang"}
	tw := &Tag{Name:"WEB"}
	tc := &Tag{Name:"并发"}
	ormer.Insert(tg); ormer.Insert(tw); ormer.Insert(tc)

	ptg := &Post{
		Title: "Goroutine",
		Content: "高并发使用Goroutine",
		User: u,
	}
	ormer.Insert(ptg); m2m := ormer.QueryM2M(ptg, "Tags"); m2m.Add(tg, tc);

	ptb := &Post{
		Title: "Beego",
		Content: "Beego入门",
		User: u,
	}
	ormer.Insert(ptb); m2m = ormer.QueryM2M(ptb, "Tags"); m2m.Add(tg, tw);
}