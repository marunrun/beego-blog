package controllers

type MainController struct {
	BaseController
}



func (this *MainController) Get() {
	this.TplName = "home.html"
	this.Data["Title"] = "首页"
}
