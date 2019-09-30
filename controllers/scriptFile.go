package controllers

import (
	"strings"
	"time"
	"wleirock/data-distribute/models"
	"wleirock/data-distribute/service"
	"wleirock/data-distribute/utils"
)

// ScriptFileController 脚本文件
type ScriptFileController struct {
	BaseController
}

type fileQueryData struct {
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
	result := fileQueryData{0, count, list}
	c.Data["json"] = &result
	c.ServeJSON()
}

// Delete 删除
func (c *ScriptFileController) Delete() {
	fileName := c.GetString("fileName")
	if fileName == "" {
		c.ErrorMsg("删除脚本文件:获取参数失败", nil)
		return
	}

	ifExist, err := service.IfSciptInfoExist(fileName)
	if err != nil {
		c.ErrorMsg("删除脚本文件失败", err)
		return
	}
	if ifExist {
		c.ErrorMsg("该文件已被引用，不能删除", nil)
		return
	}

	filePk, err := c.GetInt("filePk")
	if err != nil {
		c.ErrorMsg("删除脚本文件:获取参数失败", err)
		return
	}

	res := service.DeleteScriptFile(filePk)

	if res {
		utils.DeleteLuaFile(fileName)
		c.SuccessMsg("删除成功")
	} else {
		c.ErrorMsg("删除脚本文件失败", nil)
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

	nameAry := strings.Split(h.Filename, ".")
	if len(nameAry) != 2 {
		c.ErrorMsg("文件名称有误", err)
		return
	}
	if nameAry[1] != "lua" {
		c.ErrorMsg("文件名称有误", err)
		return
	}

	err = c.SaveToFile("file", "luafile/"+h.Filename)
	if err != nil {
		c.ErrorMsg("脚本文件上传失败", err)
		return
	}

	scriptFile := models.ScriptFile{}
	scriptFile.FileName = h.Filename
	scriptFile.FileSize = h.Size
	scriptFile.Status = "U"
	scriptFile.UploadTime = time.Now().Format("2006-01-02 15:04:05")

	res := true
	existScriptFile := service.GetScriptFileByName(h.Filename)
	if existScriptFile.FilePk != 0 {
		scriptFile.FilePk = existScriptFile.FilePk
		res = service.UpdateScriptFile(&scriptFile)
	} else {
		res = service.AddScriptFile(&scriptFile)
	}

	if res {
		c.SuccessMsg("脚本文件上传成功")
	} else {
		c.ErrorMsg("脚本文件上传失败", nil)
	}
}

// Download 下载脚本文件
func (c *ScriptFileController) Download() {
	fileName := c.GetString("fileName")
	if fileName == "" {
		c.ErrorMsg("Download 文件名称为空", nil)
		return
	}
	c.Ctx.Output.Download("luafile/" + fileName)
}
