package main

import (
		_ "github.com/BambuSolar/GoDirector/routers"
		"github.com/astaxie/beego"
		"github.com/astaxie/beego/orm"
		_ "github.com/mattn/go-sqlite3"

	"github.com/astaxie/beego/plugins/cors"
)


func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "database/orm_test.db")
}

func main() {

	err := orm.RunSyncdb("default", false, false)
	if err != nil {
		beego.Error(err)
	}


	beego.InsertFilter("*", beego.BeforeRouter,cors.Allow(&cors.Options{
		AllowOrigins: []string{"https://*.foo.com"},
		AllowMethods: []string{"PUT", "PATCH"},
		AllowHeaders: []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
	}))

	beego.Run()
}

