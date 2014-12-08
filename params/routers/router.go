package routers

import (
	"github.com/astaxie/beego"

	"github.com/bakanis/training/params/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/post/:id", &controllers.MainController{}, "get:GetPostById")
}
