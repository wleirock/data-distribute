package service

import (
	"fmt"
	"wleirock/data-distribute/models"

	"github.com/astaxie/beego/orm"
)

// GetAllPublicMethodList 获取机构集合
func GetAllPublicMethodList() []models.PublicMethod {
	o := orm.NewOrm()
	var list []models.PublicMethod
	_, err := o.Raw("select * from public_method").QueryRows(&list)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return list
}

// AddPublicMethod 添加publicMethod
func AddPublicMethod(publicMethod *models.PublicMethod) bool {
	o := orm.NewOrm()
	res, err := o.Raw("insert into public_method(method_name,data_type,script_name,description) values (?,?,?,?)", publicMethod.MethodName, publicMethod.DataType, publicMethod.ScriptName, publicMethod.Description).Exec()
	if err != nil {
		fmt.Println("AddPublicMethod error", err)
		return false
	}

	count, err := res.RowsAffected()
	if err != nil {
		fmt.Println("AddPublicMethod error", err)
		return false
	}

	if count > 0 {
		return true
	}
	return false
}

// UpdatePublicMethod 更新publicMethod
func UpdatePublicMethod(publicMethod *models.PublicMethod) bool {
	o := orm.NewOrm()
	res, err := o.Raw("update public_method set method_name = ?,data_type = ?,script_name = ?,description = ? where method_pk = ?", publicMethod.MethodName, publicMethod.DataType, publicMethod.ScriptName, publicMethod.Description, publicMethod.MethodPk).Exec()
	if err != nil {
		fmt.Println("UpdatePublicMethod error", err)
		return false
	}

	_, err = res.RowsAffected()
	if err != nil {
		fmt.Println("UpdatePublicMethod error", err)
		return false
	}

	return true
}

// GetPublicMethodByPk 根据主键查询
func GetPublicMethodByPk(methodPk int) models.PublicMethod {
	var publicMethod models.PublicMethod
	o := orm.NewOrm()
	err := o.Raw("select * from public_method where method_pk = ?", methodPk).QueryRow(&publicMethod)
	if err != nil {
		fmt.Println("GetPublicMethodByPk error", err)
	}
	return publicMethod
}

// DeletePublicMethod 根据主键删除
func DeletePublicMethod(methodPk int) bool {
	o := orm.NewOrm()
	res, err := o.Raw("delete from public_method where method_pk = ?", methodPk).Exec()
	if err != nil {
		fmt.Println("DeletePublicMethod error", err)
		return false
	}
	_, err = res.RowsAffected()
	if err != nil {
		fmt.Println("DeletePublicMethod error", err)
		return false
	}
	return true
}
