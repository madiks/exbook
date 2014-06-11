package controllers

import (
	"encoding/json"
	"exbook/models"
	//"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type MyOfferController struct {
	BaseAController
}

func (this *MyOfferController) Get() {
	o := orm.NewOrm()
	var idlist orm.ParamsList
	var Isempty bool
	num, err := o.Raw("SELECT distinct(tid) FROM offers WHERE uid = ? order by create_time desc", this.Uid).ValuesFlat(&idlist)
	if err == nil && num > 0 {
		Isempty = false
		//fmt.Println(idlist) // []{"1","2","3",...}
		//fmt.Println(num)
		var list []models.Trades
		for _, v := range idlist {
			var td models.Trades
			tids, _ := v.(string)
			tid, _ := strconv.Atoi(tids)
			//fmt.Println(tid)
			td.Id = int64(tid)
			td.Read()
			list = append(list, td)
		}
		//fmt.Print(list)

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
			Newrpnum    int64
			Shownewrp   bool
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
			queryof := new(models.Reply_offer).Query().Filter("toid", this.Uid).Filter("status", 1).Filter("tid", list[k].Id)
			countof, _ := queryof.Count()
			if countof > 0 {
				tin.Shownewrp = true
				tin.Newrpnum = countof
			} else {
				tin.Shownewrp = false
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
		//fmt.Println(tlist)
		this.Data["trades"] = tlist
	} else {
		Isempty = true
	}
	this.Data["isempty"] = Isempty
	this.Data["PageTitle"] = "我的报价"
	this.Data["IsMyOffer"] = true
	this.Layout = "layout/default.html"
	this.TplNames = "myoffer.html"
}

func (this *MyOfferController) Show() {
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
	//fmt.Print(tin)
	//读取offer和回复offer
	type Rpofr struct {
		Rfid         int64
		Oid          int64
		Uid          int64
		Uname        string
		Showreplybtn bool
		Tid          int64
		Msg          string
		CreateTime   time.Time
		Status       int64
	}

	type Ofst struct {
		Id           int64
		Tid          int64
		Uid          int64
		Uname        string
		Offerlist    []Booksinfo
		Replylist    []Rpofr
		Hasaccept    bool
		Showdeloffer bool
		Offerstatus  int64
		Ohasmoney    int64
		Omoney       int64
		Msg          string
		CreateTime   time.Time
		Tstatus      int64
	}
	var olist []*models.Offers
	queryo := new(models.Offers).Query().Filter("tid", id)
	counto, _ := queryo.Count()
	if counto > 0 {
		queryo.OrderBy("-create_time").All(&olist)
	}
	//fmt.Print(olist)
	var oflistinfo []Ofst
	for k, _ := range olist {
		var ofi Ofst
		ofi.Id = olist[k].Id
		o := orm.NewOrm()
		o.QueryTable("reply_offer").Filter("tid", olist[k].Id).Filter("status", 1).Update(orm.Params{
			"status": 2,
		})
		ofi.Tid = olist[k].Tid
		ofi.Uid = olist[k].Uid
		ofi.Uname = olist[k].Uname
		ofi.Offerstatus = olist[k].Offerstatus
		if olist[k].Uid == this.Uid {
			ofi.Showdeloffer = true
		} else {
			ofi.Showdeloffer = false
		}
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
			if oflist[k1].Uid == this.Uid {
				rf.Showreplybtn = false
			} else {
				rf.Showreplybtn = true
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
	//fmt.Print(oflistinfo)
	this.Data["oflist"] = oflistinfo
	this.Data["tin"] = tin
	this.Data["PageTitle"] = "我的报价"
	this.Layout = "layout/default.html"
	this.TplNames = "myoffer_detail.html"
}

func (this *MyOfferController) DeleteOffer() {
	tid, _ := strconv.ParseInt(this.Ctx.Input.Param(":tid"), 10, 64)
	oid, _ := strconv.ParseInt(this.Ctx.Input.Param(":oid"), 10, 64)
	new(models.Offers).Query().Filter("uid", this.Uid).Filter("id", oid).Delete()
	new(models.Reply_offer).Query().Filter("toid", this.Uid).Filter("oid", oid).Delete()
	tids := strconv.Itoa(int(tid))
	this.Redirect("/myoffer/"+tids, 302)
}

func (this *MyOfferController) DeleteAllOffer() {
	tid, _ := strconv.ParseInt(this.Ctx.Input.Param(":tid"), 10, 64)
	new(models.Offers).Query().Filter("uid", this.Uid).Filter("tid", tid).Delete()
	new(models.Reply_offer).Query().Filter("toid", this.Uid).Filter("tid", tid).Delete()
	this.Redirect("/myoffer/", 302)
}

func (this *MyOfferController) DeletePostReply() {
	tid, _ := strconv.ParseInt(this.Ctx.Input.Param(":tid"), 10, 64)
	rpid, _ := strconv.ParseInt(this.Ctx.Input.Param(":rpid"), 10, 64)
	new(models.Reply_offer).Query().Filter("uid", this.Uid).Filter("tid", tid).Filter("id", rpid).Delete()
	tids := strconv.Itoa(int(tid))
	this.Redirect("/myoffer/"+tids, 302)
}

func (this *MyOfferController) GetTraderInfo() {
	oid, _ := strconv.Atoi(this.GetString("Oid"))
	var ofr models.Offers
	//fmt.Println("****************************************************************")
	ofr.Id = int64(oid)
	ofr.Read()
	//fmt.Println(ofr)
	var uinfo models.User
	uinfo.Id = ofr.Tcid
	uinfo.Read()

	//fmt.Println(uinfo)
	var str string = "Ta的联系电话：" + uinfo.Tel + "<br>Ta的地址：" + uinfo.Address
	this.Data["json"] = str
	this.ServeJson()
}
