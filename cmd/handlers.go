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
		cli.StringFlag{Name: "file, f",Usage:"handler file name",},
		cli.StringSliceFlag{Name: "func, F",Usage:"handler function name,one or more"},
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
	fileName := c.String("file") // 生成文件名
	// funcName := c.String("func") // handler函数名
	fs := c.StringSlice("func")


	var buffs bytes.Buffer
	for _,f := range fs {
		var buff bytes.Buffer
		// handler函数模板
		err := template.Must(template.ParseFiles("./tmpl/handler.tmpl")).Execute(&buff, tmplData{
			PackageName:"handler",
			FileName:fileName,
			StructName: f +"Params",
			FuncName:  f,
		})
		if err != nil {
			return err
		}
		io.Copy(&buffs,&buff)
	}


	// 设置输出
	out, err := util.OutputHandlerFile(util.HandlerOutputRootPath,fileName)
	if err != nil {
		util.OutputWarn(err.Error())
	}

	_, err = io.Copy(out,&buffs)
	if err != nil {
		return err
	}

	// stdout 输出router代码设置
	util.OutputInfo("Generate Successful",outputTips(fileName,fs))

	return nil
}

func outputTips(fileName string ,funcNames []string) string {
	var header = `binding this handler function to app router.
	For example:`
	var footer = `According to the http method what you need,copy the code to router/router.go.`
	var examples string
	for _,f := range funcNames{
		snake := util.SnakeString(f)
		examples += fmt.Sprintf(`
		r.POST("%s/%s", handler.%s)
		
	`,fileName,snake, f)
	}

	return header + examples + footer

}


const routerCode  =  `


	r.POST("/signup", handlers.SignUp)

`