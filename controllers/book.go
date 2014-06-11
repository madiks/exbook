package controllers

import ()

type BookController struct {
	BaseAController
}

func (this *BookController) Get() {
	this.Data["PageTitle"] = "添加书籍"
	this.Layout = "layout/default.html"
	this.TplNames = "addbook.html"
}
