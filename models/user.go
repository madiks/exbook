package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//用户模型
type User struct {
	Id        int64
	Uname     string `orm:"unique"`
	Pwd       string
	Region    string    `orm:"index"`
	Tname     string    `orm:"size(300)"`
	Tel       string    `orm:"size(300)"`
	Img       string    `orm:"size(400)"`
	Remark    string    `orm:"size(500)"`
	Lastlogin time.Time `orm:"auto_now;type(datetime)"`
	TradeRank int64     `orm:"index"`
	OfferRank int64     `orm:"index"`
	Email     string
	Address   string `orm:"size(500)"`
}

func (m *User) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
