package cmd

import "github.com/urfave/cli"

// 根据config 配置文件生成配置参数解析代码
var ConfigCommand = cli.Command{
	Name:        "config",
	Usage:       "generate config loading code for yaml",
	UsageText:   "ginger-cli config [sub-command] [option]",
	Description: "generate config init code for yaml config file",
	Subcommands: []cli.Command{
		subCommandLoading,
	},
}

var subCommandLoading = cli.Command{
	Name:        "loading",
	UsageText:   "",
	Description: "generate config loading code for yaml",
	Flags: []cli.Flag{
		cli.StringFlag{Name: "path, p", Value: "config"},
	},
	Action: subCommandLoadingAction,
}

func subCommandLoadingAction(c *cli.Context) error {

	return nil
}