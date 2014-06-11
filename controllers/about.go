package controllers

import ()

type AboutController struct {
	BaseController
}

func (this *AboutController) Get() {
	this.Data["PageTitle"] = "关于网站"
	this.Data["IsAbout"] = true
	this.Layout = "layout/default.html"
	this.TplNames = "about.html"
}
