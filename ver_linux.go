package osinfo

import (
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

// GetVersion Linux returns version info
func GetVersion() *Release {
	// fetching os info for linux distros is more complicated process than it should be
	// version information is collected via `uname` and `/etc/os-release`
	// Returns:
	//		- r.Runtime
	//		- r.Arch
	//		- r.Name
	//		- r.Version
	//		- r.nix.Kernel
	//		- r.nix.Distro
	//		- r.nix.PkgMng

	inf := &Release{
		Runtime: runtime.GOOS,
		Arch:    runtime.GOARCH,
	}

	out, _ := exec.Command("uname", "-r").Output()
	kern := strings.ReplaceAll(strings.ReplaceAll(string(out), "\n", ""), "\"", "")
	inf.nix.Kernel = string(kern)

	if pathExists("/etc/os-release") == true {
		f := readFile("/etc/os-release")

		var (
			nameField  = regexp.MustCompile(`NAME=(.*?)\n|\nNAME=(.*?)\n`)
			pnameField = regexp.MustCompile(`PRETTY_NAME=(.*?)\n|\nPRETTY_NAME=(.*?)\n`)
			verField   = regexp.MustCompile(`VERSION_ID=(.*?)\n|\nVERSION_ID=(.*?)\n`)

			suse   = regexp.MustCompile(`SLES|openSUSE`)
			debian = regexp.MustCompile(`Debian|Ubuntu|Kali|Parrot|Mint`)
			rhl    = regexp.MustCompile(`Red Hat|CentOS|Fedora|Oracle`)
			arch   = regexp.MustCompile(`Arch|Manjaro`)
			alpine = regexp.MustCompile(`Alpine`)
		)

		namef := strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(nameField.FindString(f), "\n", ""), "\"", ""))
		pnamef := strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(pnameField.FindString(f), "\n", ""), "\"", ""))
		verf := strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(verField.FindString(f), "\n", ""), "\"", ""))

		name := strings.Split(namef, "=")[1]

		if pnamef != "" {
			inf.Name = strings.Split(pnamef, "=")[1]
		} else {
			inf.Name = "unknown"
		}

		if verf != "" {
			inf.Version = strings.Split(verf, "=")[1]
		} else {
			inf.Version = "unknown"
		}

		if suse.MatchString(name) == true {
			inf.nix.Distro = "opensuse"
			inf.nix.PkgMng = "zypper"

		} else if debian.MatchString(name) == true {
			inf.nix.Distro = "debian"
			inf.nix.PkgMng = "apt"

		} else if rhl.MatchString(name) == true {
			inf.nix.Distro = "fedora"
			inf.nix.PkgMng = "yum"

		} else if arch.MatchString(name) == true {
			inf.Version = "rolling"
			inf.nix.Distro = "arch"
			inf.nix.PkgMng = "pacman"

		} else if alpine.MatchString(name) == true {
			inf.nix.Distro = "alpine"
			inf.nix.PkgMng = "apk"

		} else {
			inf.nix.Distro = "unknown"
			inf.nix.PkgMng = "unknown"
		}

	} else {
		inf.Name = "unknown"
		inf.Version = "unknown"
		inf.nix.Distro = "unknown"
		inf.nix.PkgMng = "unknown"
	}
	return inf
}

func pathExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func readFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(content[:])
}
