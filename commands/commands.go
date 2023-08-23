package commands

import (
	"github.com/mr-chelyshkin/rpi4_power_controller/commands/disable"
	"github.com/mr-chelyshkin/rpi4_power_controller/commands/enable"
	"github.com/mr-chelyshkin/rpi4_power_controller/commands/facts"

	"github.com/urfave/cli/v2"
)

func Commands() []*cli.Command {
	return []*cli.Command{
		facts.Command(),
		enable.Command(),
		disable.Command(),
	}
}
