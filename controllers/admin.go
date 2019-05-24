package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
)

type AdminController struct {
	beego.Controller
}

type res struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}


func (this *AdminController) Get() {
	// 管理员名称
	adminName := this.GetSession("name")

	this.TplName = "admin/index.html"
	this.Data["Title"] = "首页"
	this.Data["Admin"] = adminName
	this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())

}

func (this *AdminController) Logout() {
	this.DelSession("uid")
	this.DelSession("name")

	this.Data["json"] = res{200,"登出成功"}
	this.ServeJSON()
}
