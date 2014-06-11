package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//报价模型
type Offers struct {
	Id          int64
	Tid         int64
	Uid         int64
	Uname       string `orm:"size(200)"`
	Offerlist   string `orm:"size(1000)"`
	Offerstatus int64
	Ohasmoney   int64
	Omoney      int64
	Tcid        int64
	Msg         string    `orm:"size(500)"`
	CreateTime  time.Time `orm:"auto_now_add;type(datetime)"`
	Tstatus     int64
}

func (m *Offers) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Offers) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Offers) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Offers) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Offers) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
