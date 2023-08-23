package hdmi

import (
	"fmt"

	"github.com/mr-chelyshkin/rpi4_power_controller/pkg/utils"
)

type debian struct{}

func debianManager() (Manager, error) {
	return &debian{}, nil
}

func (d *debian) Enable() error {
	fmt.Println("hdmi enable")
	return nil
}

func (d *debian) Disable() error {
  b, err := utils.FileRead(bootConfigFile)
	if err != nil {
		return err
	}
	data := utils.FileParseEnv(b)
	fmt.Println(data)
	return nil
}
