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
	app.Name = "ginger-cli"
	app.Version = "0.0.1"
	app.Compiled = time.Now()
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
		cmd.BuilderCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}


