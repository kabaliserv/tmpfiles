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
	app.Flags = []cli.Flag{
		cmd.FlagConf,
	}
	app.Action = cmd.StartApp

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
