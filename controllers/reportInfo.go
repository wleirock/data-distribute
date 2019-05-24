package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/yuin/gopher-lua"
	"wleirock/data-distribute/mylualib"
)

// ReportInfoController reportInfo api
type ReportInfoController struct {
	beego.Controller
}
// Distribute send reportInfo to other platform
func (c *ReportInfoController) Distribute() {
	// do job, each platform has its own lua script file
	go processLua("golib.lua")
	resultData := ResponseEntity{"0", true, "success"}
	c.Data["json"] = &resultData
	c.ServeJSON()
}

func processLua(luaFileName string) string {
	resStr := "error"
	// L := lua.NewState()
	// defer L.Close()
	L := luaPool.Get()
	defer luaPool.Put(L)

	L.PreloadModule("mymodule", mylualib.Loader)

	//加载lua文件
	if err := L.DoFile("luafile/" + luaFileName); err != nil {
		fmt.Println(err)
	}
	//调用lua中的函数
	err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("sendReport"), //获取函数引用
		NRet:    1,                         //指定返回值数量
		Protect: true,                      //如果出现异常是panic还是返回err
	}, lua.LString("http://192.168.1.93:8888/index"), lua.LString("{'checkNo':'100210','time':'2019-02-22'}")) //传递参数
	if err != nil {
		fmt.Println(err)
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
		fmt.Println("出现异常")
	}
	return resStr
}