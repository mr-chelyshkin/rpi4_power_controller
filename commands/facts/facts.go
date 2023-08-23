package facts

import (
	"errors"

	"github.com/mr-chelyshkin/rpi4_power_controller"
	"github.com/mr-chelyshkin/rpi4_power_controller/before"
	"github.com/mr-chelyshkin/rpi4_power_controller/pkg/gather"
	"github.com/mr-chelyshkin/rpi4_power_controller/pkg/pretty"

	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:        "facts",
		Usage:       "show system info",
		Description: "show system info",
		Action:      action,

		Before: func(ctx *cli.Context) error {
			return before.GetGatherFacts(ctx)
		},
	}
}

func action(cCtx *cli.Context) error {
	facts, ok := cCtx.Context.Value(rpi4_power_controller.ContextKeyFacts).(*gather.Facts)
	if !ok {
		return errors.New("Gather facts data not found")
	}

	tHead := []string{"Key", "Value"}
	tBody := [][]string{
		{"OS", facts.OsInfoOS()},
		{"Name", facts.OsInfoName()},
		{"Arch", facts.OsInfoArch()},
		{"Host", facts.OsInfoHost()},
		{"Machine", facts.OsInfoMachine()},
		{"Release", facts.OsInfoRelease()},

		{"Distr name", facts.DistrInfoName()},
		{"Distr release", facts.DistrInfoRelease()},
		{"Distr codename", facts.DistInfoCodeName()},
	}
	pretty.RenderTable(tHead, tBody)
	return nil
}
