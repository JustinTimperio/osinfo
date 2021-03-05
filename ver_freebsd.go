package main

import (
	"os/exec"
	"runtime"
	"strings"
)

func GetVersion() *Release {

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
