package main

import (
	_ "beegorm/docs"
	_ "beegorm/routers"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}