package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//回复报价模型
type Reply_offer struct {
	Id         int64
	Oid        int64
	Uid        int64
	Uname      string `orm:"size(200)"`
	Tid        int64
	Msg        string `orm:"size(500)"`
	Toid       int64
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	Status     int64
	Isprivate  int64
}

func (m *Reply_offer) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Reply_offer) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Reply_offer) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Reply_offer) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Reply_offer) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
