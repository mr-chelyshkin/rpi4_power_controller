package disable

import (
	"errors"
	"fmt"

	"github.com/mr-chelyshkin/rpi4_power_controller"
	"github.com/mr-chelyshkin/rpi4_power_controller/before"
	"github.com/mr-chelyshkin/rpi4_power_controller/pkg/gather"
	"github.com/mr-chelyshkin/rpi4_power_controller/pkg/hdmi"
	"github.com/mr-chelyshkin/rpi4_power_controller/pkg/usb"
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:        "disable",
		Usage:       "disable devices",
		Description: "disable devices",
		Action:      action,
		Flags:       flags(),

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

	switch {
	case cCtx.Bool("usb"):
		fmt.Println("disable usb ports ...")

		manager, err := usb.New(facts)
		if err != nil {
			return err
		}
		if err := manager.Disable(); err != nil {
			return err
		}

		fmt.Println("usb ports disabled")
		return nil
	case cCtx.Bool("hdmi"):
		fmt.Println("disable hdmi port ...")

		manager, err := hdmi.New(facts)
		if err != nil {
			return err
		}
		if err := manager.Disable(); err != nil {
			return err
		}

		fmt.Println("hdmi port disabled")
		return nil
	}

	return fmt.Errorf("flags not found, see '--help'")
}
