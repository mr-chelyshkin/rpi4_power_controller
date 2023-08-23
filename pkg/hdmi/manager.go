package hdmi

import (
	"errors"
	"fmt"

	"github.com/mr-chelyshkin/rpi4_power_controller/pkg/gather"
)

const (
	bootConfigFile = "/boot/config.txt"
)

type Manager interface {
	Disable() error
	Enable() error
}

func New(facts *gather.Facts) (Manager, error) {
	switch facts.DistrInfoName() {
	case "debian":
	return debianManager()
	default:
		msg := fmt.Sprintf(
			"usb manager not implemented for distr: %s", 
			facts.DistrInfoName(),
		)
		return nil, errors.New(msg)
	}
}
