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

	user := &User{Id:1}
	ormer.Read(user)

	fmt.Printf("%#v\n", *user)
	ormer.LoadRelated(user, "Profile")
	fmt.Printf("%#v\n", *user.Profile)
	ormer.LoadRelated(user, "Posts")
	for _, post := range user.Posts {
		fmt.Printf("%#v\n", *post)
	}

	profile := &Profile{Id:1}
	ormer.Read(profile)
	fmt.Printf("%#v\n", *profile)
	ormer.LoadRelated(profile, "User")
	fmt.Printf("%#v\n", *profile.User)

	post := &Post{Id:1}
	ormer.Read(post)
	fmt.Printf("%#v\n", *post)
	ormer.LoadRelated(post, "Tags")
	for _, tag := range post.Tags {
		fmt.Printf("%#v\n", *tag)
	}

	ormer.LoadRelated(post, "User")
	fmt.Printf("%#v\n", *post.User)

	tag := &Tag{Id:1}
	ormer.Read(tag)
	fmt.Printf("%#v\n", *tag)
	ormer.LoadRelated(tag, "Posts")
	for _, post := range tag.Posts {
		fmt.Printf("%#v\n", *post)
	}


}