package controllers

import (
	"fmt"
	"wleirock/data-distribute/models"
	"wleirock/data-distribute/service"
)

// ScriptFileController 脚本文件
type ScriptFileController struct {
	BaseController
}

type fileQueryDataFile struct {
	Code  int                 `json:"code"`
	Count int                 `json:"count"`
	Data  []models.ScriptFile `json:"data"`
}

// Index 脚本文件列表展示
func (c *ScriptFileController) Index() {
	c.TplName = "script_file.html"
}

// List 获取脚本文件列表
func (c *ScriptFileController) List() {
	scriptFile := models.ScriptFile{}
	err := c.ParseForm(&scriptFile)
	if err != nil {
		c.ErrorMsg("脚本文件列表查询失败", err)
	}

	list := service.GetScriptFileList(&scriptFile)
	count := service.GetScriptFileTotal(&scriptFile)
	result := fileQueryDataFile{0, count, list}
	c.Data["json"] = &result
	c.ServeJSON()
}

// Delete 删除
func (c *ScriptFileController) Delete() {
	filePk, err := c.GetInt("filePk")
	if err != nil {
		c.ErrorMsg("删除脚本设置，获取参数失败", err)
	}

	res := service.DeleteScriptFile(filePk)

	if res {
		c.SuccessMsg("删除成功")
	} else {
		c.ErrorMsg("删除失败", nil)
	}
}

// Upload 上传脚本文件
func (c *ScriptFileController) Upload() {
	f, h, err := c.GetFile("file")
	if err != nil {
		c.ErrorMsg("脚本文件上传失败", err)
		return
	}
	defer f.Close()

	fmt.Println("上传的文件--", h.Filename)

	err = c.SaveToFile("file", "luafile/"+h.Filename)
	if err != nil {
		c.ErrorMsg("脚本文件上传失败", err)
		return
	}
	c.SuccessMsg("脚本文件上传成功")
}
