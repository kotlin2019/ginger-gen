package cmd

import (
	"github.com/gofuncchan/ginger-cli/util"
	"github.com/urfave/cli"
	"os"
)

var InitCommand = cli.Command{
	Name:        "init",
	Usage:       "ginger app init.",
	UsageText:   "ginger-cli init [--name|-n] [project_name]",
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
	isExist := util.IsDir(name)
	if isExist {
		return util.OutputError("initialization failed, Project Directory Is Exist")
	}

	if !util.CheckDirMode() {
		return util.OutputError("initialization failed, please check directory permissions")
	}

	// 远程拉取ginger脚手架代码
	done := util.GitClone(name)
	if !done {
		return util.OutputError("initialization failed, please to check you already install git and network is ok")
	}
	// 删除.git本地文件,让用户自己init
	err := os.RemoveAll(name + "/.git")
	if err != nil {
		return util.OutputError(err.Error())
	}

	initGit := c.Bool("g")
	if initGit {
		util.OutputStep("`git init`")
		InitGitCmd := "cd " + name + " && git init"
		err := util.ExecShellCommand(InitGitCmd)
		if err != nil {
			return util.OutputError(err.Error())
		}

		util.OutputOk("git init successful")
	}

	// 由于使用go module 管理依赖，项目内的包需要replace到本地目录，使用go mod edit 重置
	pwd, err := os.Getwd()
	pwd = pwd + "/" + name
	if err != nil {
		return util.OutputError(err.Error())
	}

	goModCmd := "cd " + name + " && go mod edit -replace github.com/gofuncchan/ginger=" + pwd
	err = util.ExecShellCommand(goModCmd)
	if err != nil {
		return util.OutputError(err.Error())
	}

	util.OutputStep("go mod edit -replace github.com/gofuncchan/ginger=" + pwd)
	util.OutputOk("Your project `" + name + "` set up successful")

	util.OutputInfo("Tips",`
	
`)
	return nil
}
