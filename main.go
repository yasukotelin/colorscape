package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ysbrothersk/colorscape/conv"

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

	app.Commands = []cli.Command{
		{
			Name:    "code",
			Aliases: []string{"c", "hex", "h"},
			Usage:   "convert to rgb from hex color code",
			Action:  codeAction,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func codeAction(c *cli.Context) error {
	code := c.Args().First()
	rgb, err := conv.ToRgb(code)
	if err != nil {
		return err
	}

	fmt.Println(rgb[0], rgb[1], rgb[2])

	return nil
}
