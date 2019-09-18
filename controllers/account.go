package controllers

import (
	"github.com/astaxie/beego"
)

// AccountController 账户
type AccountController struct {
	beego.Controller
}

// Login 登录
func (c *AccountController) Login() {
	username := c.GetString("username")
	password := c.GetString("password")
	// 先做个假登录
	if username == "admin" && password == "admin" {
		c.SetSession("username", username)
		c.Data["username"] = username
		c.TplName = "panel.html"
	} else {
		c.Data["msg"] = "用户名或密码错误"
		c.Redirect("/", 301)
	}
}

// Logout 退出
func (c *AccountController) Logout() {
	c.DelSession("username")
	c.Redirect("/", 301)
}
