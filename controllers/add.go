package controllers

import (
	"exbook/models"
	"strconv"
	"time"
)

type AddController struct {
	BaseAController
}

func (this *AddController) AddTrade() {
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
	this.Data["Bks"] = Bks
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
	this.Data["PageTitle"] = "添加交换"
	this.Data["IsAddTrade"] = true
	this.Layout = "layout/default.html"
	this.TplNames = "addtrade.html"
}

func (this *AddController) AddBook() {
	this.Data["PageTitle"] = "我的报价"
	this.Data["IsMyOffer"] = true
	this.Layout = "layout/default.html"
	this.TplNames = "myoffer.html"
}

func (this *AddController) getTime() time.Time {
	return time.Now().UTC()
}

func (this *AddController) AddToLeft() {
	ivid, _ := strconv.Atoi(this.GetString("Ivid"))
	var list models.Inventory
	query := new(models.Inventory).Query().Filter("uid", this.Uid).Filter("id", ivid)
	count, _ := query.Count()
	if count > 0 {
		query.One(&list)
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
	var bse bookshe
	bse.Id = list.Id
	bse.Uid = list.Uid
	bse.Bid = list.Bid
	bse.NOrate = list.Newoldrate
	bse.Updatetime = list.UpdateTime
	bse.Createtime = list.CreateTime
	bse.Status = list.Status
	var bk models.Books
	bk.Id = list.Bid
	bk.Read()
	bse.Title = bk.Title
	bse.Subtitle = bk.Subtitle
	bse.Author = bk.Author
	bse.Image = bk.Image
	bse.Dbalt = bk.Dbalt
	this.Data["json"] = bse
	this.ServeJson()
}

func (this *AddController) AddToRight() {
	bid, _ := strconv.Atoi(this.GetString("Bid"))
	var bk models.Books
	bk.Id = int64(bid)
	bk.Read()
	this.Data["json"] = bk
	this.ServeJson()
}

func (this *AddController) AddOffer() {
	tid, _ := strconv.Atoi(this.GetString("Tid"))
	tcid, _ := strconv.Atoi(this.GetString("Tcid"))
	offerl := this.GetString("Offerl")
	remark := this.GetString("Remark")
	ohasmoney, _ := strconv.Atoi(this.GetString("Ohasmoney"))
	omoney, _ := strconv.Atoi(this.GetString("Omoney"))
	var of models.Offers
	of.Uid = this.Uid
	of.Uname = this.Uname
	of.Offerlist = offerl
	of.Msg = remark
	of.Ohasmoney = int64(ohasmoney)
	of.Omoney = int64(omoney)
	of.Tid = int64(tid)
	of.Tcid = int64(tcid)
	of.Tstatus = 1
	of.Offerstatus = 1
	of.CreateTime = this.getTime()
	of.Insert()
	td := models.Trades{Id: int64(tid)}
	if td.Read() == nil {
		td.UpdateTime = this.getTime()
		td.Update()
	}
	this.Data["json"] = "ok"
	this.ServeJson()
}
