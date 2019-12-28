package writer

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// 输出
func OutputCacheFile(root, module string) (io.Writer, error) {
	packName := module + "Cache"

	// 创建输出目录
	err := os.MkdirAll("cache/"+packName, 0755)
	if err != nil {
		// 如目录创建失败，则标准输出
		return os.Stdout, err
	}

	// 如: /cache/tokenCache/token_cache.go
	filename := "cache/" + packName + "/" + module + "_cache.go"

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
		out = addCacheImportContent(root,packName)
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

func addCacheImportContent(root,packName string) io.Reader {
	return bytes.NewBuffer([]byte(fmt.Sprintf(`package %s

import(
	redigo "github.com/garyburd/redigo/redis"
	"%s/dao/redis"
)

/*
This code is generated with ginger-gen.
You should handling errors in cache function,and return data or result to caller.

For example:

func SetKey(key, value string) bool {
	rs, _ := redigo.String(redis.R("SET", key, value, "EX", 3600))
	return rs == "OK"
}

*/

	`, packName,root)))
}
