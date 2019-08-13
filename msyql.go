package main

import (
	_ "github.com/go-sql-driver/mysql" //mysql driver
	"github.com/ilibs/gosql"
)

func main() {
	configs := make(map[string]*gosql.Config)

	configs["default"] = &gosql.Config{
		Enable:  true,
		Driver:  "mysql",
		Dsn:     "root:ztinfo2017@tcp(192.168.1.46:3308)/test?charset=utf8&parseTime=True&loc=Asia%2FShanghai",
		ShowSql: true,
	}

	//connection database
	gosql.Connect(configs)

	gosql.DB().QueryRowx("select * from users where id = 1")
}
