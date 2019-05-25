package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}
var IsLogin bool = false

func (this *BaseController) Prepare() {


	if err := this.GetSession("uid"); err != nil {
		IsLogin = true
		// 管理员名称
		adminName := this.GetSession("name")
		this.Data["Admin"] = adminName
	}
	this.Data["IsLogin"] = IsLogin
}
