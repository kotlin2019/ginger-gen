package cmd

import (
	"github.com/urfave/cli"
)

// 根据config 配置文件生成配置参数解析代码
var ConfigCommand = cli.Command{
	Name:        "config",
	Usage:       "generate config parse code for yaml file ",
	UsageText:   "ginger-cli config [sub-command] [option]",
	Description: "generate config init code for yaml config file",
	Subcommands: []cli.Command{
		subCommandParse,
	},
}

var subCommandParse = cli.Command{
	Name:        "parse",
	UsageText:   "",
	Description: "generate config parse code for yaml",
	Flags: []cli.Flag{
		cli.StringFlag{Name: "path, p", Value: "config"},
	},
	Action: subCommandParseAction,
}

func subCommandParseAction(c *cli.Context) error {
	return nil
}