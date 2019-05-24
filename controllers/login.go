package controllers

import (
	"beego-blog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"html/template"
	"reflect"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get(){
	// 解析 flash 闪存的数据
	_ = beego.ReadFromRequest(&this.Controller)

	this.TplName = "admin/login.html"
	this.Data["Title"] = "后台登录"
	this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())
}

func (this *LoginController) Post() {

	// 处理错误异常
	defer func() {
		if err := recover(); err != nil {
			flash := beego.NewFlash()
			v := reflect.ValueOf(err)
			flash.Error(v.String())
			flash.Store(&this.Controller)
			this.Redirect("/admin/login",302)
			return
		}
	}()

	// 账号和密码的验证
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	// 新建验证
	valid := validation.Validation{}
	valid.Required(username,"username").Message("用户名不能为空")
	valid.Required(password,"password").Message("密码不能为空")

	// 如果有错误信息，将错误信息flash起来
	if valid.HasErrors() {
		var err_msg string
		for _, err := range valid.Errors {
			err_msg += err.Message+"\r\n"
		}

		panic(err_msg)
		return
	}
	// 如果没错，开始验证账号密码

	var admin models.Admin
	admin.GetByName(username)

	// 检查账号密码是否正确，不正确就抛出异常
	if res := admin.CheckPwd(password); !res{
		panic("账号或者密码错误")
		return
	}

	this.SetSession("uid",admin.Id)
	this.SetSession("name",admin.Username)
	this.Redirect("/admin",302)
}