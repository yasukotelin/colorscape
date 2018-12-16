package main

import (
	"os"

	"github.com/urfave/cli"
)

const (
	appName = "colorscape"
	usage   = "colorscape is a color code conversion tool."
	version = "0.0.1"
)

func main() {
	app := cli.NewApp()

	app.Name = appName
	app.Usage = usage
	app.Version = version

	app.Run(os.Args)
}
