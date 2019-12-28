package writer

import (
	"bytes"
	"fmt"
	"github.com/gofuncchan/ginger-gen/util"
	"io"
	"os"
)


// 输出
func OutputRepoFile(root, module string) (io.Writer, error) {
	// 创建输出目录
	err := os.MkdirAll("repository/"+module+"Repo", 0755)
	if err != nil {
		// 如目录创建失败，则标准输出
		return os.Stdout, err
	}

	// 如: /repository/logRepo/log_repository.go
	filename := "repository/" + module + "Repo/" + module + "_repository.go"

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
		out = addRepoImportContent(root, module)
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


func addRepoImportContent(root,module string) io.Reader {
	return bytes.NewBuffer([]byte(fmt.Sprintf(`package repository

import(
   	mongo "%s/dao/mongodb"
	"%s/util/e"
	"gopkg.in/mgo.v2"
)

/*
This code is generated with ginger-gen.
You should handling errors in repository function,and return data or result to caller.

For example:

func InsertPost(dataMap map[string]interface{}) (b bool) {

	// Use mongodb dao common function
	err := mongo.Insert(MongoPostCollection,dataMap)
	
	if !e.Em(err) {
		return false
	}

	return true
}
*/

const Mongo%sCollection = "%s"

	`,root,root,util.CamelString(module),module)))
}
