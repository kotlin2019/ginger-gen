package cmd

import (
	"bytes"
	forgeDao "github.com/gofuncchan/ginger-forge/lib/dao"
	forgeSchema "github.com/gofuncchan/ginger-forge/lib/schema"
	forgeUtil "github.com/gofuncchan/ginger-forge/lib/util"
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
	schemaArgs := &forgeUtil.SchemaArgs{
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
	_, err := io.Copy(&buff, forgeSchema.AddImportContent(schemaArgs.Table))
	if err != nil {
		return forgeUtil.OutputError(err.Error())
	}
	// 生成go 结构体 代码
	_, err = forgeSchema.GenSchema(&buff, schemaArgs)
	if err != nil {
		return forgeUtil.OutputError(err.Error())
	}

	// 设置输出
	out, err := forgeUtil.OutputFile(schemaArgs.Table,c.Command.Name)
	if err != nil {
		forgeUtil.OutputWarn(err.Error())
	}

	_, err = io.Copy(out, &buff)
	if err != nil {
		return forgeUtil.OutputError(err.Error())
	}

	forgeUtil.OutputOk("Generate go struct from table schema successful")

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
	// 接收参数
	schemaArgs := &forgeUtil.SchemaArgs{
		Database: c.String("database"),
		Table:    c.String("table"),
		User:     c.String("user"),
		Password: c.String("password"),
		Host:     c.String("host"),
		Port:     c.Int("port"),

	}

	// 可写缓冲
	var buff bytes.Buffer
	// 加包名和导入包代码
	_, err := io.Copy(&buff, forgeDao.AddImportContent(schemaArgs.Table))
	if err != nil {
		return forgeUtil.OutputError(err.Error())
	}

	// 生成go结构体代码
	structName, err := forgeSchema.GenSchema(&buff, schemaArgs)
	if err != nil {
		return forgeUtil.OutputError(err.Error())
	}

	// structName := forgeSchema.ConvertUnderScoreToCamel(schemaArgs.Table)
	// 生成curd方法
	err = forgeDao.GenDao(&buff,schemaArgs.Table, structName)
	if err != nil {
		return forgeUtil.OutputError(err.Error())
	}

	// 设置输出
	out, err := forgeUtil.OutputFile(schemaArgs.Table,c.Command.Name)
	if err != nil {
		forgeUtil.OutputWarn(err.Error())
	}

	_, err = io.Copy(out, &buff)
	if err != nil {
		return  forgeUtil.OutputError(err.Error())
	}
	return nil
}
