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

// SaveScriptInfo 保存脚本信息
func SaveScriptInfo(scriptInfo *models.ScriptInfo) bool {
	o := orm.NewOrm()
	res, err := o.Raw("insert into script_info(hospital_fk,hospital_name,data_type,script_name,description,api_ip,api_port,status) values (?,?,?,?,?,?,?,?)", scriptInfo.HospitalFk, scriptInfo.HospitalName, scriptInfo.DataType, scriptInfo.ScriptName, scriptInfo.Description, scriptInfo.ApiIp, scriptInfo.ApiPort, scriptInfo.Status).Exec()
	if err != nil {
		fmt.Println("SaveScriptInfo error", err)
		return false
	}
	count, err := res.RowsAffected()
	if err != nil {
		fmt.Println("SaveScriptInfo error", err)
		return false
	}
	if count > 0 {
		return true
	}
	return false
}

// UpdateScriptInfo 更新脚本信息
func UpdateScriptInfo(scriptInfo *models.ScriptInfo) bool {
	o := orm.NewOrm()
	res, err := o.Raw("update script_info set hospital_fk = ?,hospital_name = ?,data_type = ?,script_name = ?,description = ?,api_ip = ?,api_port = ?,status = ? where info_pk = ?", scriptInfo.HospitalFk, scriptInfo.HospitalName, scriptInfo.DataType, scriptInfo.ScriptName, scriptInfo.Description, scriptInfo.ApiIp, scriptInfo.ApiPort, scriptInfo.Status, scriptInfo.InfoPk).Exec()
	if err != nil {
		fmt.Println("UpdateScriptInfo error", err)
		return false
	}
	count, err := res.RowsAffected()
	if err != nil {
		fmt.Println("UpdateScriptInfo error", err)
		return false
	}
	if count > 0 {
		return true
	}
	return false
}

// GetScriptInByPk 根据主键查询
func GetScriptInByPk(infoPk int) models.ScriptInfo {
	var scriptInfo models.ScriptInfo
	o := orm.NewOrm()
	err := o.Raw("select info_pk,hospital_fk,hospital_name,data_type,script_name,description,api_ip,api_port,status from script_info where info_pk = ?", infoPk).QueryRow(&scriptInfo)
	if err != nil {
		fmt.Println("GetScriptInByPk error", err)
	}
	return scriptInfo
}
