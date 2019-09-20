package controllers

import (
	"fmt"
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
		c.ErrorMsg("查询失败")
	}

	list := service.GetScriptInfoList(&scriptInfo)
	count := service.GetScriptInfoTotal(&scriptInfo)
	result := queryData{0, count, list}
	c.Data["json"] = &result
	c.ServeJSON()
}

// Add 新增或修改页面
func (c *ScriptController) Add() {
	infoPk, err := c.GetInt("infoPk")
	if err != nil {
		fmt.Println(err)
	}
	scriptInfo := service.GetScriptInByPk(infoPk)
	c.Data["scriptInfo"] = scriptInfo
	c.TplName = "script_add.html"
}

// Save 保存或更新
func (c *ScriptController) Save() {
	res := true
	scriptInfo := models.ScriptInfo{}
	err := c.ParseForm(&scriptInfo)
	if err != nil {
		c.ErrorMsg("保存失败")
	}

	if scriptInfo.InfoPk == 0 {
		// 新增
		scriptInfo.Status = "A"
		res = service.SaveScriptInfo(&scriptInfo)
	} else {
		res = service.UpdateScriptInfo(&scriptInfo)
	}

	if res {
		c.SuccessMsg("保存成功")
	} else {
		c.ErrorMsg("保存失败")
	}
}
