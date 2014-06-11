package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//私信模型
type Trades struct {
	Id          int64
	Uid         int64
	Uname       string `orm:"size(200)"`
	Selllist    string `orm:"size(1000)"`
	Wantlist    string `orm:"size(1000)"`
	Hasanyoffer int64
	Shasmoney   int64
	Smoney      int64
	Whasmoney   int64
	Wmoney      int64
	CreateTime  time.Time `orm:"auto_now_add;type(datetime)"`
	Status      int64
	Region      string    `orm:"size(100)"`
	UpdateTime  time.Time `orm:"auto_now;type(datetime)"`
	Remark      string    `orm:"size(500)"`
	Dealuid     int64
	Dealuname   string `orm:"size(200)"`
	TEstimation int64
	OEstimation int64
	TCommment   string `orm:"size(500)"`
	OCommment   string `orm:"size(500)"`
}

func (m *Trades) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Trades) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Trades) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Trades) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Trades) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
