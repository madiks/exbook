package controllers

import (
	"exbook/models"
	"strconv"
	"time"
)

type AddFavController struct {
	BaseController
}

func (this *AddFavController) AddTrade() {
	tid, _ := strconv.Atoi(this.GetString("Tid"))
	result := struct {
		Existed string `json:"existed"`
	}{"0"}
	if this.IsLogin {
		var bk models.Bookmark
		c_query := bk.Query()
		c_count, _ := c_query.Filter("tid", tid).Filter("uid", this.Uid).Filter("type", 1).Count()

		if c_count > 0 {
			result.Existed = "1"
		} else {
			result.Existed = "0"
			bk.Uid = this.Uid
			bk.Type = 1
			bk.Tid = int64(tid)
			bk.CreateTime = this.getTime()
			bk.Status = 1
			bk.Insert()
		}
	} else {
		result.Existed = "2"
	}
	this.Data["json"] = result
	this.ServeJson()
}

func (this *AddFavController) AddBook() {
	bid, _ := strconv.Atoi(this.GetString("Bid"))
	result := struct {
		Existed string `json:"existed"`
	}{"0"}
	if this.IsLogin {

		var bk models.Bookmark
		c_query := bk.Query()
		c_count, _ := c_query.Filter("bid", bid).Filter("uid", this.Uid).Filter("type", 2).Count()

		if c_count > 0 {
			result.Existed = "1"
		} else {
			result.Existed = "0"
			bk.Uid = this.Uid
			bk.Type = 2
			bk.Bid = int64(bid)
			bk.CreateTime = this.getTime()
			bk.Status = 1
			bk.Insert()
		}
	} else {
		result.Existed = "2"
	}
	this.Data["json"] = result
	this.ServeJson()
}

func (this *AddFavController) getTime() time.Time {
	return time.Now().UTC()
}

func (this *AddFavController) RmBook() {
	var id int64
	var err error
	if id, err = strconv.ParseInt(this.Ctx.Input.Param(":invid"), 10, 64); err != nil {

	} else {
		inv := models.Inventory{Id: id}
		if inv.Read() == nil {
			//需删除交易中包含该书的
			inv.Delete()
		}
		this.Redirect("/bookshelf", 302)
	}
}
