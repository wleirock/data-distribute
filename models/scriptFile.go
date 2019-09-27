package models

// ScriptFile 脚本文件类型
type ScriptFile struct {
	FilePk     int    `json:"filePk" form:"filePk"`
	FileName   string `json:"fileName" form:"fileName"`
	FileSize   int64  `json:"fileSize"`
	Status     string `json:"status"`
	UploadTime string `json:"uploadTime"`
	Page       int    `form:"page" form:"page"`
	Limit      int    `form:"limit" form:"limit"`
}
