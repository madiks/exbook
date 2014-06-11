package controllers

import (
	"exbook/models"
	"strings"
	"time"
)

type RegisterController struct {
	BaseController
}

func (this *RegisterController) Reg() {
	uid := this.GetSession("uid")
	if uid == nil {
		var rlist []*models.Regionset
		query := new(models.Regionset).Query()
		count, _ := query.Count()
		if count > 0 {
			query.All(&rlist)
		}
		this.Data["rlist"] = rlist
		this.Data["PageTitle"] = "用户注册"
		this.TplNames = "register.html"
	} else {
		this.Redirect("/", 302)
	}

}

func (this *RegisterController) DoReg() {
	name := this.GetString("uname")

	password := this.GetString("password")
	//this.Ctx.WriteString(password)
	region := this.GetString("address")
	if name != "" && password != "" {
		var userinfo models.User
		userinfo.Uname = name
		userinfo.Pwd = models.Md5([]byte(password))
		userinfo.Region = region
		flag := userinfo.Insert()
		if flag == nil {
			this.SetSession("uname", userinfo.Uname)
			this.SetSession("uid", userinfo.Id)
			this.SetSession("region", userinfo.Region)
			this.Redirect("/", 302)
		}
	}
	this.Data["PageTitle"] = "用户注册"
	this.TplNames = "register.html"
}

//获取用户IP地址
func (this *RegisterController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

func (this *RegisterController) getTime() time.Time {
	return time.Now().UTC()
}

func (this *RegisterController) Checkuname() {
	name := this.GetString("chkvalue")
	var usera models.User
	c_query := usera.Query()
	c_count, _ := c_query.Filter("uname", name).Count()
	result := struct {
		Existed bool `json:"existed"`
	}{true}

	if c_count > 0 {
		result.Existed = true
	} else {
		result.Existed = false
	}
	this.Data["json"] = result
	this.ServeJson()
	//this.StopRun()
}
