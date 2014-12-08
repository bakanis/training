package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) GetPostById() {
	postId, err := c.GetInt(":id")

	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	beego.Debug("post id: ", postId)

	//but we might need int64 if we do a sql query
	postId64, err := c.GetInt64(":id")

	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	beego.Debug("post id as int64: ", postId64)

	c.TplNames = "index.tpl"
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}
