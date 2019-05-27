package controllers

import (
	"beego-blog/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"html/template"
	"strconv"
)

type CategoryController struct {
	BaseController
}

// 分类管理列表
func (this *CategoryController) Get() {
	var page int
	var err error
	if res := this.Input().Get("page"); res == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(res)
		if err != nil {
			panic(err)
		}
	}

	if page < 1 {
		page = 1
	}
	this.TplName = "admin/category.html"
	this.Data["Title"] = "分类管理"
	this.Data["xsrftoken"] = template.HTML(this.XSRFToken())
	var list []models.Category
	o := orm.NewOrm()
	o.QueryTable(new(models.Category)).OrderBy("Created").Limit(1).Offset(1 * (page - 1)).All(&list)

	total, _ := o.QueryTable(new(models.Category)).Count()

	pageHtml := this.paginate(page, total, 1)
	this.Data["pageHtml"] = template.HTML(pageHtml)
	this.Data["categorys"] = list
}

// 添加分类
func (this *CategoryController) Post() {
	title := this.Input().Get("title")
	fmt.Println(title)
	if title == "" {
		this.Data["json"] = res{400, "分类名称不能为空"}
		this.ServeJSON()
		return
	}
	var category models.Category
	category.Title = title
	o := orm.NewOrm()
	_, err := o.Insert(&category)
	if err != nil {
		this.Data["json"] = res{400, err.Error()}
		this.ServeJSON()
		return
	}
	this.Data["json"] = res{0, "添加成功"}
	this.ServeJSON()
	return
}

// 删除分类
func (this *CategoryController) Delete() {
	id, err := this.GetInt("id")

	if err != nil {
		this.Data["json"] = res{400, err.Error()}
		this.ServeJSON()
		return
	}

	o := orm.NewOrm()
	if _, err := o.Delete(&models.Category{Id: id}); err != nil {
		this.Data["json"] = res{400, err.Error()}
		this.ServeJSON()
		return
	}

	this.Data["json"] = res{0, "删除成功"}
	this.ServeJSON()
	return
}

// 更新分类
func (this *CategoryController) Put() {
	params := this.Input()
	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		this.Data["json"] = res{400, err.Error()}
		this.ServeJSON()
		return
	}

	o := orm.NewOrm()
	category := models.Category{Id:id,Title:params.Get("title")}

	_, err = o.Update(&category, "Title")
	if err != nil {
		this.Data["json"] = res{400, err.Error()}
		this.ServeJSON()
		return
	}

	this.Data["json"] = res{0, "修改成功"}
	this.ServeJSON()
	return
}
