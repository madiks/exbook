package controllers

import (
	"encoding/json"
	"exbook/models"
	//"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type SearchController struct {
	BaseAController
}

func (this *SearchController) Get() {
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
	this.Data["PageTitle"] = "搜索"
	this.Data["IsSearch"] = true
	this.Layout = "layout/default.html"
	this.TplNames = "search.html"
}

func (this *SearchController) getTime() time.Time {
	return time.Now().UTC()
}

func (this *SearchController) SearchT() {
	l := this.GetString("left")
	r := this.GetString("right")
	var page int
	var err error
	var num int64
	var list []models.Trades
	if page, err = strconv.Atoi(this.Ctx.Input.Param(":page")); err != nil || page < 1 {
		page = 1
	}
	var pagesize int = 8
	//fmt.Println("----------------------------参数-------------------------------")
	//fmt.Println(l)
	//fmt.Println(r)
	if l == "0" && r != "0" {
		var ividl []int64
		var sqlcon string = ""
		o := orm.NewOrm()
		o.Raw("select id from inventory where bid = ?", r).QueryRows(&ividl)
		//fmt.Println("**********************左执行结果******************************888")
		//fmt.Println(ividl)
		//o.Raw("SELECT count(*) FROM trades WHERE region= ? AND wantlist LIKE ? ORDER BY update_time desc", this.Region, "[%"+r+"%]").QueryRow(&num)
		for k, v := range ividl {
			sqlcon += "selllist LIKE '%," + strconv.Itoa(int(v)) + ",%'"
			if k != len(ividl)-1 {
				sqlcon += " OR "
			}
		}
		q1 := "SELECT count(*) FROM trades WHERE region= '" + this.Region + "' AND (" + sqlcon + ") ORDER BY update_time desc"
		o.Raw(q1).QueryRow(&num)
		q2 := "SELECT * FROM trades WHERE region= '" + this.Region + "' AND (" + sqlcon + ") ORDER BY update_time desc limit " + strconv.Itoa((page-1)*pagesize) + "," + strconv.Itoa(pagesize)
		o.Raw(q2).QueryRows(&list)
	} else if l != "0" && r == "0" {
		//fmt.Println("0000000**********************右执行结果******************************888")
		o := orm.NewOrm()
		o.Raw("SELECT count(*) FROM trades WHERE region= ? AND wantlist LIKE ? ORDER BY update_time desc", this.Region, "[%,"+l+",%]").QueryRow(&num)
		o.Raw("SELECT * FROM trades WHERE region= ? AND wantlist LIKE ? ORDER BY update_time desc limit ?,?", this.Region, "[%,"+l+",%]", (page-1)*pagesize, pagesize).QueryRows(&list)
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
	if num > 0 {
		this.Data["hasres"] = false
	} else {
		this.Data["hasres"] = true
	}
	this.Data["trades"] = tlist
	this.Data["pagebar"] = models.NewPager(int64(page), num, int64(pagesize), "/trade/search?left="+l+"&right="+r).ToString()
	//fmt.Println(num)
	//fmt.Println("-----------------------doS--------------------------")
	this.Data["PageTitle"] = "搜索结果"
	this.Data["IsAddTrade"] = true
	this.Layout = "layout/default.html"
	this.TplNames = "search_result.html"
}
