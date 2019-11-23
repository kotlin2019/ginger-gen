package cmd

// 生成mongo dao代码的命令集合
import (
	"bytes"
	"fmt"
	"github.com/gofuncchan/ginger-cli/util"
	"github.com/urfave/cli"
	"io"
	"text/template"
)

// 创建mongo repo方法
var RepoCommand = cli.Command{
	Name:        "repo",
	Usage:       "generate repo file and function code for mongodb repository",
	UsageText:   "ginger-cli repo [option]",
	Description: "generate repo file and function code for mongodb repository",
	Flags: []cli.Flag{
		cli.StringFlag{Name: "module, m", Usage: "repo file name",},
		cli.StringSliceFlag{Name: "func, F", Usage: "repo function name,one or more"},
	},
	Action: repoCommandAction,
}

type repoTmplData struct {
	ModuleName     string
	CollectionName string
	FuncName       string
}

func repoCommandAction(c *cli.Context) error {

	module := c.String("module")
	fs := c.StringSlice("func")

	var buffs bytes.Buffer
	for _, f := range fs {
		var buff bytes.Buffer
		// handler函数模板
		err := template.Must(template.ParseFiles("./tmpl/repo.tmpl")).Execute(&buff, repoTmplData{
			ModuleName:     util.CamelString(module),
			CollectionName: module,
			FuncName:       f,
		})
		if err != nil {
			return err
		}
		io.Copy(&buffs, &buff)
	}

	// 设置输出
	out, err := util.OutputFile(util.RepoOutput, module)
	if err != nil {
		util.OutputWarn(err.Error())
	}

	_, err = io.Copy(out, &buffs)
	if err != nil {
		return err
	}

	// stdout 输出router代码设置
	util.OutputInfo("Generate Successful", outputRepoTips(module, fs))

	return nil
}

func outputRepoTips(moduleName string, funcNames []string) string {
	header := "Please reset input or output params of repository function.\n"
	footer := "You should handle errors in repository function,and return data or result to caller.\n"
	examples := "For example:\n"

	for _, f := range funcNames {
		examples += fmt.Sprintf(`
		func %s(arg1 int, arg2 string) int64 {
			// TODO model logic code
			// ...
			}
	`,  f)
	}

	return header + examples + footer

}
