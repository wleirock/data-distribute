package models

// ScriptInfo 脚本信息
type ScriptInfo struct {
	InfoPk       int    `json:"infoPk" form:"infoPk"`
	HospitalFk   string `json:"hospitalFk" form:"hospitalFk"`
	HospitalName string `json:"hospitalName" form:"hospitalName"`
	DataType     string `json:"dataType" form:"dataType"`
	ScriptName   string `json:"scriptName" form:"scriptName"`
	Description  string `json:"description" form:"description"`
	ApiIp        string `json:"apiIp" form:"apiIp"`
	ApiPort      int    `json:"apiPort" form:"apiPort"`
	Status       string `json:"status" form:"status"`
	Page         int    `form:"page"`
	Limit        int    `form:"limit"`
}
