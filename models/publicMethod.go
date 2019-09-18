package models

// PublicMethod 公共方法
type PublicMethod struct {
	MethodPk    int    `json:"methodPk"`
	MethodName  string `json:"methodName"`
	DataType    string `json:"dataType"`
	ScriptName  string `json:"scriptName"`
	Description string `json:"description"`
}
