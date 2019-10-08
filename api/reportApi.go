package api

import (
	"strconv"
	"wleirock/data-distribute/service"
)

// ReportAPIController 报告类型数据处理
type ReportAPIController struct {
	DataAPIController
}

// Report 接收报告数据
func (c *ReportAPIController) Report() {
	dataType := "REPORT"

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
