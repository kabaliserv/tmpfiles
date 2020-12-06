package main

import (
	"log"
	"os"

	"github.com/kabaliserv/tmpfiles/cmd"

	"github.com/urfave/cli/v2"
)


func main() {
	
	app := cli.NewApp()
	app.Name = "TMPFiles"
	app.Usage = "temporary file storage web application"
	app.Commands = []*cli.Command{
		{
			Name:   "start",
			Usage:  "Run App",
			Action: cmd.StartApp,
			Flags: []cli.Flag{
				cmd.FlagConf,
			},
		},
		{
			Name:   "clean",
			Usage:  "Remove expire file to filesystem",
			Action: cmd.CleanUp,
			Flags: []cli.Flag{
				cmd.FlagConf,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
