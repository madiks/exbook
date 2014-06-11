package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//库存模型
type Inventory struct {
	Id         int64
	Uid        int64
	Bid        int64
	Newoldrate string    `orm:"size(50)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
	Status     int64
}

func (m *Inventory) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Inventory) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Inventory) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Inventory) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Inventory) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
