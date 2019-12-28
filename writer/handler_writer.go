package writer

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// 输出
func OutputHandlerFile(root, module string) (io.Writer, error) {
	// 创建输出目录
	err := os.MkdirAll("handler", 0755)
	if err != nil {
		// 如目录创建失败，则标准输出
		return os.Stdout, err
	}

	filename := "handler/" + module + "_handler.go"

	// 创建输出的目录并创建输出的go文件
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	// 如存在，直接输出该文件
	if err == nil {
		return file, nil
	}

	// 不存在，创建文件并加文件头
	var std = os.Stdout
	var out io.Reader
	if os.IsNotExist(err) {
		out = addHandlerImportContent(root)
		file, err = os.Create(filename)
		if err != nil {
			io.Copy(std, out)
			return std, err
		}
		io.Copy(file, out)
		return file, nil
	}

	// 有其它错误，则标准输出
	return std, err
}

func addHandlerImportContent(root string) io.Reader {
	return bytes.NewBuffer([]byte(fmt.Sprintf(`package handler

import(
    "%s/common"
    "github.com/gin-gonic/gin"
)

/*
This code is generated with ginger-gen.
You must reset Request Params, and implement biz logic code.
*/

	`, root)))
}
