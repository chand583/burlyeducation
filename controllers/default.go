package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "burlyeducation.com"
	c.Data["Email"] = "chand583@gmail.com"
	c.TplName = "index.tpl"
}
