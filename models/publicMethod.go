package models

// PublicMethod 公共方法
type PublicMethod struct {
	MethodPk    int    `json:"methodPk" form:"methodPk"`
	MethodName  string `json:"methodName" form:"methodName"`
	DataType    string `json:"dataType" form:"dataType"`
	ScriptName  string `json:"scriptName" form:"scriptName"`
	Description string `json:"description" form:"description"`
}
