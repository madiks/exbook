package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//
type Activity_log struct {
	Id         int64
	Aid        int64 `orm:"index"`
	Uid        int64
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
}

func (m *Activity_log) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Activity_log) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Activity_log) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Activity_log) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Activity_log) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
