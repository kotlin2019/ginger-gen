package cmd

import (
	"bytes"
	libSchema "github.com/gofuncchan/ginger-gen/lib/schema"
	"github.com/gofuncchan/ginger-gen/xprint"
	"github.com/urfave/cli"
	"io"
)

var MysqlCommand = cli.Command{
	Name:        "schema",
	Usage:       "generate mysql table schema to go struct",
	UsageText:   "generate mysql table schema to go struct",
	Description: "generate sql builder code for dao which fork didi/gendry",
	Flags: []cli.Flag{
		cli.StringFlag{Name: "host, H", Value: "localhost"},
		cli.IntFlag{Name: "port, P", Value: 3306},
		cli.StringFlag{Name: "user, u",Value:"root"},
		cli.StringFlag{Name: "password, p",Value:"123456"},
		cli.StringFlag{Name: "database, d", Required: true},
		cli.StringFlag{Name: "table, t", Required: true},
	},
	Action: commandSchemaAction,

}

func commandSchemaAction(c *cli.Context) error {
	// 接收参数
	libSchemaArgs := &libSchema.SchemaArgs{
		Database: c.String("database"),
		Table:    c.String("table"),
		User:     c.String("user"),
		Password: c.String("password"),
		Host:     c.String("host"),
		Port:     c.Int("port"),
	}

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
	out, err := libSchema.OutputFile(libSchema.DefaultSchemaOutputRootPath,libSchemaArgs.Table)
	if err != nil {
		xprint.Warn(err.Error())
	}

	_, err = io.Copy(out, &buff)
	if err != nil {
		return xprint.Error(err.Error())
	}

	xprint.Ok("Generate go struct from table libSchema successful")

	return nil
}
