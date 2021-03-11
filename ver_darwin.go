package osinfo

import (
	"os/exec"
	"runtime"
	"strings"
)

// GetVersion Darwin returns version info
// fetching info for this os is fairly simple
// version information is all fetched via `sw_vers`
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

	version, err := exec.Command("sw_vers").Output()
	if err == nil {
		str := strings.Split(string(version), "\n")
		for _, s := range str {
			if strings.HasPrefix(s, "ProductVersion:\t") {
				info.Version = strings.TrimPrefix(s, "ProductVersion:\t")
			}
		}
	}

	switch idx := strings.LastIndex(info.Version, "."); info.Version[0:idx] {
	case "10.6":
		info.Name = "MacOS: Snow Leopard"
	case "10.7":
		info.Name = "MacOS: Lion"
	case "10.8":
		info.Name = "MacOS: Mountain Lion"
	case "10.9":
		info.Name = "MacOS: Mavericks"
	case "10.10":
		info.Name = "MacOS: Yosemite"
	case "10.11":
		info.Name = "MacOS: El Capitan"
	case "10.12":
		info.Name = "MacOS: Sierra"
	case "10.13":
		info.Name = "MacOS: High Sierra"
	case "10.14":
		info.Name = "MacOS: Mojave"
	case "10.15":
		info.Name = "MacOS: Catalina"
	case "11.0":
		info.Name = "MacOS: Big Sur"
	default:
		info.Name = "MacOS: Unknown Version"
	}
	return info
}
