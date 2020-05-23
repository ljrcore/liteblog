package routers

import (
	"github.com/astaxie/beego"
	"liteblog/controllers"
)

func init() {
	beego.Include(&controllers.IndexController{})
}
