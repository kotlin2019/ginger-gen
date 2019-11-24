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
		cli.StringSliceFlag{Name: "func, f", Usage: "repo function name,one or more"},
	},
	Action: repoCommandAction,
}

type repoTmplData struct {
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
		err := template.Must(template.New("repository").Parse(repoTmplCode)).Execute(&buff, repoTmplData{
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
	util.OutputInfo("Generate Successful", outputRepoTips(fs))

	return nil
}

func outputRepoTips(funcNames []string) string {
	header := "Please reset input or output params of repository function.\n"
	footer := "You should handling errors in repository function,and return data or result to caller.\n"
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

const repoTmplCode  = `
/*
TODO Description: What does {{ .FuncName }}  do ...
TODO reset this function input args and output args
@param

@return
*/
func {{ .FuncName }}(args... interface{}) (err error) {
    // TODO your repository logic code for mongodb store ...

    return nil
}
`