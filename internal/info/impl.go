package info

import "strings"

type appInfoImpl struct {
	AppName      string
	AppVersion   string
	AppBuildTime string
}

func getAppInfoInternal() AppInfo {
	return appInfoImpl{
		AppName:      strings.ToUpper(Name),
		AppVersion:   Version,
		AppBuildTime: BuildTime,
	}
}

func (impl appInfoImpl) Name() string {
	return impl.AppName
}

func (impl appInfoImpl) Version() string {
	return impl.AppVersion
}

func (impl appInfoImpl) BuildTime() string {
	return impl.AppBuildTime
}

func (impl appInfoImpl) Print() {
	printInfo(impl)
}
