package main

import (
	_ "wleirock/data-distribute/routers"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(192.168.1.93:3306)/test?charset=utf8")
}

func main() {
	beego.Run()
}
