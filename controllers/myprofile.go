package controllers

import (
	"exbook/models"
)

type MyProfileController struct {
	BaseAController
}

func (this *MyProfileController) Get() {
	var usr models.User
	usr.Id = this.Uid
	usr.Read()
	type Regi struct {
		Id     int64
		Region string
		Isthis bool
	}

	var rlist []*models.Regionset
	query := new(models.Regionset).Query()
	count, _ := query.Count()
	if count > 0 {
		query.All(&rlist)
	}
	var reglist []Regi
	for k, _ := range rlist {
		var r1 Regi
		r1.Id = rlist[k].Id
		r1.Region = rlist[k].Region
		if rlist[k].Region == usr.Region {
			r1.Isthis = true
		} else {
			r1.Isthis = false
		}
		reglist = append(reglist, r1)
	}
	this.Data["uinfo"] = usr
	this.Data["rglist"] = reglist
	this.Data["PageTitle"] = "我的资料"
	this.Layout = "layout/default.html"
	this.TplNames = "myprofile.html"
}

func (this *MyProfileController) UpdateProfile() {
	tname := this.GetString("tname")
	tel := this.GetString("tel")
	address := this.GetString("address")
	//region := this.GetString("region")
	email := this.GetString("email")
	pf := models.User{Id: this.Uid}
	if pf.Read() == nil {
		pf.Tel = tel
		pf.Tname = tname
		pf.Address = address
		//pf.Region = region
		pf.Email = email
		pf.Update()
	}
	this.Redirect("/myprofile", 302)
}

func (this *MyProfileController) ChangePwd() {
	pwd := this.GetString("pwd")
	pwds := this.GetString("pwds")
	if pwd == pwds {
		enpwd := models.Md5([]byte(pwd))
		pf := models.User{Id: this.Uid}
		if pf.Read() == nil {
			pf.Pwd = enpwd
			pf.Update()
		}
	}
	this.DelSession("uid")
	this.DelSession("uname")
	this.DelSession("region")
	this.Redirect("/login", 302)

}
