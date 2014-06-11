package controllers

import (
	"encoding/json"
	"exbook/models"
	"fmt"
	"strconv"
	"time"
)

type TradeController struct {
	BaseController
}

func (this *TradeController) Show() {
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
		Tid        int64
		Showonleft bool
		Toid       int64
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
		Offerstatus int64
		Ohasmoney   int64
		Hasaccept   bool
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
			rf.Toid = oflist[k1].Toid
			if oflist[k1].Uid == oflist[k1].Toid {
				rf.Showonleft = true
			} else {
				rf.Showonleft = false
			}
			rf.Uname = oflist[k1].Uname
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
	this.Data["PageTitle"] = "交换"
	this.Layout = "layout/default.html"
	this.TplNames = "trade.html"
}
