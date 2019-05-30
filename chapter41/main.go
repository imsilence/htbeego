package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:881019@tcp(127.0.0.1:3306)/htbeego?charset=utf8mb4&loc=Asia%2FShanghai")

	ormer := orm.NewOrm()

	builder, _ := orm.NewQueryBuilder("mysql")
	builder = builder.Select("id", "name")
	builder = builder.From("user")
	builder = builder.Where("name like ?")
	builder = builder.OrderBy("id").Desc()
	builder = builder.Limit(3).Offset(2)

	fmt.Println(builder.String())
	rawset := ormer.Raw(builder.String(), "%kk%")

	var values []orm.Params
	rawset.Values(&values)
	for _, v := range values {
		fmt.Printf("%#v\n", v)
	}

	builder, _ = orm.NewQueryBuilder("mysql")
	builder = builder.Select("name", "count(*) as cnt")
	builder = builder.From("user")
	builder = builder.GroupBy("name")
	builder = builder.Having("cnt >= ?")

	fmt.Println(builder.String())
	rawset = ormer.Raw(builder.String(), 6)

	rawset.Values(&values)
	for _, v := range values {
		fmt.Printf("%#v\n", v)
	}
}