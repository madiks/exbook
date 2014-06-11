package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//注册model
	orm.RegisterModel(new(Admin), new(User), new(Activity), new(Regionset), new(Activity_log), new(Bookmark), new(Books), new(Inventory), new(Message), new(Message_reply), new(Offers), new(Reply_offer), new(Trades))
}

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func DecodeS(buf string) []byte {
	var news []byte
	news = append(news, '[')
	for i := 2; i < len(buf)-2; i++ {
		news = append(news, buf[i])
	}
	news = append(news, ']')
	return news
}

func DecodeSS(buf []byte) []byte {
	var news []byte
	news = append(news, '[')
	news = append(news, buf[2:len(buf)-2]...)
	news = append(news, ']')
	return news
}
