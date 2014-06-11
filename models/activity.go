package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//同城活动模型
type Activity struct {
	Id         int64
	Author     string `orm:"size(100)"`
	Aid        int64
	Title      string `orm:"size(500)"`
	Condition  string `orm:"size(700)"`
	Content    string `orm:"size(300)"`
	Like       int64  `orm:"index"`
	Limit      int64
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	Join       int64     `orm:"index"`
	Region     string    `orm:"size(100)"`
	Status     int64
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
	StartTime  time.Time
	EndTime    time.Time
}

func (m *Activity) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Activity) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Activity) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Activity) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Activity) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
