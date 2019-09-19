package controllers

import (
	"github.com/astaxie/beego"
)

// ResponseEntity 定义返回的JSON字段
type ResponseEntity struct {
	ResultCode string
	IsSuccess  bool
	ResultMsg  string
}

// BaseController 首页
type BaseController struct {
	beego.Controller
}

//Get 登录界面
func (c *BaseController) Get() {
	c.TplName = "login.html"
}

// ErrorListMsg ajax请求数据发生错误时返回数据
func (c *BaseController) ErrorListMsg() {
	msg := map[string]int{"code": -1}
	c.Data["json"] = &msg
	c.ServeJSON()
}
