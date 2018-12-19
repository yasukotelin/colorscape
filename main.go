package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/ysbrothersk/colorscape/conv"
)

const (
	appName = "colorscape"
	usage   = "colorscape is a color code conversion tool."
	version = "0.0.1"
	author  = "ysbrothersk"
)

const (
	codeFlag1 = "code"
	codeFlag2 = "c"
	rgbFlag1  = "rgb"
	rgbFlag2  = "r"
)

func main() {
	app := cli.NewApp()

	app.Name = appName
	app.Usage = usage
	app.Version = version
	app.Author = author

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  fmt.Sprintf("%s, %s", codeFlag1, codeFlag2),
			Usage: "Convert hex color code to RGB.",
		},
		cli.StringFlag{
			Name:  fmt.Sprintf("%s, %s", rgbFlag1, rgbFlag2),
			Usage: "Convert RGB to hex color code.",
		},
	}

	app.Action = func(c *cli.Context) error {
		hexArg := c.String(codeFlag1)

		if len(hexArg) > 0 {
			rgb, err := conv.ToRgb(hexArg)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(rgb[0], rgb[1], rgb[2])
		}

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
