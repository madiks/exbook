package main

import (
	"exbook/g"
	_ "exbook/routers"
	"github.com/astaxie/beego"
)

func main() {
	g.InitEnv()
	beego.Run()
}
