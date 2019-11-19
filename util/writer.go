package util

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

const HandlerOutputRootPath = "handlers/"

// 输出
func OutputHandlerFile(rootPath, handlerName string) (io.Writer, error) {
	// 创建输出目录
	err := os.MkdirAll(rootPath, 0755)
	if err != nil {
		// 如目录创建失败，则标准输出
		return os.Stdout, err
	}

	// 默认于./dao/{table_name}/schema.go ,如./dao/user/schema.go
	filename := rootPath + handlerName + ".go"


	// 创建输出的目录并创建输出的go文件
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	// 如存在，直接输出该文件
	if err == nil{
		OutputInfo("OpenFile","File is exist,the code will append to this file.")
		return file,nil
	}

	// 不存在，创建文件并加文件头
	var std = os.Stdout
	if os.IsNotExist(err) {
		OutputInfo("OpenFile","File is not exist,create new file.")
		out := addImportContent("handler")

		file, err = os.Create(filename)
		if err != nil {
			io.Copy(std,out)
			return std, err
		}

		io.Copy(file, out)
		return file,nil
	}

	// 有其它错误，则标准输出
	return std, err
}


func addImportContent(packageName string) io.Reader {
	return bytes.NewBuffer([]byte(fmt.Sprintf(`package %s

import(
    "github.com/gofuncchan/ginger/common"
    "github.com/gin-gonic/gin"
)

/*
This code is generated with ginger-cli.
You must reset Request Params, and implement biz logic code.
*/

	`, packageName)))
}