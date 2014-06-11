package controllers

import (
	"exbook/models"
	"strings"
	"time"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Login() {
	uid := this.GetSession("uid")
	if uid == nil {
		this.Data["PageTitle"] = "用户登录"
		this.TplNames = "login.html"
	} else {
		this.Redirect("/", 302)
	}

}

func (this *LoginController) DoLogin() {
	name := this.GetString("uname")

	password := this.GetString("pass")
	//this.Ctx.WriteString(password)
	remember := this.GetString("remember")
	if name != "" && password != "" {
		var userinfo models.User
		userinfo.Uname = name
		if userinfo.Read("uname") != nil || userinfo.Pwd != models.Md5([]byte(password)) {
			this.Data["errmsg"] = "帐号或密码错误!"
		} else {
			userinfo.Lastlogin = this.getTime()
			userinfo.Update()
			if remember == "yes" {
				this.Ctx.SetCookie("uname", userinfo.Uname, 7*86400, "/")
				this.SetSession("uname", userinfo.Uname)
				this.SetSession("uid", userinfo.Id)
				this.SetSession("region", userinfo.Region)
			} else {
				this.SetSession("uname", userinfo.Uname)
				this.SetSession("uid", userinfo.Id)
				this.SetSession("region", userinfo.Region)
			}

			this.Redirect("/", 302)
		}
	}
	this.Data["PageTitle"] = "用户登录"
	this.TplNames = "login.html"
}

func (this *LoginController) Logout() {
	this.DelSession("uid")
	this.DelSession("uname")
	this.DelSession("region")
	this.Redirect("/login", 302)
}

//获取用户IP地址
func (this *LoginController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

func (this *LoginController) getTime() time.Time {
	return time.Now().UTC()
}

func (this *LoginController) CheckLogin() {
	if this.IsLogin {
		this.Ctx.WriteString("ok")
	} else {
		this.Ctx.WriteString("fail")
	}
}
