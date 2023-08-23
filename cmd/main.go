package main

import (
	"log"
	"os"

	"github.com/mr-chelyshkin/rpi4_power_controller"
	"github.com/mr-chelyshkin/rpi4_power_controller/commands"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
    Name:  rpi4_power_controller.Name,
    Usage: rpi4_power_controller.Usage,
    Commands: commands.Commands(),
  }

  if err := app.Run(os.Args); err != nil {
      log.Fatal(err)
  }
}
