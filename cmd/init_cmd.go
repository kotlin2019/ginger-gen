package cmd

import (
	libInit "github.com/gofuncchan/ginger-gen/lib/init"
	"github.com/gofuncchan/ginger-gen/xprint"
	"github.com/urfave/cli"
	"os"
	"runtime"
)

var InitCommand = cli.Command{
	Name:        "init",
	Usage:       "ginger app init.",
	UsageText:   "ginger-gen init [--name|-n] [project_name]",
	Description: "The init command create a new gin application in current directory，this command will generate some necessary folders and files,which make up project scaffold.",
	Flags: []cli.Flag{
		cli.StringFlag{Name: "name, n", Usage: "project name"},
		cli.BoolFlag{Name: "g", Usage: "git init"},
	},
	Action: initCommandFunc,
}

func initCommandFunc(c *cli.Context) error {
	var name string
	name = c.String("name")
	if c.NArg() > 0 {
		name = c.Args().Get(0)
	}
	if name == "" {
		name = "ginger_app"
	}

	// 检查项目目录是否已存在
	s, err := os.Stat(name)
	if err == nil && s.IsDir() {
		return xprint.Error("initialization failed, Project Directory Is Exist")
	}

	xprint.Step("Env Checking")
	if !libInit.CheckDirMode() {
		return xprint.Error("initialization failed, please check directory permissions")
	}
	xprint.Ok("Current directory is readable and writable. ")

	// 远程拉取ginger脚手架代码

	// git clone 并删除.git本地文件,让用户自己init
	xprint.Step("git clone from " + libInit.GitUrl)
	err = libInit.GitClone(name)
	if err != nil {
		return xprint.Error("initialization failed, please check your network is ok")
	} else {
		xprint.Ok("Clone ginger scaffold successful")
	}

	err = os.RemoveAll(name + "/.git")
	if err != nil {
		return xprint.Error(err.Error())
	}

	// 本地 git 初始化
	initGit := c.Bool("g")
	if initGit {
		xprint.Step("`git init`")
		InitGitCmd := "cd " + name + " && git init"
		err := libInit.ExecShellCommand(InitGitCmd)
		if err != nil {
			return xprint.Error(err.Error())
		}

		xprint.Ok("git init successful")
	}

	// 由于使用go module 管理依赖，项目内的包需要replace到本地目录，使用go mod edit 重置
	if v, err := libInit.GetMinVer(runtime.Version()); err == nil && v < 13 && v >= 11 {
		xprint.Step(runtime.Version())
		xprint.Step("`export GO111MODULE=on`")
		err = libInit.ExecShellCommand("export GO111MODULE=on")
		if err != nil {
			return xprint.Error(err.Error())
		}

		pwd, err := os.Getwd()
		pwd = pwd + "/" + name
		if err != nil {
			return xprint.Error(err.Error())
		}

		goModCmd := "cd " + name + " && go mod edit -replace github.com/gofuncchan/ginger=" + pwd
		err = libInit.ExecShellCommand(goModCmd)
		if err != nil {
			return xprint.Error(err.Error())
		}

		xprint.Step("go mod edit -replace github.com/gofuncchan/ginger=" + pwd)
		xprint.Ok("go mod replace successful")

	} else {
		return xprint.Error(err.Error())
	}
	xprint.Ok("Your project `" + name + "` set up successful")

	xprint.Info("Tips", `
	1.Because ginger uses go module to manage dependency packages, you can start with your config by default;

	2.The default root package is github.com/gofuncchan/ginger. If you need to change it, please replace it globally and modify the go.mod file;

	3.Once the root package is replaced, most of the commands of the tool will take the -r parameter to set your custom root package, so that the generated code is consistent with your project;

	4.If the IDE does not recognize the package, please use go mod tidy and go mod vendor for localization,

	5.Use the -g option to initialize git automatically when init
`)
	return nil
}
