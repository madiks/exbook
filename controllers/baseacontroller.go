package controllers

import (
	"exbook/models"
	"github.com/astaxie/beego"
)

type BaseAController struct {
	beego.Controller
	IsLogin     bool
	Uname       string
	Uid         int64
	Region      string
	Newoffernum int64
	NewRpnum    int64
	Showofnum   bool
	Showrpnum   bool
}

func (this *BaseAController) Prepare() {
	this.Data["Favicon"] = beego.AppConfig.String("favicon")
	this.Data["SiteTitle"] = beego.AppConfig.String("site_title")
	this.Data["AdminName"] = beego.AppConfig.String("admin_name")
	this.Data["AdminEmail"] = beego.AppConfig.String("admin_email")
	this.Data["AdminLink"] = beego.AppConfig.String("admin_link")
	this.CheckLogin()
	this.GetNewOfferNum()
	this.GetNewReplyNum()
	this.Data["NewOfferNum"] = this.Newoffernum
	this.Data["NewRpNum"] = this.NewRpnum
	this.Data["Showofnum"] = this.Showofnum
	this.Data["Showrpnum"] = this.Showrpnum
	this.Data["IsLogin"] = this.IsLogin
	this.Data["Uname"] = this.Uname
	this.Data["Uid"] = this.Uid
	this.Data["Region"] = this.Region
}

func (this *BaseAController) CheckLogin() {
	uid := this.GetSession("uid")
	if uid == nil {
		this.IsLogin = false
		this.Data["IsLogin"] = this.IsLogin
		this.Region = "长春"
		this.Data["Region"] = this.Region
		this.Redirect("/login", 302)
	} else {
		this.IsLogin = true
		this.Uname = this.GetSession("uname").(string)
		this.Uid = this.GetSession("uid").(int64)
		this.Region = this.GetSession("region").(string)
	}
}

func (this *BaseAController) GetNewOfferNum() {
	queryo := new(models.Offers).Query().Filter("tcid", this.Uid).Filter("offerstatus", 1)
	counto, _ := queryo.Count()
	if counto > 0 {
		this.Showofnum = true
		this.Newoffernum = counto
	} else {
		this.Showofnum = false
	}
}

func (this *BaseAController) GetNewReplyNum() {
	query := new(models.Reply_offer).Query().Filter("toid", this.Uid).Filter("status", 1)
	count, _ := query.Count()
	if count > 0 {
		this.NewRpnum = count
		this.Showrpnum = true
	} else {
		this.Showrpnum = false
	}
}
