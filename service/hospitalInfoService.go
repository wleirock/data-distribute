package service

import (
	"fmt"
	"wleirock/data-distribute/models"

	"github.com/astaxie/beego/orm"
)

// GetHospitalInfoList 获取机构集合
func GetHospitalInfoList() []models.HospitalInfo {
	o := orm.NewOrm()
	var list []models.HospitalInfo
	_, err := o.Raw("select * from hospital_info").QueryRows(&list)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return list
}
