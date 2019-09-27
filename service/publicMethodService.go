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
