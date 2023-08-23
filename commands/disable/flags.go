package disable

import "github.com/urfave/cli/v2"

func flags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name: "usb",
			Usage: "disable usb ports",
		},
		&cli.BoolFlag{
			Name: "hdmi",
			Usage: "disable hdmi port",
		},
	}
}
