package utils

import (
	"fmt"
	"io/ioutil"
)

// LuaFile lua文件
type LuaFile struct {
	FileName string
}

// LuaFileExist 判断某lua脚本是否存在
func LuaFileExist(fileName string) bool {
	fileList := GetAllLuaFile()
	for _, file := range fileList {
		if file.FileName == fileName {
			return true
		}
	}
	return false
}

// GetAllLuaFile 获取所有文件
func GetAllLuaFile() []LuaFile {
	// 声明一个文件列表切片
	var fileList []LuaFile
	rd, err := ioutil.ReadDir("luafile")
	if err != nil {
		fmt.Println("GetAllLuaFile error", err)
		return fileList
	}

	for _, fi := range rd {
		fileList = append(fileList, LuaFile{fi.Name()})
	}
	return fileList
}
