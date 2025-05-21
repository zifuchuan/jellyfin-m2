package main

import (
	"webapp/models"
	_ "webapp/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func initDb() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:mysql@123@tcp(127.0.0.1:3306)/testdb?charset=utf8")
	orm.RegisterModel(new(models.PyClass))
	orm.RunSyncdb("default", false, true)
}

func main() {
	initDb()
	beego.Run()
}
