package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"html/template"
	"log"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get(){
	this.TplName = "login.html"
	this.Data["Title"] = "后台登录"
	this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())
}

func (this *LoginController) Post() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	valid := validation.Validation{}
	valid.Required(username,"用户名")
	valid.Required(password,"密码")
	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	this.Ctx.WriteString(username)
}