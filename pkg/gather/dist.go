package gather

import "github.com/mr-chelyshkin/rpi4_power_controller/pkg/utils"

type distInfo struct {
	name    string
	code    string
	release string
}

func getDistInfo() (*distInfo, error) {
	b, err := utils.FileRead("/etc/os-release")
	if err != nil {
		return nil, err
	}

	res := utils.FileParseEnv(b)
	return &distInfo{
		release: res["VERSION_ID"],
		name:    res["ID"],
		code:    res["VERSION_CODENAME"],
	}, nil
}
