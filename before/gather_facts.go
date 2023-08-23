package before

import (
	"context"

	"github.com/mr-chelyshkin/rpi4_power_controller"
	"github.com/mr-chelyshkin/rpi4_power_controller/pkg/gather"

	"github.com/urfave/cli/v2"
)

func GetGatherFacts(cCtx *cli.Context) error {
	facts, err := gather.New()
	if err != nil {
		return err
	}
  
	nCtx := context.WithValue(
		cCtx.Context, 
		rpi4_power_controller.ContextKeyFacts, 
		facts,
	)
	cCtx.Context = nCtx
	return nil
}

