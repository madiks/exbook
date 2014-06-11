package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//私信模型
type Message_reply struct {
	Id         int64
	Mid        int64
	Uid        int64
	Uname      string    `orm:"size(200)"`
	Msg        string    `orm:"size(500)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	Status     int64
	Tag        string `orm:"size(200)"`
}

func (m *Message_reply) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Message_reply) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Message_reply) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Message_reply) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Message_reply) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
