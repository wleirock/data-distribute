package controllers

import (
	"wleirock/data-distribute/models"
	"wleirock/data-distribute/service"

	"github.com/astaxie/beego"
)

// ScriptController 脚本管理
type ScriptController struct {
	beego.Controller
}

type queryData struct {
	Code  int                 `json:"code"`
	Count int                 `json:"count"`
	Data  []models.ScriptInfo `json:"data"`
}

// Index 页面
func (c *ScriptController) Index() {
	c.TplName = "script_list.html"
}

// List 列表数据
func (c *ScriptController) List() {
	pageIndex, err := c.GetInt("page")
	if err != nil {
		pageIndex = 1
	}
	pageSize, err := c.GetInt("limit")
	if err != nil {
		pageSize = 10
	}

	list := service.GetScriptInfoList(pageIndex, pageSize)
	count := service.GetScriptInfoTotal()
	result := queryData{0, count, list}
	c.Data["json"] = &result
	c.ServeJSON()
}
