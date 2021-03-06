package osinfo

import (
	"os/exec"
	"runtime"
	"strings"
)

// GetVersion FreeBSD returns version info
// fetching info for this os is fairly simple
// version information is all fetched via `uname`
// Returns:
//		- r.Runtime
//		- r.Arch
//		- r.Name
//		- r.Version
//		- r.bsd.Kernel
func GetVersion() *Release {

	inf := &Release{
		Runtime: runtime.GOOS,
		Arch:    runtime.GOARCH,
	}

	fullName, _ := exec.Command("uname", "-or").Output()
	inf.Name = strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(string(fullName), "\n", ""), "\"", ""))

	version, _ := exec.Command("uname", "-r").Output()
	inf.Version = strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(string(version), "\n", ""), "\"", ""))

	kernel, _ := exec.Command("uname", "-K").Output()
	inf.bsd.Kernel = strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(string(kernel), "\n", ""), "\"", ""))

	inf.bsd.PkgMng = "pkg"

	return inf
}
