package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//管理员模型
type Admin struct {
	Id        int64
	Uname     string `orm:"unique"`
	Pwd       string
	Ip        string
	Lastlogin time.Time `orm:"auto_now;type(datetime)"`
}

func (m *Admin) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Admin) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Admin) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Admin) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Admin) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
