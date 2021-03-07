package osinfo

import (
	"os/exec"
	"runtime"
	"strings"
)

// GetVersion Darwin returns version info
// fetching info for this os is fairly simple
// version information is all fetched via `uname`
// Returns:
//		- r.Runtime
//		- r.Arch
//		- r.Name
//		- r.Version
//		- r.bsd.Kernel
func GetVersion() Release {
	info := Release{
		Runtime: runtime.GOOS,
		Arch:    runtime.GOARCH,
		Name:    "unknown",
		Version: "unknown",
	}

	version, err := exec.Command("uname", "-r").Output()
	if err == nil {
		info.Version = cleanString(string(version))
	}

	if strings.HasPrefix(version, "10") {
		info.Name = "MacOS: Snow Leopard"

	} else if strings.HasPrefix(version, "11") {
		info.Name = "MacOS: Lion"

	} else if strings.HasPrefix(version, "12") {
		info.Name = "MacOS: Mountain Lion"

	} else if strings.HasPrefix(version, "13") {
		info.Name = "MacOS: Mavericks"

	} else if strings.HasPrefix(version, "14") {
		info.Name = "MacOS: Yosemite"

	} else if strings.HasPrefix(version, "15") {
		info.Name = "MacOS: El Capitan"

	} else if strings.HasPrefix(version, "16") {
		info.Name = "MacOS: Sierra"

	} else if strings.HasPrefix(version, "17") {
		info.Name = "MacOS: High Sierra"

	} else if strings.HasPrefix(version, "18") {
		info.Name = "MacOS: Mojave"

	} else if strings.HasPrefix(version, "19") {
		info.Name = "MacOS: Catalina"

	} else if strings.HasPrefix(version, "20") {
		info.Name = "MacOS: Big Sur"

	} else {
		info.Name = "MacOS: Unknown Version"
	}

	return info
}
