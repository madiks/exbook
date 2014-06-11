package models

import (
	"github.com/astaxie/beego/orm"
)

//地区模型
type Regionset struct {
	Id     int64
	Region string `orm:"size(200)"`
}

func (m *Regionset) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Regionset) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Regionset) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Regionset) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Regionset) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
