package util

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// 检查目录权限
func CheckDirMode() bool {
	OutputStep("Env Checking")
	// 获取当前目录
	dir, err := os.Getwd()
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	// 目录是否可读可写
	err = syscall.Access(dir, syscall.O_RDWR)
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {
		OutputOk("Current directory is readable and writable. ")
		return true
	}
}

const GitUrl = "https://github.com/gofuncchan/ginger.git"

func GitClone(appName string) bool {
	shellCmd := "`git clone " + GitUrl + " " + appName + "`"
	OutputStep(shellCmd)
	err, _, _ := ExecShellCommand(shellCmd)
	if err != nil {
		OutputError(err.Error())
		return false
	}
	OutputOk("Clone ginger scaffold successful")

	return true
}

const ShellToUse = "bash"

func ExecShellCommand(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

// 判断目录是否存在
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		OutputError(err.Error())
		return false
	}
	return s.IsDir()
}
