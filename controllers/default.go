package controllers

import (
	"encoding/json"
	"exbook/models"
	//"fmt"
	"strconv"
	"time"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	var (
		list      []*models.Trades
		pageSize  int64
		err       error
		pageNo    int64
		Isnotrade bool
	)

	if pageNo, err = strconv.ParseInt(this.Ctx.Input.Param(":page"), 10, 64); err != nil || pageNo < 1 {
		pageNo = 1
	}
	pageSize = 8
	query := new(models.Trades).Query().Filter("status", 1).Filter("region", this.Region)
	count, _ := query.Count()
	if count > 0 {
		Isnotrade = false
		query.OrderBy("-update_time").Limit(pageSize, (pageNo-1)*pageSize).All(&list)
	} else {
		Isnotrade = true
	}
	this.Data["isnotrade"] = Isnotrade
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

	//fmt.Print(tlist)
	this.Data["trades"] = tlist
	this.Data["pagebar"] = models.NewPager(int64(pageNo), int64(count), int64(pageSize), "").ToString()
	this.Data["PageTitle"] = "首页"
	this.Data["IsIndex"] = true
	this.Layout = "layout/default.html"
	this.TplNames = "index.html"
}
