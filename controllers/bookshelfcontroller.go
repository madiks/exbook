package controllers

import (
	"exbook/models"
	"fmt"
	"strconv"
	"time"
)

type BookShelfController struct {
	BaseAController
}

func (this *BookShelfController) Get() {
	var list []*models.Inventory
	var Isempty bool
	query := new(models.Inventory).Query().Filter("uid", this.Uid)
	count, _ := query.Count()
	if count > 0 {
		Isempty = false
		query.OrderBy("-update_time").All(&list)
	} else {
		Isempty = true
	}
	this.Data["isempty"] = Isempty
	//fmt.Println(list)
	type bookshe struct {
		Id         int64
		Uid        int64
		Bid        int64
		NOrate     string
		Status     int64
		Updatetime time.Time
		Createtime time.Time
		Title      string
		Subtitle   string
		Author     string
		Image      string
		Dbalt      string
	}
	var Bks []bookshe
	for k, _ := range list {
		var bse bookshe
		bse.Id = list[k].Id
		bse.Uid = list[k].Uid
		bse.Bid = list[k].Bid
		bse.NOrate = list[k].Newoldrate
		bse.Updatetime = list[k].UpdateTime
		bse.Createtime = list[k].CreateTime
		bse.Status = list[k].Status
		var bk models.Books
		bk.Id = list[k].Bid
		bk.Read()
		bse.Title = bk.Title
		bse.Subtitle = bk.Subtitle
		bse.Author = bk.Author
		bse.Image = bk.Image
		bse.Dbalt = bk.Dbalt
		Bks = append(Bks, bse)
	}
	fmt.Print(Bks)
	this.Data["Bks"] = Bks
	this.Data["PageTitle"] = "我的书架"
	this.Data["IsBookShelf"] = true
	this.Layout = "layout/default.html"
	this.TplNames = "bookshelf.html"
}

func (this *BookShelfController) GetInv() {
	if this.IsLogin {
		var list []*models.Inventory
		query := new(models.Inventory).Query().Filter("uid", this.Uid)
		count, _ := query.Count()
		if count > 0 {
			query.OrderBy("-update_time").All(&list)
		}
		//fmt.Println(list)
		type bookshe struct {
			Id         int64
			Uid        int64
			Bid        int64
			NOrate     string
			Status     int64
			Updatetime time.Time
			Createtime time.Time
			Title      string
			Subtitle   string
			Author     string
			Image      string
			Dbalt      string
		}
		var Bks []bookshe
		str := ""
		for k, _ := range list {
			var bse bookshe
			sid := strconv.Itoa(int(list[k].Id))
			bse.Id = list[k].Id
			bse.Uid = list[k].Uid
			bse.Bid = list[k].Bid
			bse.NOrate = list[k].Newoldrate
			bse.Updatetime = list[k].UpdateTime
			bse.Createtime = list[k].CreateTime
			bse.Status = list[k].Status
			var bk models.Books
			bk.Id = list[k].Bid
			bk.Read()
			bse.Title = bk.Title
			bse.Subtitle = bk.Subtitle
			bse.Author = bk.Author
			bse.Image = bk.Image
			bse.Dbalt = bk.Dbalt
			Bks = append(Bks, bse)
			var stri = "<div class='col-md-2 column productbox' id='ofgooddiv" + sid + "'><img src='" + bse.Image + "' class='img-responsive' style='height:110px;weight:80px'><div class='producttitle'>" + bse.Title + "</div><div class='productprice'><div class='pull-right'><button actname='" + sid + "' class='btn btn-info btn-sm ofgood' role='button'>添加</button></div><div class='pricetext'>" + bse.NOrate + "</div></div></div>"
			str = str + stri
		}
		fmt.Print(Bks)

		this.Data["json"] = str
	} else {
		this.Data["json"] = "you are not login!"
	}
	this.ServeJson()
}
