package gather

type Facts struct {
	*osInfo
	*distInfo
}

func New() (*Facts, error) {
	osInfo, err := getOsInfo()
	if err != nil {
		return nil, err
	}
	distInfo, err := getDistInfo()
	if err != nil {
		return nil, err
	}

	return &Facts{
		osInfo:     osInfo,
		distInfo:   distInfo,
	}, nil
}

func (f *Facts) OsInfoOS() string {
	return f.osInfo.os
}

func (f *Facts) OsInfoName() string {
	return f.osInfo.name
}

func (f *Facts) OsInfoArch() string {
	return f.osInfo.arch
}

func (f *Facts) OsInfoHost() string {
	return f.osInfo.host
}

func (f *Facts) OsInfoMachine() string {
	return f.osInfo.machine
}

func (f *Facts) OsInfoRelease() string {
	return f.osInfo.release
}

func (f *Facts) DistrInfoRelease() string {
	return f.distInfo.release
}

func (f *Facts) DistrInfoName() string {
	return f.distInfo.name
}

func (f *Facts) DistInfoCodeName() string {
	return f.distInfo.code
}

