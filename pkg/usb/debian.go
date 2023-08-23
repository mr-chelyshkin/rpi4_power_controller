package usb

import (
	"fmt"

	"github.com/mr-chelyshkin/rpi4_power_controller/pkg/utils"
)

type debian struct{}

func debianManager() (Manager, error) {
	return &debian{}, nil
}

func (d *debian) Enable() error {
	err := utils.FileWrite(driverPathBind, []byte("1-1"))
	if err != nil && err.Error() == "device or resource busy" {
		return fmt.Errorf("ports already enabled")
	}
	return err
}

func (d *debian) Disable() error {
	err := utils.FileWrite(driverPathUnbind, []byte("1-1"))
	if err != nil && err.Error() == "no such device" {
		return fmt.Errorf("ports already disabled")
	}
	return err
}
