package osinfo

import (
	"os/exec"
	"runtime"
	"strings"
)

// GetVersion FreeBSD returns version info
func GetVersion() *Release {
	// fetching info for this os is fairly simple
	// version information is all fetched via `uname`
	// Returns:
	//		- r.Runtime
	//		- r.Arch
	//		- r.Name
	//		- r.Version
	//		- r.bsd.Kernel

	inf := &Release{
		Runtime: runtime.GOOS,
		Arch:    runtime.GOARCH,
	}

	out0, _ := exec.Command("uname", "-or").Output()
	inf.Name = strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(string(out0), "\n", ""), "\"", ""))

	out1, _ := exec.Command("uname", "-r").Output()
	inf.Version = strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(string(out1), "\n", ""), "\"", ""))

	out2, _ := exec.Command("uname", "-K").Output()
	inf.bsd.Kernel = strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(string(out2), "\n", ""), "\"", ""))

	inf.bsd.PkgMng = "pkg"

	return inf
}
