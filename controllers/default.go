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

//Get 首页
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
