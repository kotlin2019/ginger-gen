package cmd

import (
	"fmt"
	"github.com/urfave/cli"
)

var MigrateCommand = cli.Command{
	Name:        "migrate",
	Usage:       "generate database scheme.",
	UsageText:   "ginger-cli migrate [sql file]",
	Description: "Import database according to SQL script file.",
	Flags: []cli.Flag{
		cli.StringFlag{Name: "file, f"},
	},
	Action: migrateCommandFunc,
}

func migrateCommandFunc(c *cli.Context) error {
	fmt.Println("file:", c.String("file"))
	return nil
}
