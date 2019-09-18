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

//MainController 首页
type MainController struct {
	beego.Controller
}

//Get 登录界面
func (c *MainController) Get() {
	c.TplName = "login.html"
}
