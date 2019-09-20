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

// AjaxMsg 页面ajax返回值定义
type AjaxMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// BaseController 首页
type BaseController struct {
	beego.Controller
}

//Get 登录界面
func (c *BaseController) Get() {
	c.TplName = "login.html"
}

// ErrorMsg ajax请求数据发生错误时返回数据
func (c *BaseController) ErrorMsg(content string) {
	msg := AjaxMsg{-1, content}
	c.Data["json"] = &msg
	c.ServeJSON()
}

// SuccessMsg ajax请求数据成功时返回数据
func (c *BaseController) SuccessMsg(content string) {
	msg := AjaxMsg{0, content}
	c.Data["json"] = &msg
	c.ServeJSON()
}
