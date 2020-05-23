package main

import (
	_ "liteblog/routers"
	"github.com/astaxie/beego"
	"strings"
)

func main() {
	initTemplate()
	beego.Run()
}

func initTemplate()  {
	beego.AddFuncMap("equrl", func(x,y string) bool {
		x1 := strings.Trim(x, "/")
		y1 := strings.Trim(y, "/")
		return strings.Compare(x1, y1) == 0
	})
}