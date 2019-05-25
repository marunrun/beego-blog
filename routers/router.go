package routers

import (
	"beego-blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin/login", &controllers.LoginController{})
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/admin/logout",&controllers.AdminController{},"post:Logout")
	beego.Router("/admin/user",&controllers.AdminController{},"get:User;put:ChangePwd")
	beego.Router("/admin/category",&controllers.CategoryController{})
}
