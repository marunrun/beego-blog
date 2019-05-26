package controllers

import (
	"beego-blog/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"html/template"
)

type CategoryController struct {
	BaseController
}

// 分类管理列表
func (this *CategoryController) Get() {
	this.TplName = "admin/category.html"
	this.Data["Title"] = "分类管理"
	this.Data["xsrftoken"] = template.HTML(this.XSRFToken())
	var list  []models.Category
	o := orm.NewOrm()
	o.QueryTable(new(models.Category)).All(&list)

	this.Data["categorys"] = list


}

// 添加分类
func (this *CategoryController) Post() {
	title := this.Input().Get("title")
	fmt.Println(title)
	if title == ""{
		this.Data["json"] = res{400,"分类名称不能为空"}
		this.ServeJSON()
		return
	}
	var category models.Category
	category.Title = title
	o := orm.NewOrm()
	_, err := o.Insert(&category)
	if err != nil{
		this.Data["json"] = res{400,err.Error()}
		this.ServeJSON()
		return
	}
	this.Data["json"] = res{0,"添加成功"}
	this.ServeJSON()
	return
}