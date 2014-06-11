package controllers

import (
	"encoding/json"
	"exbook/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type MyTradeController struct {
	BaseAController
}

func (this *MyTradeController) Get() {
	var list []*models.Trades
	var err error
	var Isempty bool
	query := new(models.Trades).Query().Filter("status", 1).Filter("uid", this.Uid)
	count, _ := query.Count()
	if count > 0 {
		Isempty = false
		query.OrderBy("-update_time").All(&list)
	} else {
		Isempty = true
	}
	this.Data["isempty"] = Isempty
	//fmt.Print(list)

	type Booksinfo struct {
		Bid    int64
		Ivid   int64
		Bname  string
		NOrate string
		Bimg   string
	}

	type Tinfos struct {
		Id           int64
		Uid          int64
		Uname        string
		Sells        []Booksinfo
		Wants        []Booksinfo
		Showpop      bool
		Newoffernum  int64
		Shownewoffer bool
		Hasanyoffer  bool
		SHasmoney    bool
		SMoney       int64
		WHasmoney    bool
		WMoney       int64
		CreateTime   time.Time
		Region       string
		UpdateTime   time.Time
		Remark       string
	}
	var tlist []Tinfos
	for k, _ := range list {
		var tin Tinfos
		tin.Id = list[k].Id
		tin.Uid = list[k].Uid
		tin.Uname = list[k].Uname
		tin.CreateTime = list[k].CreateTime
		tin.UpdateTime = list[k].UpdateTime
		queryof := new(models.Offers).Query().Filter("tcid", this.Uid).Filter("offerstatus", 1).Filter("tid", list[k].Id)
		countof, _ := queryof.Count()
		if countof > 0 {
			tin.Shownewoffer = true
			tin.Newoffernum = countof
		} else {
			tin.Shownewoffer = false
		}
		//该交易的更新时间与现在相差30分钟则显示pop按钮
		if time.Since(list[k].UpdateTime).Minutes() > 30 {
			tin.Showpop = true
		}
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
				//fmt.Print(inv)
				if inv.Newoldrate != "" {
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

	//fmt.Print(tlist)
	this.Data["trades"] = tlist
	this.Data["PageTitle"] = "我的交换"
	this.Data["IsMyTrade"] = true
	this.Layout = "layout/default.html"
	this.TplNames = "mytrade.html"
}

func (this *MyTradeController) RemoveTrade() {
	var id int64
	var err error
	if id, err = strconv.ParseInt(this.Ctx.Input.Param(":tid"), 10, 64); err != nil {

	} else {
		new(models.Trades).Query().Filter("uid", this.Uid).Filter("id", id).Delete()
		this.Redirect("/mytrade", 302)
	}
}

func (this *MyTradeController) PopTrade() {
	tid, _ := strconv.Atoi(this.GetString("tid"))
	result := struct {
		Existed string `json:"existed"`
	}{"0"}
	inv := models.Trades{Id: int64(tid)}
	if inv.Read() == nil {
		inv.UpdateTime = this.getTime()
		if err := inv.Update(); err == nil {
			result.Existed = "1"
		}
	}
	this.Data["json"] = result
	this.ServeJson()
}

func (this *MyTradeController) getTime() time.Time {
	return time.Now().UTC()
}

func (this *MyTradeController) PostTrade() {
	selllist := this.GetString("Selll")
	wantlist := this.GetString("Wantl")
	remark := this.GetString("Remark")
	hasanyoffer, _ := strconv.Atoi(this.GetString("Hasanyoffer"))
	shasmoney, _ := strconv.Atoi(this.GetString("Shasmoney"))
	smoney, _ := strconv.Atoi(this.GetString("Smoney"))
	whasmoney, _ := strconv.Atoi(this.GetString("Whasmoney"))
	wmoney, _ := strconv.Atoi(this.GetString("Wmoney"))
	var td models.Trades
	td.Uname = this.Uname
	td.Uid = this.Uid
	td.Status = 1
	td.Region = this.Region
	td.CreateTime = this.getTime()
	td.UpdateTime = this.getTime()
	td.Remark = remark
	td.Selllist = selllist
	td.Wantlist = wantlist
	td.Hasanyoffer = int64(hasanyoffer)
	td.Shasmoney = int64(shasmoney)
	td.Whasmoney = int64(whasmoney)
	td.Smoney = int64(smoney)
	td.Wmoney = int64(wmoney)
	td.Insert()
	this.Data["json"] = "ok"
	this.ServeJson()
}

func (this *MyTradeController) Show() {
	var trade models.Trades
	var err error
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	trade.Id = int64(id)
	trade.Read()
	//fmt.Println(id)
	//fmt.Print(trade)
	type Booksinfo struct {
		Bid    int64
		Ivid   int64
		Bname  string
		NOrate string
		Bimg   string
		Dbalt  string
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

	var tin Tinfos
	tin.Id = trade.Id
	tin.Uid = trade.Uid
	tin.Uname = trade.Uname
	tin.CreateTime = trade.CreateTime
	tin.UpdateTime = trade.UpdateTime
	tin.Region = trade.Region
	tin.Remark = trade.Remark
	if trade.Hasanyoffer == 0 {
		tin.Hasanyoffer = false
	} else {
		tin.Hasanyoffer = true
	}
	if trade.Shasmoney == 0 {
		tin.SHasmoney = false
	} else {
		tin.SHasmoney = true
		tin.SMoney = trade.Smoney
	}
	if trade.Whasmoney == 0 {
		tin.WHasmoney = false
	} else {
		tin.WHasmoney = true
		tin.WMoney = trade.Wmoney
	}
	if trade.Selllist != "" {
		var j1 []int
		err = json.Unmarshal(models.DecodeSS([]byte(trade.Selllist)), &j1)
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
				bi.Dbalt = bok.Dbalt
				bi.NOrate = inv.Newoldrate
				//fmt.Print(bi)
				//result[year] = append(result[year], v)
				bis = append(bis, bi)
			}
		}
		tin.Sells = bis
	}
	if trade.Wantlist != "" {
		var j2 []int
		err = json.Unmarshal(models.DecodeSS([]byte(trade.Wantlist)), &j2)
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
			bi.Dbalt = bok.Dbalt
			bi.Ivid = bok.Id
			//fmt.Print(bi)
			//result[year] = append(result[year], v)
			bis = append(bis, bi)
		}
		tin.Wants = bis
	}
	fmt.Print(tin)
	//读取offer和回复offer
	type Rpofr struct {
		Rfid       int64
		Oid        int64
		Uid        int64
		Uname      string
		Showonleft bool
		Toid       int64
		Tid        int64
		Msg        string
		CreateTime time.Time
		Status     int64
	}

	type Ofst struct {
		Id          int64
		Tid         int64
		Uid         int64
		Uname       string
		Offerlist   []Booksinfo
		Replylist   []Rpofr
		Hasaccept   bool
		Offerstatus int64
		Ohasmoney   int64
		Omoney      int64
		Msg         string
		CreateTime  time.Time
		Tstatus     int64
	}
	var olist []*models.Offers
	queryo := new(models.Offers).Query().Filter("tid", id)
	counto, _ := queryo.Count()
	if counto > 0 {
		queryo.OrderBy("-create_time").All(&olist)
	}
	fmt.Print(olist)
	var oflistinfo []Ofst
	for k, _ := range olist {
		var ofi Ofst
		ofi.Id = olist[k].Id
		o := orm.NewOrm()
		o.QueryTable("offers").Filter("id", olist[k].Id).Filter("offerstatus", 1).Update(orm.Params{
			"offerstatus": 2,
		})
		ofi.Tid = olist[k].Tid
		ofi.Uid = olist[k].Uid
		ofi.Uname = olist[k].Uname
		ofi.Offerstatus = olist[k].Offerstatus
		if olist[k].Offerstatus == 3 {
			ofi.Hasaccept = true
		} else {
			ofi.Hasaccept = false
		}
		ofi.Ohasmoney = olist[k].Ohasmoney
		ofi.Omoney = olist[k].Omoney
		ofi.Msg = olist[k].Msg
		ofi.CreateTime = olist[k].CreateTime
		ofi.Tstatus = olist[k].Tstatus
		if olist[k].Offerlist != "" {
			var j1 []int
			err = json.Unmarshal(models.DecodeSS([]byte(olist[k].Offerlist)), &j1)
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
					bi.Dbalt = bok.Dbalt
					bi.NOrate = inv.Newoldrate
					//fmt.Print(bi)
					//result[year] = append(result[year], v)
					bis = append(bis, bi)
				}
			}
			ofi.Offerlist = bis
		}
		var rfi []Rpofr
		var oflist []*models.Reply_offer
		queryof := new(models.Reply_offer).Query().Filter("status", 1).Filter("tid", id).Filter("oid", olist[k].Id)
		countof, _ := queryof.Count()
		if countof > 0 {
			queryof.OrderBy("create_time").All(&oflist)
		}
		for k1, _ := range oflist {
			var rf Rpofr
			rf.Rfid = oflist[k1].Id
			rf.Oid = oflist[k1].Oid
			rf.Uid = oflist[k1].Uid
			rf.Uname = oflist[k1].Uname
			rf.Toid = oflist[k1].Toid
			if oflist[k1].Uid == this.Uid {
				rf.Showonleft = false
			} else {
				rf.Showonleft = true
			}
			rf.Tid = oflist[k1].Tid
			rf.Msg = oflist[k1].Msg
			rf.CreateTime = oflist[k1].CreateTime
			rf.Status = oflist[k1].Status
			rfi = append(rfi, rf)
		}
		ofi.Replylist = rfi
		oflistinfo = append(oflistinfo, ofi)
	}
	fmt.Print(oflistinfo)
	this.Data["oflist"] = oflistinfo
	this.Data["tin"] = tin
	this.Data["PageTitle"] = "我的交换"
	this.Layout = "layout/default.html"
	this.TplNames = "mytrade_detail.html"
}

