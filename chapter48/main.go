package main

import (
	"time"
	"fmt"
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

	ormer := orm.NewOrm()

	queryset := ormer.QueryTable(&User{})
	fmt.Println(queryset.Count())
	fmt.Println(queryset.Filter("profile__algo", "sha256").Count())
	fmt.Println(queryset.Filter("profile__algo", "none").Count())
	fmt.Println(queryset.Filter("posts__title__contains", "beego").Count())
	fmt.Println(queryset.Filter("posts__title__icontains", "beego").Count())

	queryset = ormer.QueryTable(&Profile{})
	fmt.Println(queryset.Count())
	fmt.Println(queryset.Filter("user__name", "kk").Count())
	fmt.Println(queryset.Filter("user__name", "silence").Count())


	queryset = ormer.QueryTable(&Post{})
	fmt.Println(queryset.Count())
	fmt.Println(queryset.Filter("user__name", "kk").Count())
	fmt.Println(queryset.Filter("user__name", "silence").Count())


	queryset = ormer.QueryTable(&Post{})
	fmt.Println(queryset.Count())
	fmt.Println(queryset.Filter("tags__tag__name__icontains", "golang").Count())
	fmt.Println(queryset.Filter("tags__tag__name__icontains", "web").Count())
	fmt.Println(queryset.Filter("tags__tag__name__contains", "golang").Count())


	queryset = ormer.QueryTable(&Tag{})
	fmt.Println(queryset.Count())
	fmt.Println(queryset.Filter("posts__post__title__contains", "beego").Count())
	fmt.Println(queryset.Filter("posts__post__title__icontains", "beego").Count())
}