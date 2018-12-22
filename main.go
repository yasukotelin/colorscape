package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ysbrothersk/colorscape/conv"

	"github.com/urfave/cli"
)

const (
	appName = "colorscape"
	usage   = "colorscape is a color code conversion tool."
	version = "1.0.0"
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
		{
			Name:    "rgb",
			Aliases: []string{"r"},
			Usage:   "convert to hex color code from rgb",
			Action:  rgbAction,
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

func rgbAction(c *cli.Context) error {
	rgb := c.Args()
	r, err := strconv.Atoi(rgb[0])
	if err != nil {
		return err
	}
	g, err := strconv.Atoi(rgb[1])
	if err != nil {
		return err
	}
	b, err := strconv.Atoi(rgb[2])
	if err != nil {
		return err
	}

	colorCode, err := conv.ToColorCode(uint8(r), uint8(g), uint8(b))
	if err != nil {
		return err
	}

	fmt.Println(colorCode)

	return nil
}
