package osinfo

import (
	"fmt"
)

type Release struct {
	Runtime string
	Arch    string
	Name    string
	Version string
	win     WindowsRelease
	nix     LinuxRelease
	bsd     BSDRelease
	osx     DarwinRelease
}

type WindowsRelease struct {
	Build string
}

type LinuxRelease struct {
	Distro string
	Kernel string
	PkgMng string
}

type BSDRelease struct {
	Kernel string
	PkgMng string
}

type DarwinRelease struct {
	Kernel string
}

func PrintInfo(i *Release) {

	fmt.Println("Runtime:", i.Runtime)
	fmt.Println("Architecture:", i.Arch)
	fmt.Println("OS Name:", i.Name)
	fmt.Println("Version:", i.Version)

	if i.Runtime == "freebsd" {
		fmt.Println("Kernel:", i.bsd.Kernel)

	} else if i.Runtime == "linux" {
		fmt.Println("Kernel:", i.nix.Kernel)
		fmt.Println("Distro:", i.nix.Distro)
		fmt.Println("Package Manager:", i.nix.PkgMng)

	} else if i.Runtime == "darwin" {
		fmt.Println("Kernel:", i.osx.Kernel)

	} else if i.Runtime == "windows" {
		fmt.Println("Build Number:", i.win.Build)
	}
}
