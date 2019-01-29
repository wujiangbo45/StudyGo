package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type TestController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type Result interface {
	SetData() ResultStruct
}

type ret struct {
}

type Success struct {
	sct interface{}
}
type Error struct {
	message string
}
type ResultStruct struct {
	ResultCode int32       `json:"result_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func (su Success) SetData() ResultStruct {
	return ResultStruct{ResultCode: 200, Message: "成功", Data: su.sct}
}

func (er Error) SetData() ResultStruct {
	return ResultStruct{ResultCode: 500, Message: er.message, Data: nil}
}

func (c *TestController) Get() {
	var s = c.Ctx.Input.Param(":id")
	var m Result
	if s == "0" {
		m = &Success{sct: struct {
			Name string `json:"name"`
		}{"wujiangbo"}}
	} else {
		m = &Error{message: "错误"}
	}

	c.Data["json"] = m.SetData()
	c.ServeJSON()
}
