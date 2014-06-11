package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//收藏模型
type Bookmark struct {
	Id         int64
	Uid        int64
	Type       int64 `orm:"index"`
	Tid        int64
	Bid        int64
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	Status     int64
}

func (m *Bookmark) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Bookmark) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Bookmark) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Bookmark) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Bookmark) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
