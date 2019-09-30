package service

import (
	"fmt"
	"strings"
	"wleirock/data-distribute/models"

	"github.com/astaxie/beego/orm"
)

// GetScriptFileList 查询脚本信息集合
func GetScriptFileList(scriptFile *models.ScriptFile) []models.ScriptFile {
	o := orm.NewOrm()

	var buildStr strings.Builder
	buildStr.WriteString("select * from script_file where 1=1")
	if scriptFile.FileName != "" {
		buildStr.WriteString(" and file_name like '%")
		buildStr.WriteString(scriptFile.FileName)
		buildStr.WriteString("%'")
	}
	buildStr.WriteString(" order by file_pk desc")
	buildStr.WriteString(" limit ? , ?")

	var list []models.ScriptFile
	_, err := o.Raw(buildStr.String(), (scriptFile.Page-1)*scriptFile.Limit, scriptFile.Limit).QueryRows(&list)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return list
}

// GetScriptFileTotal 查询总数
func GetScriptFileTotal(scriptFile *models.ScriptFile) int {
	o := orm.NewOrm()

	var buildStr strings.Builder
	buildStr.WriteString("select count(*) from script_file where 1=1")
	if scriptFile.FileName != "" {
		buildStr.WriteString(" and file_name like '%")
		buildStr.WriteString(scriptFile.FileName)
		buildStr.WriteString("%'")
	}

	var totalCount int
	err := o.Raw(buildStr.String()).QueryRow(&totalCount)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return totalCount
}

// AddScriptFile 保存脚本信息
func AddScriptFile(scriptFile *models.ScriptFile) bool {
	o := orm.NewOrm()
	res, err := o.Raw("insert into script_file(file_name,file_size,status,upload_time) values (?,?,?,?)", scriptFile.FileName, scriptFile.FileSize, scriptFile.Status, scriptFile.UploadTime).Exec()
	if err != nil {
		fmt.Println("AddScriptFile error", err)
		return false
	}
	count, err := res.RowsAffected()
	if err != nil {
		fmt.Println("AddScriptFile error", err)
		return false
	}
	if count > 0 {
		return true
	}
	return false
}

// UpdateScriptFile 更新脚本信息
func UpdateScriptFile(scriptFile *models.ScriptFile) bool {
	o := orm.NewOrm()
	res, err := o.Raw("update script_file set file_name = ?,file_size = ?,status = ?,upload_time = ? where file_pk = ?", scriptFile.FileName, scriptFile.FileSize, scriptFile.Status, scriptFile.UploadTime, scriptFile.FilePk).Exec()
	if err != nil {
		fmt.Println("UpdateScriptFile error", err)
		return false
	}
	_, err = res.RowsAffected()
	if err != nil {
		fmt.Println("UpdateScriptFile error", err)
		return false
	}
	return true
}

// GetScriptFileByPk 根据主键查询
func GetScriptFileByPk(filePk int) models.ScriptFile {
	var scriptFile models.ScriptFile
	o := orm.NewOrm()
	err := o.Raw("select file_pk,file_name,file_size,status,upload_time from script_file where file_pk = ?", filePk).QueryRow(&scriptFile)
	if err != nil {
		fmt.Println("GetScriptFileByPk error", err)
	}
	return scriptFile
}

// DeleteScriptFile 根据主键删除脚本信息
func DeleteScriptFile(filePk int) bool {
	o := orm.NewOrm()
	res, err := o.Raw("delete from script_file where file_pk = ?", filePk).Exec()
	if err != nil {
		fmt.Println("DeleteScriptFile error", err)
		return false
	}
	_, err = res.RowsAffected()
	if err != nil {
		fmt.Println("DeleteScriptFile error", err)
		return false
	}
	return true
}

// GetScriptFileByName 根据文件名称查询
func GetScriptFileByName(fileName string) models.ScriptFile {
	var scriptFile models.ScriptFile
	o := orm.NewOrm()
	err := o.Raw("select file_pk,file_name,file_size,status,upload_time from script_file where file_name = ?", fileName).QueryRow(&scriptFile)
	if err != nil {
		fmt.Println("GetScriptFileByName error", err)
	}
	return scriptFile
}
