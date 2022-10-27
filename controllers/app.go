package controllers

import (
	"github.com/astaxie/beego"
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

func (a *AppController) Response(res interface{}, message string) {
	a.Data["json"] = Success{Status: 1, Data: res, Message: message}
	a.ServeJSON()
}

func (a *AppController) ResponseWithError(code int, msg string) {
	a.Data["json"] = Error{Status: 0, Code: code, Message: msg}
	a.ServeJSON()
}
