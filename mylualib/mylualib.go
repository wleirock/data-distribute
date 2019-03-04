package mylualib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

// Loader 装载函数
func Loader(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), exports)
	// returns the module
	L.Push(mod)
	return 1
}

var exports = map[string]lua.LGFunction{
	"sendHTTPGet":  sendHTTPGet,
	"sendHTTPPost": sendHTTPPost,
}

func sendHTTPGet(L *lua.LState) int {
	url := L.CheckString(1)
	res, err := http.Get(url)
	if err != nil {
		L.Push(lua.LString("http error"))
		return 1
	}
	defer res.Body.Close()

	getRes, getErr := ioutil.ReadAll(res.Body)
	if getErr != nil {
		fmt.Println(getErr)
		L.Push(lua.LString("read param error"))
		return 1
	}

	resultStr := string(getRes)
	L.Push(lua.LString(resultStr))
	return 1 //此处是返回值个数，并不是返回值本身
}

func sendHTTPPost(L *lua.LState) int {
	url := L.CheckString(1)
	param := L.CheckString(2)
	contentType := L.CheckString(3)
	resp, err := http.Post(url, contentType, strings.NewReader(param))
	if err != nil {
		fmt.Println("error:", err)
		L.Push(lua.LString("http error"))
		return 1
	}
	defer resp.Body.Close()

	getRes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
		L.Push(lua.LString("read param error"))
		return 1
	}
	resultStr := string(getRes)
	L.Push(lua.LString(resultStr))
	return 1
}
