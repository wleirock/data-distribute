package controllers

import (
	"wleirock/data-distribute/models"
	"wleirock/data-distribute/service"
)

// ScriptController 脚本管理
type ScriptController struct {
	BaseController
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
	scriptInfo := models.ScriptInfo{}
	err := c.ParseForm(&scriptInfo)
	if err != nil {
		c.ErrorListMsg()
	}

	list := service.GetScriptInfoList(&scriptInfo)
	count := service.GetScriptInfoTotal(&scriptInfo)
	result := queryData{0, count, list}
	c.Data["json"] = &result
	c.ServeJSON()
}
