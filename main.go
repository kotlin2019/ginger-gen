package main

import (
	"fmt"
	"ginger-cli/cmd"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "Ginger-cli"
	app.Version = "0.0.1"
	app.Compiled = time.Now()
	app.HelpName = "help"
	app.Usage = "A client of ginger scaffold."
	app.UsageText = "ginger-cli [option] [command] [args]"
	app.ArgsUsage = "[args and such]"
	app.UseShortOptionHandling = true

	app.Action = func(c *cli.Context) error {
		fmt.Println("Ginger-cli is a client of ginger scaffold.")
		return nil
	}
	app.Commands = []cli.Command{
		cmd.InitCommand,
		cmd.MigrateCommand,
		cmd.BuilderCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
