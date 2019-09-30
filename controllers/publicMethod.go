package controllers

import (
	"fmt"
	"wleirock/data-distribute/models"
	"wleirock/data-distribute/service"
)

// PublicMethodController 公共方法
type PublicMethodController struct {
	BaseController
}

type methodQueryData struct {
	Code int                   `json:"code"`
	Data []models.PublicMethod `json:"data"`
}

// Index 页面
func (c *PublicMethodController) Index() {
	c.TplName = "method_list.html"
}

// List 列表数据
func (c *PublicMethodController) List() {
	list := service.GetAllPublicMethodList()
	result := methodQueryData{0, list}
	c.Data["json"] = &result
	c.ServeJSON()
}

// Add 新增或修改页面
func (c *PublicMethodController) Add() {
	methodPk, err := c.GetInt("methodPk")
	if err != nil {
		fmt.Println(err)
	}
	method := service.GetPublicMethodByPk(methodPk)
	c.Data["method"] = method
	c.TplName = "method_add.html"
}

// Save 保存或更新
func (c *PublicMethodController) Save() {
	res := true
	method := models.PublicMethod{}
	err := c.ParseForm(&method)
	if err != nil {
		c.ErrorMsg("公共方法保存失败", err)
		return
	}

	if method.MethodPk == 0 {
		// 新增
		res = service.AddPublicMethod(&method)
	} else {
		// 修改
		res = service.UpdatePublicMethod(&method)
	}

	if res {
		c.SuccessMsg("保存成功")
	} else {
		c.ErrorMsg("公共方法保存失败", nil)
	}
}

// Delete 删除
func (c *PublicMethodController) Delete() {
	scriptName := c.GetString("scriptName")
	if scriptName == "" {
		c.ErrorMsg("删除脚本文件:获取参数失败", nil)
		return
	}

	ifExist, err := service.IfSciptInfoExist(scriptName)
	if err != nil {
		c.ErrorMsg("删除脚本文件失败", err)
		return
	}
	if ifExist {
		c.ErrorMsg("该文件已被引用，不能删除", nil)
		return
	}

	methodPk, err := c.GetInt("methodPk")
	if err != nil {
		c.ErrorMsg("删除公共方法，获取参数失败", err)
	}

	res := service.DeletePublicMethod(methodPk)

	if res {
		c.SuccessMsg("删除成功")
	} else {
		c.ErrorMsg("删除失败", nil)
	}
}
