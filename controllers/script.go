package controllers

import (
	"fmt"
	"wleirock/data-distribute/models"
	"wleirock/data-distribute/service"
	"wleirock/data-distribute/utils"
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
		c.ErrorMsg("脚本列表查询失败", err)
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
		c.ErrorMsg("脚本设置保存失败", err)
		return
	}

	// 校验脚本文件是否存在(只有自定义的脚本校验，公共方法不校验)
	if scriptInfo.Status == "U" {
		if !utils.LuaFileExist(scriptInfo.ScriptName) {
			c.ErrorMsg("脚本文件不存在", nil)
			return
		}
	}

	if scriptInfo.InfoPk == 0 {
		// 新增
		if scriptInfo.Status == "" {
			scriptInfo.Status = "U"
		}
		res = service.SaveScriptInfo(&scriptInfo)
	} else {
		// 修改
		res = service.UpdateScriptInfo(&scriptInfo)
	}

	if res {
		c.SuccessMsg("保存成功")
	} else {
		c.ErrorMsg("脚本设置保存失败", nil)
	}
}

// Delete 删除
func (c *ScriptController) Delete() {
	infoPk, err := c.GetInt("infoPk")
	if err != nil {
		c.ErrorMsg("删除脚本设置，获取参数失败", err)
	}

	res := service.DeleteScriptInfo(infoPk)

	if res {
		c.SuccessMsg("删除成功")
	} else {
		c.ErrorMsg("删除失败", nil)
	}
}

// GetHospitalList 获取hospital列表
func (c *ScriptController) GetHospitalList() {
	hosList := service.GetHospitalInfoList()
	c.Data["json"] = &hosList
	c.ServeJSON()
}

// GetMethodList 获取公共方法列表
func (c *ScriptController) GetMethodList() {
	methodList := service.GetAllPublicMethodList()
	c.Data["json"] = &methodList
	c.ServeJSON()
}
