package util

import (
	"io"
	"os"
)

const HandlerOutputRootPath = "handlers/"

// 输出
func OutputHandlerFile(rootPath, handlerName string) (io.Writer, error) {
	// 创建输出目录
	err := os.MkdirAll(rootPath, 0755)
	if err != nil {
		// 如文件不能输出，则标准输出
		return os.Stdout, err
	}

	// 默认于./dao/{table_name}/schema.go ,如./dao/user/schema.go
	filename := rootPath + handlerName + ".go"

	// 创建输出的目录并创建输出的go文件
	file, err := os.OpenFile(filename, os.O_RDWR, 0755)
	if os.IsNotExist(err) {
		file, err = os.Create(filename)
		if err != nil {
			return os.Stdout, err
		}
	}

	return file, nil
}
