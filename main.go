package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

const (
	appName = "colorscape"
	usage   = "colorscape is a color code conversion tool."
	version = "0.0.1"
	author  = "ysbrothersk"
)

func main() {
	app := cli.NewApp()

	app.Name = appName
	app.Usage = usage
	app.Version = version
	app.Author = author

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "hex, x",
			Usage: "Hex color code",
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println(c.String("hex"))
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
