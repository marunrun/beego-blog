package main

import (
	"beego-blog/models"
	_ "beego-blog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

)

func init()  {
	models.RegisterDB()
}

func main() {
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	beego.Run()
}
