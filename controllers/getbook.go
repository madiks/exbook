package controllers

import (
	"exbook/models"
	//"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type GetBookController struct {
	BaseController
}

func (this *GetBookController) getTime() time.Time {
	return time.Now().UTC()
}

type DbBookJson struct {
	Subtitle  string
	Pubdate   string
	Author    []string
	Publisher string
	Isbn10    string
	Isbn13    string
	Price     string
	Image     string
	Title     string
	Binding   string
	Alt       string
	Id        string
}

func (this *GetBookController) Getbinfo() {
	isbn := this.GetString("Isbn")
	var bk models.Books
	cond := orm.NewCondition()
	cond1 := cond.And("isbn10", isbn).Or("isbn13", isbn)
	query := new(models.Books).Query().SetCond(cond1)

	count, _ := query.Count()
	if count > 0 {
		query.One(&bk)
		this.Data["json"] = bk
	} else {
		req := httplib.Get("https://api.douban.com/v2/book/isbn/" + isbn)
		//str, _ := req.String()
		var result DbBookJson
		req.ToJson(&result)
		var bki models.Books
		if result.Title != "" {
			bki.Title = result.Title
			bki.Subtitle = result.Subtitle
			bki.Author = result.Author[0]
			bki.Pubdate = result.Pubdate
			bki.Image = result.Image
			bki.Binding = result.Binding
			bki.Dbalt = result.Alt
			bki.Isbn10 = result.Isbn10
			bki.Isbn13 = result.Isbn13
			bki.Publisher = result.Publisher
			bki.Dbid = result.Id
			bki.Price = result.Price
			bki.Isprivate = 0
			bki.Insert()
			this.Data["json"] = bki
			//fmt.Print("db1111")
		}

	}

	this.ServeJson()
}

func (this *GetBookController) AddInv() {
	norate := this.GetString("Norate")
	bid, _ := strconv.Atoi(this.GetString("Bid"))
	var inv models.Inventory
	inv.Uid = this.Uid
	inv.Bid = int64(bid)
	inv.Newoldrate = norate
	inv.CreateTime = this.getTime()
	inv.UpdateTime = this.getTime()
	inv.Status = 1
	inv.Insert()
	this.Data["json"] = "ok"
	this.ServeJson()
}

func (this *GetBookController) Searchbinfo() {
	isbn := this.GetString("Isbn")
	var bk models.Books
	cond := orm.NewCondition()
	cond1 := cond.And("isbn10", isbn).Or("isbn13", isbn)
	query := new(models.Books).Query().SetCond(cond1)

	count, _ := query.Count()
	if count > 0 {
		query.One(&bk)
		this.Data["json"] = bk
	} else {
		req := httplib.Get("https://api.douban.com/v2/book/isbn/" + isbn)
		//str, _ := req.String()
		var result DbBookJson
		err := req.ToJson(&result)
		if err == nil {
			var bki models.Books
			if result.Title != "" {
				bki.Title = result.Title
				bki.Subtitle = result.Subtitle
				//bki.Author = result.Author[0]
				bki.Pubdate = result.Pubdate
				bki.Image = result.Image
				bki.Binding = result.Binding
				bki.Dbalt = result.Alt
				bki.Isbn10 = result.Isbn10
				bki.Isbn13 = result.Isbn13
				bki.Publisher = result.Publisher
				bki.Dbid = result.Id
				bki.Price = result.Price
				bki.Isprivate = 0
				bki.Insert()
				this.Data["json"] = bki
			}
		}
	}

	this.ServeJson()
}
