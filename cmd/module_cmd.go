package cmd

import (
	"bytes"
	libModule "github.com/gofuncchan/ginger-gen/lib/module"
	libSchema "github.com/gofuncchan/ginger-gen/lib/schema"
	"github.com/gofuncchan/ginger-gen/util"
	"github.com/gofuncchan/ginger-gen/xprint"
	"github.com/urfave/cli"
	"strconv"
	"text/template"

	"io"
)

// 生成curd模块
var ModuleCommand = cli.Command{
	Name:        "module",
	Usage:       "generate CURD actions for a module",
	UsageText:   "ginger-gen module [option]",
	Description: "generate CURD actions for a module",
	Flags: []cli.Flag{
		cli.StringFlag{Name: "root, r", Value: "github.com/gofuncchan/ginger", Usage: "root package name"},
		cli.StringFlag{Name: "host, H", Value: "localhost"},
		cli.IntFlag{Name: "port, P", Value: 3306},
		cli.StringFlag{Name: "user, u", Value: "root"},
		cli.StringFlag{Name: "password, p", Value: "123456"},
		cli.StringFlag{Name: "database, d", Required: true},
		cli.StringFlag{Name: "table, t", Required: true},
	},
	Action: moduleCommandAction,
}

func moduleCommandAction(c *cli.Context) error {
	// 接收参数
	libSchemaArgs := &libSchema.SchemaArgs{
		Database: c.String("database"),
		Table:    c.String("table"),
		User:     c.String("user"),
		Password: c.String("password"),
		Host:     c.String("host"),
		Port:     c.Int("port"),
	}

	root := c.String("root")
	module := c.String("table")

	// 1.先连接数据库生成schema struc
	// 可写缓冲
	var buff bytes.Buffer
	// 加包名
	_, err := io.Copy(&buff, libSchema.AddImportContent(libSchemaArgs.Table))
	if err != nil {
		return xprint.Error(err.Error())
	}
	// 生成go 结构体 代码
	_, err = libSchema.GenSchemaStruct(&buff, libSchemaArgs)
	if err != nil {
		return xprint.Error(err.Error())
	}

	// 设置输出
	out, err := libSchema.OutputFile(libSchema.DefaultSchemaOutputRootPath, libSchemaArgs.Table)
	if err != nil {
		xprint.Warn(err.Error())
	}

	_, err = io.Copy(out, &buff)
	if err != nil {
		return xprint.Error(err.Error())
	}

	// 2.生成相应的 curd handler/model
	// 获取表字段，作为curd handler传参
	columnSlices, err := libSchema.GetSchemaField(libSchemaArgs)
	if err != nil {
		return xprint.Error(err.Error())
	}

	// handler
	var handlerBuff bytes.Buffer
	err = template.Must(template.New("module_handler").Parse(libModule.ModuleHandlerTmplCode)).Execute(&handlerBuff, libModule.ModuleHandlerTmplData{
		RootPackageName: root,
		ModuleName:      module,
		CamelModuleName: util.CamelString(module),
		GetListParams:   libModule.GenGetListParams(),
		GetOneParams:    libModule.GenGetOneParams(),
		CreateParams:    libModule.GenCreateParams(columnSlices),
		UpdateParams:    libModule.GenUpdateParams(columnSlices),
		DeleteParams:    libModule.GenDeleteParams(),
		CreateArgs:      libModule.GenCreateArgs(columnSlices),
		UpdateArgs:      libModule.GenUpdateArgs(columnSlices),
	})
	if err != nil {
		return xprint.Error(err.Error())
	}

	handlerWriter, err := libModule.OutputHandlerFile(module)
	if err != nil {
		return xprint.Error(err.Error())
	}
	written, err := io.Copy(handlerWriter, &handlerBuff)
	if err != nil {
		return xprint.Error(err.Error())
	}
	xprint.Ok("handler written:" + strconv.Itoa(int(written)))

	// model
	var modelBuff bytes.Buffer
	err = template.Must(template.New("module_model").Parse(libModule.ModuleModelTmplCode)).Execute(&modelBuff, libModule.ModuleModelTmplData{
		RootPackageName:  root,
		ModuleName:       module,
		CamelModuleName:  util.CamelString(module),
		CreateArgs:       libModule.GenCreateModelArgs(columnSlices),
		CreateArgsForMap: libModule.GenCreateArgsForMap(columnSlices),
	})
	if err != nil {
		return xprint.Error(err.Error())
	}

	modelWriter, err := libModule.OutputModuleModelFile(module)
	if err != nil {
		return xprint.Error(err.Error())
	}
	i, err := io.Copy(modelWriter, &modelBuff)
	if err != nil {
		return xprint.Error(err.Error())
	}
	xprint.Ok("model written:" + strconv.Itoa(int(i)))
	xprint.Ok("Generate module CURD successful")

	// 输出需添加的routes
	xprint.Info("Tips",libModule.PrintHandlerTips(module))

	return nil
}
