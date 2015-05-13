package main

import (
	"github.com/codegangsta/cli"
	"github.com/zchee/golang-sandbox/commands"
	"github.com/zchee/golang-sandbox/version"
	"os"
	"path"
)

func main() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Version = version.VERSION + " (" + version.GITCOMMIT + ")"
	app.Usage = "golang sandbox library"
	app.Author = "zchee"
	app.Email = "zcheeee@gmail.com"
	app.Commands = commands.Commands

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "Enable debug mode",
		},
	}

	app.Run(os.Args)
}
