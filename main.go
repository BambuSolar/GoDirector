package main

import (
		_ "github.com/BambuSolar/GoDirector/routers"
		"github.com/astaxie/beego"
		"github.com/astaxie/beego/orm"
		_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego/plugins/cors"
	"os"
	"fmt"
)


func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	fmt.Print(os.Getenv("MySQLConection"))
	orm.RegisterDataBase("default", "mysql", os.Getenv("MySQLConection"))
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

