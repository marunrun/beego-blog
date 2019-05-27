package controllers

import (
	"beego-blog/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"html/template"
)

type AdminController struct {
	BaseController
}

type res struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

// 后台首页GET
func (this *AdminController) Get() {
	this.TplName = "admin/home.html"
	this.Data["Title"] = "首页"
	this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())
	this.Data["xsrftoken"] = template.HTML(this.XSRFToken())

}

// 退出登陆
func (this *AdminController) Logout() {
	this.DelSession("uid")
	this.DelSession("name")

	this.Data["json"] = res{200,"登出成功"}
	this.ServeJSON()
}

// 修改密码的页面 GET 请求
func (this *AdminController) User() {
	this.TplName = "admin/changePwd.html"
	this.Data["Title"] = "修改密码"
	this.Data["xsrftoken"] = template.HTML(this.XSRFToken())
}

// 修改密码 PUT 请求
func (this *AdminController) ChangePwd() {
	// 获取当前登陆的用户id
	uid := this.GetSession("uid")
	if uid == nil {
		this.Data["json"] = res{500,"登陆已过期，请重新登陆"}
		this.ServeJSON()
		return
	}


	old_pwd := this.Input().Get("old_pwd")
	new_pwd := this.Input().Get("new_pwd")
	// 新建验证是否传了原始密码和新密码
	valid := validation.Validation{}
	valid.Required(old_pwd,"old_pwd").Message("原始密码不能为空")
	valid.Required(new_pwd,"new_pwd").Message("新密码不能为空")
	if valid.HasErrors() {
		var err_msg string
		for _, err := range valid.Errors {
			err_msg += err.Message+"\r\n"
		}
		this.Data["json"] = res{400,err_msg}
		this.ServeJSON()
		return
	}

	//验证原始密码是否正确
	admin := models.Admin{Id:uid.(int)}
	o := orm.NewOrm()
	err := o.Read(&admin)

	if err != nil || !admin.CheckPwd(old_pwd){
		this.Data["json"] = res{400,"原始密码错误"}
		this.ServeJSON()
		return
	}

	admin.Password = admin.Encode(new_pwd)

	_, err = o.Update(&admin,"password")
	if err != nil {
		this.Data["json"] = res{400,"系统错误，请重试"}
		this.ServeJSON()
		return
	}
	this.DelSession("uid")
	this.Data["json"] = res{0,"修改成功"}
	this.ServeJSON()
	return
}

func (this *AdminController) Upload() {
	params := this.Input()
	fmt.Println(params)
}