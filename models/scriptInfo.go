package models

// ScriptInfo 脚本信息
type ScriptInfo struct {
	InfoPk       int    `json:"infoPk"`
	HospitalFk   string `json:"hospitalFk"`
	HospitalName string `json:"hospitalName"`
	DataType     string `json:"dataType"`
	ScriptName   string `json:"scriptName"`
	Description  string `json:"description"`
	ApiIp        string `json:"apiIp"`
	ApiPort      int    `json:"apiPort"`
}
