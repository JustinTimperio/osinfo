package osinfo

import (
	"fmt"
)

// Release contains all the
type Release struct {
	Runtime string
	Arch    string
	Name    string
	Version string
	win     windowsRelease
	nix     linuxRelease
	bsd     bsdRelease
	osx     darwinRelease
}

type windowsRelease struct {
	Build string
}

type linuxRelease struct {
	Distro string
	Kernel string
	PkgMng string
}

type bsdRelease struct {
	Kernel string
	PkgMng string
}

type darwinRelease struct {
	Kernel string
}

// PrintInfo prints all the fields in a release struct
func (i *Release) PrintInfo() {

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
