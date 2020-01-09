package module

import (
	"io"
	"os"
)

// 输出
func OutputModuleModelFile(module string) (io.Writer, error) {
	packName := module + "Model"
	// 创建输出目录
	var err error
	_ = os.MkdirAll("model/"+packName, 0755)

	// 如: /model/userModel/user_model.go
	filename := "model/" + packName + "/" + module + "_model.go"

	// 创建输出的目录并创建输出的go文件
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	// 如存在，直接输出该文件
	if err == nil {
		return file, nil
	}

	// 不存在，创建文件并加文件头
	if os.IsNotExist(err) {
		file, err = os.Create(filename)
		if err != nil {
			return os.Stdout, err
		}
		return file, nil
	}

	// 有其它错误，则标准输出
	return os.Stdout, err
}
