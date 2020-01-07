package model

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// 输出
func OutputModelFile(root, module string) (io.Writer, error) {
	packName := module+"Model"
	// 创建输出目录
	err := os.MkdirAll("model/"+packName, 0755)
	if err != nil {
		// 如目录创建失败，则标准输出
		return os.Stdout, err
	}

	// 如: /model/userModel/user_model.go
	filename := "model/" + packName + "/" + module + "_model.go"

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
		out = addModelImportContent(root, packName)
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

func addModelImportContent(root, packName string) io.Reader {
	return bytes.NewBuffer([]byte(fmt.Sprintf(`package %s

import(
	"%s/dao/mysql/schema"
    "%s/util/e"
)

/*
This code is generated with ginger-gen.
You should handling errors in model function,and return data or result to caller.

For example:

func CreateUserByPhone(name, phone, passwd, salt string) int64 {
	var data []map[string]interface{}
	data = append(data, map[string]interface{}{
		"name":   name,
		"phone":  phone,
		"password": password,
		"salt":   salt,
	})

	id, err := builder.Insert(data)
	if !e.Em(err) {
		return -1
	}
	return id
}

*/

	`, packName,root, root)))
}
