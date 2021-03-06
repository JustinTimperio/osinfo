package osinfo

import (
	"runtime"

	"golang.org/x/sys/windows/registry"
)

// GetVersion Windows returns version info
func GetVersion() *Release {
	// fetching os info for modern windows versions is fairly simple
	// version info is easily fetched in the registry
	// Returns:
	//		- r.Runtime
	//		- r.Arch
	//		- r.Name
	//		- r.Version
	//		- r.win.Build

	inf := &Release{
		Runtime: runtime.GOOS,
		Arch:    runtime.GOARCH,
	}

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		inf.Name = "unknown"
		inf.Version = "unknown"
		inf.win.Build = "unknown"
		return inf
	}
	defer k.Close()

	// extract info
	pname, _, err = k.GetStringValue("ProductName")
	if err != nil {
		inf.Name = "unknown"
	} else {
		inf.Name = pname
	}

	ver, _, err = k.GetIntegerValue("CurrentMajorVersionNumber")
	if err != nil {
		inf.Version = "unknown"
	} else {
		inf.Version = ver
	}

	build, _, err = k.GetStringValue("CurrentBuild")
	if err != nil {
		inf.win.Build = "unknown"
	} else {
		inf.win.Build = build
	}

	return inf
}
