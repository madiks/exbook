package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//私信模型
type Message struct {
	Id         int64
	FromUid    int64     `orm:"index"`
	FromUname  string    `orm:"size(200)"`
	ForUid     int64     `orm:"index"`
	ForUname   string    `orm:"size(200)"`
	Msg        string    `orm:"size(500)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	Status     int64
	Tid        int64
	Istrade    int64
	Tag        string `orm:"size(200)"`
}

func (m *Message) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Message) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Message) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Message) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Message) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
