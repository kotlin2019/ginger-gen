package cmd

import (
	"fmt"
	"github.com/urfave/cli"
)

var BuilderCommand = cli.Command{
	Name:        "builder",
	Usage:       "generate dao code",
	UsageText:   "ginger-cli builder [sub-command] [option]",
	Description: "generate sql builder code for dao which fork didi/gendry",
	Subcommands: []cli.Command{
		subCommandTable,
		subCommandDao,
	},
}

var subCommandTable = cli.Command{
	Name:        "table",
	UsageText:   "",
	Description: "generate mysql table struct",
	Flags: []cli.Flag{
		cli.StringFlag{Name: "host, h", Value: "localhost"},
		cli.IntFlag{Name: "port, P", Value: 3306},
		cli.StringFlag{Name: "user, u"},
		cli.StringFlag{Name: "database, d"},
		cli.StringFlag{Name: "password, p"},
		cli.StringFlag{Name: "table, t"},
	},
	Action: subCommandTableFuncfunc,
}

func subCommandTableFuncfunc(c *cli.Context) error {
	fmt.Println("host:", c.String("host"))
	fmt.Println("port:", c.Int("port"))
	fmt.Println("user:", c.String("user"))
	fmt.Println("database:", c.String("database"))
	fmt.Println("table:", c.String("table"))
	fmt.Println("password:", c.String("password"))

	return nil
}

var subCommandDao = cli.Command{
	Name:      "dao",
	UsageText: "generate mysql table struct and CURD code",
	Flags: []cli.Flag{
		cli.StringFlag{Name: "host, h", Value: "localhost"},
		cli.IntFlag{Name: "port, P", Value: 3306},
		cli.StringFlag{Name: "user, u"},
		cli.StringFlag{Name: "database, d"},
		cli.StringFlag{Name: "password, p"},
		cli.StringFlag{Name: "table, t"},
	},
	Action: subCommandDaoFunc,
}

func subCommandDaoFunc(c *cli.Context) error {
	fmt.Println("host:", c.String("host"))
	fmt.Println("port:", c.Int("port"))
	fmt.Println("user:", c.String("user"))
	fmt.Println("database:", c.String("database"))
	fmt.Println("table:", c.String("table"))
	fmt.Println("password:", c.String("password"))
	return nil
}
