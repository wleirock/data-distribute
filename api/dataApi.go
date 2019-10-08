package api

import (
	"fmt"

	"wleirock/data-distribute/mylualib"

	"github.com/astaxie/beego"
	lua "github.com/yuin/gopher-lua"
)

// DataAPIController 接收数据并根据配置分发
type DataAPIController struct {
	beego.Controller
}

// ResponseEntity 定义返回的JSON字段
type ResponseEntity struct {
	Code    string `json:"code"`
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

// ErrorJSON 返回错误数据
func (c *DataAPIController) ErrorJSON(code string, msg string, err error) {
	if err != nil {
		fmt.Println(err)
	}
	resultData := ResponseEntity{code, false, msg}
	c.Data["json"] = &resultData
	c.ServeJSON()
}

// SuccessJSON 返回正确数据
func (c *DataAPIController) SuccessJSON() {
	resultData := ResponseEntity{"0", true, "success"}
	c.Data["json"] = &resultData
	c.ServeJSON()
}

// ProcessLua 调用lua文件进行数据分发
func (c *DataAPIController) ProcessLua(luaFileName string, url string, param string) (string, error) {
	resStr := "error"
	// L := lua.NewState()
	// defer L.Close()
	L := luaPool.Get()
	defer luaPool.Put(L)

	L.PreloadModule("mymodule", mylualib.Loader)

	//加载lua文件
	if err := L.DoFile("luafile/" + luaFileName); err != nil {
		fmt.Println(err)
		return "error", err
	}
	//调用lua中的函数
	err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("sendData"), //获取函数引用
		NRet:    1,                       //指定返回值数量
		Protect: true,                    //如果出现异常是panic还是返回err
	}, lua.LString(url), lua.LString(param)) //传递参数
	if err != nil {
		fmt.Println(err)
		return "error", err
	}
	// 获取返回结果
	result := L.Get(-1)
	//从堆栈中扔掉返回结果
	L.Pop(1)
	//打印结果
	res, ok := result.(lua.LString)
	if ok {
		fmt.Println(res)
		resStr = "success"
	} else {
		resStr = "false"
	}
	return resStr, nil
}
