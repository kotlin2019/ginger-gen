package main

import (
	"fmt"
	"github.com/gofuncchan/ginger-cli/cmd"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "ginger-cli"
	app.Version = "0.2.0"
	app.Compiled = time.Now()
	app.Usage = "A client of ginger scaffold."
	app.UsageText = "ginger-cli [option] [command] [args]"
	app.ArgsUsage = "[args and such]"
	app.UseShortOptionHandling = true

	app.Action = func(c *cli.Context) error {
		fmt.Println("ginger-cli is a client of ginger scaffold.")
		return nil
	}
	app.Commands = []cli.Command{
		cmd.InitCommand,
		cmd.MysqlCommand,
		cmd.HandlerCommand,
		cmd.ModelCommand,
		cmd.RepoCommand,
		cmd.CacheCommand,
		cmd.ConfigCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
