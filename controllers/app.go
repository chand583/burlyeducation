package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	//"github.com/astaxie/beego"
)

type AppController struct {
	beego.Controller
}

type Error struct {
	Status  int    `json:"status"`
	Code    int    `json:"error_code"`
	Message string `json:"error_msg"`
}
type Success struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (ac *AppController) Response(res interface{}, message string) {
	ac.Data["json"] = Success{Status: 1, Data: res, Message: message}
	ac.ServeJSON()
}
func (ac *AppController) ResponseWithError(code int, message string) {
	ac.Data["json"] = Error{Status: 0, Code: code, Message: message}
	ac.ServeJSON()
}
