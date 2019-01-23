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
	SetData(sct interface{}) ResultStruct
	SetMessage(msg string) ResultStruct
}

type Success struct {
	Result
}
type Error struct {
	Result
}
type ResultStruct struct {
	ResultCode int32       `json:"result_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func (su Success) SetData(sct interface{}) ResultStruct {
	return ResultStruct{ResultCode: 200, Message: "成功", Data: sct}
}

func (er Error) SetMessage(msg string) ResultStruct {
	return ResultStruct{ResultCode: 500, Message: msg, Data: nil}
}

func (c *TestController) Get() {
	var s = c.Ctx.Input.Param(":id")
	var m Result
	var res ResultStruct
	if s == "0" {
		m = Success{}
		res = m.SetData(struct {
			Name string `json:"name"`
		}{"wujiangbo"})
	} else {
		m = Error{}
		res = m.SetMessage("调用错误")
	}

	c.Data["json"] = res
	c.ServeJSON()
}
