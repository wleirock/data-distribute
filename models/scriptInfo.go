package models

// ScriptInfo 脚本信息
type ScriptInfo struct {
	InfoPk       int    `json:"infoPk" form:"-"`
	HospitalFk   string `json:"hospitalFk" form:"hospitalFk"`
	HospitalName string `json:"hospitalName" form:"hospitalName"`
	DataType     string `json:"dataType"`
	ScriptName   string `json:"scriptName" form:"scriptName"`
	Description  string `json:"description"`
	ApiIp        string `json:"apiIp"`
	ApiPort      int    `json:"apiPort"`
	Page         int    `form:"page"`
	Limit        int    `form:"limit"`
}
