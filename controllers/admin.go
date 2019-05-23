package controllers

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get() {
	if err := this.GetSession("uid"); err == nil {
		this.Redirect("/admin/login",302)
		return
	}

	this.Ctx.WriteString("Admin")
}

