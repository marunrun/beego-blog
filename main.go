package main

import (
	"beego-blog/models"
	_ "beego-blog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/context"
)

func init()  {
	// 注册ORM
	models.RegisterDB()

	// 过滤器，中间件，检查是否登录
	var FilterUser = func(ctx *context.Context) {
		_, ok := ctx.Input.Session("uid").(int)
		if !ok && ctx.Request.RequestURI != "/admin/login" {
			ctx.Redirect(302, "/admin/login")
		}
	}

	// 所有的admin路由都会执行中间件
	beego.InsertFilter("/admin/*",beego.BeforeRouter,FilterUser)

}

func main() {
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	beego.Run()
}
