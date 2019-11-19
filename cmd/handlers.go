package cmd

import (
	"bytes"
	"fmt"
	"github.com/gofuncchan/ginger-cli/util"
	"github.com/urfave/cli"
	"io"
	"text/template"

)

// 创建handlers方法
var HandlerCommand = cli.Command{
	Name:        "handler",
	Usage:       "generate handler function code",
	UsageText:   "ginger-cli handler [sub-command] [option]",
	Description: "generate handler function code and request params validator struct",
	Flags: []cli.Flag{
		cli.StringFlag{Name: "name, n"},
		cli.StringFlag{Name: "func, f"},
	},
	Action: handlerCommandAction,
}

type tmplData struct {
	PackageName string
	FileName string
	FuncName string
	StructName string

}

func handlerCommandAction(c *cli.Context) error {
	fileName := c.String("name") // 生成文件名
	funcName := c.String("func") // handler函数名

	var buff bytes.Buffer
	err := template.Must(template.New("handler").Parse(handlerCode)).Execute(&buff, tmplData{
		PackageName:"handlers",
		FileName:fileName,
		StructName: funcName +"Params",
		FuncName:  funcName,
	})
	if err != nil {
		return err
	}

	// 设置输出
	out, err := util.OutputHandlerFile(util.HandlerOutputRootPath,c.Command.Name)
	if err != nil {
		util.OutputWarn(err.Error())
	}

	_, err = io.Copy(out,&buff)
	if err != nil {
		return err
	}

	// TODO stdout 输出router代码设置

	return nil
}

func AddImportContent(packageName string) io.Reader {
	return bytes.NewBuffer([]byte(fmt.Sprintf(`package %s

	`, packageName)))
}

const handlerCode = `package {{.PackageName}}
import (
\	"errors"
	"time"
	"context"
\
)

/*
This code is generated with ginger-cli
*/

// {{ .StructName }} is a mapping object for params in request
type {{.StructName}} struct {
{{- range .FieldList }}
	{{ .Name }} {{ .Type }} {{ .StructTag }}
{{- end}}
}


`