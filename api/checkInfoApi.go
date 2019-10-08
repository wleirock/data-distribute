package api

import (
	"strconv"
	"wleirock/data-distribute/service"
)

// CheckInfoAPIController 报告类型数据处理
type CheckInfoAPIController struct {
	DataAPIController
}

// CheckInfo 接收检查数据
func (c *CheckInfoAPIController) CheckInfo() {
	dataType := "CHECK"

	hosID, err := c.GetInt("hospitalId")
	if err != nil {
		c.ErrorJSON("-1", "hospitalId获取失败", err)
		return
	}

	param := string(c.Ctx.Input.RequestBody)

	scriptList := service.GetScriptInfoListByHosFk(hosID, dataType)
	for _, script := range scriptList {
		go c.ProcessLua(script.ScriptName, script.ApiIp+":"+strconv.Itoa(script.ApiPort), param)
	}

	resultData := ResponseEntity{"0", true, "success"}
	c.Data["json"] = &resultData
	c.ServeJSON()
}