func (this *MyTradeController) RefuseOffer() {
	tid, _ := strconv.ParseInt(this.Ctx.Input.Param(":tid"), 10, 64)
	oid, _ := strconv.ParseInt(this.Ctx.Input.Param(":oid"), 10, 64)
	new(models.Offers).Query().Filter("tcid", this.Uid).Filter("id", oid).Delete()
	new(models.Reply_offer).Query().Filter("tid", tid).Filter("oid", oid).Delete()
	tids := strconv.Itoa(int(tid))
	this.Redirect("/mytrade/"+tids, 302)
}

func (this *MyTradeController) AcceptOffer() {
	tid, _ := strconv.ParseInt(this.Ctx.Input.Param(":tid"), 10, 64)
	oid, _ := strconv.ParseInt(this.Ctx.Input.Param(":oid"), 10, 64)
	ofr := models.Offers{Id: oid}
	if ofr.Read() == nil {
		ofr.Offerstatus = 3
		ofr.Update()
	}
	tids := strconv.Itoa(int(tid))
	this.Redirect("/mytrade/"+tids, 302)
}

func (this *MyTradeController) CancelOffer() {
	tid, _ := strconv.ParseInt(this.Ctx.Input.Param(":tid"), 10, 64)
	oid, _ := strconv.ParseInt(this.Ctx.Input.Param(":oid"), 10, 64)
	ofr := models.Offers{Id: oid}
	if ofr.Read() == nil {
		ofr.Offerstatus = 2
		ofr.Update()
	}
	tids := strconv.Itoa(int(tid))
	this.Redirect("/mytrade/"+tids, 302)
}

