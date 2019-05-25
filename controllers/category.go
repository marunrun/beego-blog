package controllers

type CategoryController struct {
	BaseController
}

func (this *CategoryController) Get() {
	this.TplName = "admin/category.html"

}
