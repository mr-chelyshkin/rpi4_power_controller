package enable

import "github.com/urfave/cli/v2"

func flags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:  "usb",
			Usage: "enable usb ports",
		},
		&cli.BoolFlag{
			Name: "hdmi",
			Usage: "enable hdmi port",
		},
	}
}
