package cmd

import (
	"bytes"
	"fmt"
	"github.com/gofuncchan/ginger-forge/lib/schema"
	"github.com/gofuncchan/ginger-forge/lib/util"
	"github.com/urfave/cli"
	"io"
)

var BuilderCommand = cli.Command{
	Name:        "builder",
	Usage:       "generate dao code",
	UsageText:   "ginger-cli builder [sub-command] [option]",
	Description: "generate sql builder code for dao which fork didi/gendry",
	Subcommands: []cli.Command{
		subCommandSchema,
		subCommandDao,
	},
}

var subCommandSchema = cli.Command{
	Name:        "schema",
	UsageText:   "",
	Description: "generate mysql table schema to go struct",
	Flags: []cli.Flag{
		cli.StringFlag{Name: "host, h", Value: "localhost"},
		cli.IntFlag{Name: "port, P", Value: 3306},
		cli.StringFlag{Name: "user, u",Value:"root"},
		cli.StringFlag{Name: "password, p",Value:"123456"},
		cli.StringFlag{Name: "database, d", Required: true},
		cli.StringFlag{Name: "table, t", Required: true},
	},
	Action: subCommandSchemaAction,
}

func subCommandSchemaAction(c *cli.Context) error {
	// 接收参数
	schemaArgs := &util.SchemaArgs{
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
	_, err := io.Copy(&buff, schema.AddImportContent(schemaArgs.Table))
	if err != nil {
		return util.OutputError(err.Error())
	}
	// 生成go 结构体 代码
	_, err = schema.GenSchema(&buff, schemaArgs)
	if err != nil {
		return util.OutputError(err.Error())
	}

	// 设置输出
	out, err := util.OutputFile(schemaArgs.Table,c.Command.Name)
	if err != nil {
		util.OutputWarn(err.Error())
	}

	_, err = io.Copy(out, &buff)
	if err != nil {
		return util.OutputError(err.Error())
	}

	util.OutputOk("Generate go struct from table schema successful")

	return nil
}

var subCommandDao = cli.Command{
	Name:      "dao",
	UsageText: "generate mysql table schema to go struct and CURD code",
	Flags: []cli.Flag{
		cli.StringFlag{Name: "host, h", Value: "localhost"},
		cli.IntFlag{Name: "port, P", Value: 3306},
		cli.StringFlag{Name: "user, u",Value:"root"},
		cli.StringFlag{Name: "password, p",Value:"123456"},
		cli.StringFlag{Name: "database, d", Required: true},
		cli.StringFlag{Name: "table, t", Required: true},
	},
	Action: subCommandDaoAction,
}

func subCommandDaoAction(c *cli.Context) error {
	fmt.Println("host:", c.String("host"))
	fmt.Println("port:", c.Int("port"))
	fmt.Println("user:", c.String("user"))
	fmt.Println("password:", c.String("password"))
	fmt.Println("database:", c.String("database"))
	fmt.Println("table:", c.String("table"))
	return nil
}