func (this *MyTradeController) PostReply() {
	tid, _ := strconv.Atoi(this.GetString("Tid"))
	oid, _ := strconv.Atoi(this.GetString("Oid"))
	toid, _ := strconv.Atoi(this.GetString("Toid"))
	msg := this.GetString("Msg")
	var rf models.Reply_offer
	rf.Tid = int64(tid)
	rf.Oid = int64(oid)
	rf.Toid = int64(toid)
	rf.Msg = msg
	rf.Uid = this.Uid
	rf.Uname = this.Uname
	rf.Status = 1
	rf.Isprivate = 0
	rf.CreateTime = this.getTime()
	rf.Insert()
	ofr := models.Offers{Id: int64(oid)}
	if ofr.Read() == nil {
		ofr.Offerstatus = 4
		ofr.Update()
	}
	td := models.Trades{Id: int64(tid)}
	if td.Read() == nil {
		td.UpdateTime = this.getTime()
		td.Update()
	}
	this.Data["json"] = "ok"
	this.ServeJson()
}

func (this *MyTradeController) DelReply() {
	tid, _ := strconv.ParseInt(this.Ctx.Input.Param(":tid"), 10, 64)
	rpid, _ := strconv.ParseInt(this.Ctx.Input.Param(":rpid"), 10, 64)
	ofr := models.Reply_offer{Id: rpid}
	if ofr.Read() == nil {
		counts, _ := new(models.Reply_offer).Query().Filter("oid", ofr.Oid).Count()
		if counts == 1 {
			o := orm.NewOrm()
			of := models.Offers{Id: ofr.Oid}
			of.Offerstatus = 2
			o.Update(&of, "offerstatus")
		}
	}
	new(models.Reply_offer).Query().Filter("id", rpid).Delete()
	tids := strconv.Itoa(int(tid))
	this.Redirect("/mytrade/"+tids, 302)
}
