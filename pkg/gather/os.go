package gather

import (
	"runtime"
	"syscall"

	"github.com/mr-chelyshkin/rpi4_power_controller/pkg/utils"
)

type osInfo struct {
	os      string
	name    string
	arch    string
	host    string
	machine string
	release string
}

func getOsInfo() (*osInfo, error) {
	var utsname syscall.Utsname
	if err := syscall.Uname(&utsname); err != nil {
		return nil, err
	}

	release := utils.Int8ToString(utsname.Release[:])
	machine := utils.Int8ToString(utsname.Machine[:])
	host := utils.Int8ToString(utsname.Nodename[:])
	name := utils.Int8ToString(utsname.Sysname[:])

	return &osInfo{
		os:      runtime.GOOS,
		arch:    runtime.GOARCH,
		release: release,
		machine: machine,
		host:    host,
		name:    name,
	}, nil
}
