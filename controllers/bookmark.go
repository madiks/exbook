package controllers

import (
	"encoding/json"
	"exbook/models"
	"fmt"
	"strconv"
	"time"
)

type BookMarkController struct {
	BaseAController
}

func (this *BookMarkController) Get() {
	//收藏的交易
	var listt []*models.Bookmark
	query := new(models.Bookmark).Query().Filter("uid", this.Uid).Filter("type", 1)
	count, _ := query.Count()
	if count > 0 {
		query.OrderBy("-create_time").All(&listt)
	}
	var list []models.Trades
	var err error
	for k, _ := range listt {
		var l models.Trades
		l.Id = listt[k].Tid
		l.Read()
		list = append(list, l)
	}
	type Booksinfo struct {
		Bid    int64
		Ivid   int64
		Bname  string
		NOrate string
		Bimg   string
	}

	type Tinfos struct {
		Id          int64
		Uid         int64
		Uname       string
		Sells       []Booksinfo
		Wants       []Booksinfo
		Hasanyoffer bool
		SHasmoney   bool
		SMoney      int64
		WHasmoney   bool
		WMoney      int64
		CreateTime  time.Time
		Region      string
		UpdateTime  time.Time
		Remark      string
	}
	var tlist []Tinfos
	for k, _ := range list {
		var tin Tinfos
		tin.Id = list[k].Id
		tin.Uid = list[k].Uid
		tin.Uname = list[k].Uname
		tin.CreateTime = list[k].CreateTime
		tin.UpdateTime = list[k].UpdateTime
		tin.Region = list[k].Region
		tin.Remark = list[k].Remark
		if list[k].Hasanyoffer == 0 {
			tin.Hasanyoffer = false
		} else {
			tin.Hasanyoffer = true
		}
		if list[k].Shasmoney == 0 {
			tin.SHasmoney = false
		} else {
			tin.SHasmoney = true
			tin.SMoney = list[k].Smoney
		}
		if list[k].Whasmoney == 0 {
			tin.WHasmoney = false
		} else {
			tin.WHasmoney = true
			tin.WMoney = list[k].Wmoney
		}
		if list[k].Selllist != "" {
			var j1 []int
			err = json.Unmarshal(models.DecodeSS([]byte(list[k].Selllist)), &j1)
			if err != nil {
				panic(err)
			}
			var bis []Booksinfo
			for k1, _ := range j1 {
				var inv models.Inventory
				inv.Id = int64(j1[k1])
				inv.Read()
				if inv.Newoldrate != "" {
					//fmt.Print(inv)
					var bok models.Books
					bok.Id = inv.Bid
					bok.Read()
					var bi Booksinfo
					bi.Bid = inv.Bid
					bi.Bimg = bok.Image
					bi.Bname = bok.Title
					bi.Ivid = inv.Id
					bi.NOrate = inv.Newoldrate
					//fmt.Print(bi)
					//result[year] = append(result[year], v)
					bis = append(bis, bi)
				}
			}
			tin.Sells = bis
		}
		if list[k].Wantlist != "" {
			var j2 []int
			err = json.Unmarshal(models.DecodeSS([]byte(list[k].Wantlist)), &j2)
			if err != nil {
				panic(err)
			}
			var bis []Booksinfo
			for k2, _ := range j2 {
				var bok models.Books
				bok.Id = int64(j2[k2])
				bok.Read()
				var bi Booksinfo
				bi.Bid = bok.Id
				bi.Bimg = bok.Image
				bi.Bname = bok.Title
				bi.Ivid = bok.Id
				//fmt.Print(bi)
				//result[year] = append(result[year], v)
				bis = append(bis, bi)
			}
			tin.Wants = bis
		}
		tlist = append(tlist, tin)
	}
	fmt.Print(tlist)
	//fmt.Println("1111")
	//收藏的书籍
	var listb []*models.Bookmark
	querys := new(models.Bookmark).Query().Filter("uid", this.Uid).Filter("type", 2)
	counts, _ := querys.Count()
	if counts > 0 {
		querys.OrderBy("-create_time").All(&listb)
	}
	type bookt struct {
		Id         int64
		Uid        int64
		Bid        int64
		Createtime time.Time
		Title      string
		Subtitle   string
		Author     string
		Image      string
		Dbalt      string
	}
	var bks []bookt
	for k, _ := range listb {
		var bkm bookt
		bkm.Id = listb[k].Id
		bkm.Uid = listb[k].Uid
		bkm.Bid = listb[k].Bid
		bkm.Createtime = listb[k].CreateTime
		var bk models.Books
		bk.Id = listb[k].Bid
		bk.Read()
		bkm.Title = bk.Title
		bkm.Subtitle = bk.Subtitle
		bkm.Author = bk.Author
		bkm.Image = bk.Image
		bkm.Dbalt = bk.Dbalt
		bks = append(bks, bkm)
	}
	this.Data["bks"] = bks
	this.Data["bts"] = tlist
	this.Data["PageTitle"] = "我的收藏"
	this.Data["IsMyBookMark"] = true
	this.Layout = "layout/default.html"
	this.TplNames = "bookmark.html"
}

func (this *BookMarkController) RemoveBT() {
	var id int64
	var err error
	if id, err = strconv.ParseInt(this.Ctx.Input.Param(":tid"), 10, 64); err != nil {

	} else {
		new(models.Bookmark).Query().Filter("uid", this.Uid).Filter("tid", id).Delete()
		this.Redirect("/mybookmark", 302)
	}
}

func (this *BookMarkController) RemoveBK() {
	var id int64
	var err error
	if id, err = strconv.ParseInt(this.Ctx.Input.Param(":bid"), 10, 64); err != nil {

	} else {
		inv := models.Bookmark{Id: id}
		if inv.Read() == nil {
			//需删除交易中包含该书的
			inv.Delete()
		}
		this.Redirect("/mybookmark", 302)
	}
}
