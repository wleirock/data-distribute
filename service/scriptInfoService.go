package service

import (
	"fmt"
	"strings"
	"wleirock/data-distribute/models"

	"github.com/astaxie/beego/orm"
)

// GetScriptInfoList 查询脚本信息集合
func GetScriptInfoList(scriptInfo *models.ScriptInfo) []models.ScriptInfo {
	o := orm.NewOrm()

	var buildStr strings.Builder
	buildStr.WriteString("select * from script_info where 1=1")
	if scriptInfo.HospitalName != "" {
		buildStr.WriteString(" and hospital_name like '%")
		buildStr.WriteString(scriptInfo.HospitalName)
		buildStr.WriteString("%'")
	}
	if scriptInfo.ScriptName != "" {
		buildStr.WriteString(" and script_name='")
		buildStr.WriteString(scriptInfo.ScriptName)
		buildStr.WriteString("'")
	}
	buildStr.WriteString(" limit ? , ?")
	var list []models.ScriptInfo
	_, err := o.Raw(buildStr.String(), (scriptInfo.Page-1)*scriptInfo.Limit, scriptInfo.Limit).QueryRows(&list)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return list
}

// GetScriptInfoTotal 查询总数
func GetScriptInfoTotal(scriptInfo *models.ScriptInfo) int {
	o := orm.NewOrm()

	var buildStr strings.Builder
	buildStr.WriteString("select count(*) from script_info where 1=1")
	if scriptInfo.HospitalName != "" {
		buildStr.WriteString(" and hospital_name like '%")
		buildStr.WriteString(scriptInfo.HospitalName)
		buildStr.WriteString("%'")
	}
	if scriptInfo.ScriptName != "" {
		buildStr.WriteString(" and script_name='")
		buildStr.WriteString(scriptInfo.ScriptName)
		buildStr.WriteString("'")
	}

	var totalCount int
	err := o.Raw(buildStr.String()).QueryRow(&totalCount)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return totalCount
}
