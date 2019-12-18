package cmd

import (
	"bytes"
	"fmt"
	"github.com/gofuncchan/ginger-gen/util"
	"github.com/urfave/cli"
	"io"
	"text/template"
)

type modelTmplData struct {
	FuncName string
}

// 创建mysql model方法
var ModelCommand = cli.Command{
	Name:        "model",
	Usage:       "generate biz logic model function code for mysql builder",
	UsageText:   "ginger-gen model [option]",
	Description: "generate biz logic model function code for mysql builder",
	Flags: []cli.Flag{
		cli.StringFlag{Name: "root, r",Value:"github.com/gofuncchan/ginger", Usage: "root package name"},
		cli.StringFlag{Name: "module, m", Usage: "module name",},
		cli.StringSliceFlag{Name: "func, f", Usage: "model function name,one or more"},
	},
	Action: modelCommandAction,
}

func modelCommandAction(c *cli.Context) error {
	root := c.String("root")
	module := c.String("module")
	fs := c.StringSlice("func")

	var buffs bytes.Buffer
	for _, f := range fs {
		var buff bytes.Buffer
		// handler函数模板
		err := template.Must(template.New("model").Parse(modelTmplCode)).Execute(&buff, modelTmplData{
			FuncName: f,
		})
		if err != nil {
			return err
		}
		_, err = io.Copy(&buffs, &buff)
		if err != nil {
			return err
		}
	}

	// 设置输出
	out, err := util.OutputFile(root,util.ModelPkgName, module)
	if err != nil {
		util.OutputWarn(err.Error())
	}

	_, err = io.Copy(out, &buffs)
	if err != nil {
		return err
	}

	// stdout 输出router代码设置
	util.OutputInfo("Generate Successful", outputModelTips(fs))

	return nil
}

func outputModelTips(funcNames []string) string {
	header := "Please reset input or output params of model function.\n"
	footer := "You should handling errors in model function,and return data or result to caller.\n"
	examples := "For example:\n"

	for _, f := range funcNames {
		examples += fmt.Sprintf(`
		func %s(arg1 int, arg2 string) int64 {
			// TODO model logic code
			// ...
			}
	`, f)
	}

	return header + examples + footer
}


const modelTmplCode = `
/*
TODO Description: What does {{ .FuncName }}  do ...
TODO reset this function input args and output args
@param

@return
*/
func {{ .FuncName }}(args... interface{}) (b bool) {
    // TODO your model logic code for mysql store ...

    return
}
`