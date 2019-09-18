package service

import (
	"fmt"
	"wleirock/data-distribute/models"

	"github.com/astaxie/beego/orm"
)

// GetScriptInfoList 查询脚本信息集合
func GetScriptInfoList(pageIndex int, pageSize int) []models.ScriptInfo {
	o := orm.NewOrm()
	var list []models.ScriptInfo
	_, err := o.Raw("select * from script_info limit ? , ?", (pageIndex-1)*pageSize, pageSize).QueryRows(&list)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return list
}

// GetScriptInfoTotal 查询总数
func GetScriptInfoTotal() int {
	o := orm.NewOrm()
	var totalCount int
	err := o.Raw("select count(*) from script_info").QueryRow(&totalCount)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return totalCount
}
