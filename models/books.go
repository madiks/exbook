package models

import (
	"github.com/astaxie/beego/orm"
)

//书籍模型
type Books struct {
	Id         int64
	Title      string `orm:"size(300)"`
	Subtitle   string `orm:"size(300)"`
	Author     string `orm:"size(300)"`
	Pubdate    string `orm:"size(200)"`
	Image      string `orm:"size(400)"`
	Binding    string `orm:"size(50)"`
	Translator string `orm:"size(100)"`
	Dbalt      string `orm:"size(400)"`
	Publisher  string `orm:"size(200)"`
	Isbn10     string `orm:"size(200)"`
	Isbn13     string `orm:"size(200)"`
	Dbid       string `orm:"size(300)"`
	Price      string `orm:"size(50)"`
	Isprivate  int64  `orm:"index"`
	Pages      int64
}

func (m *Books) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Books) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Books) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Books) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Books) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
