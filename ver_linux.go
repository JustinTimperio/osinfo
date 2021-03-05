package osinfo

import (
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

func GetVersion() *Release {

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
			name_field  = regexp.MustCompile(`NAME=(.*?)\n|\nNAME=(.*?)\n`)
			pname_field = regexp.MustCompile(`PRETTY_NAME=(.*?)\n|\nPRETTY_NAME=(.*?)\n`)
			ver_field   = regexp.MustCompile(`VERSION_ID=(.*?)\n|\nVERSION_ID=(.*?)\n`)

			suse   = regexp.MustCompile(`SLES|openSUSE`)
			debian = regexp.MustCompile(`Debian|Ubuntu|Kali|Parrot`)
			rhl    = regexp.MustCompile(`Red Hat|CentOS|Fedora|Oracle`)
			arch   = regexp.MustCompile(`Arch|Manjaro`)
			alpine = regexp.MustCompile(`Alpine`)
		)

		name_f := strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(name_field.FindString(f), "\n", ""), "\"", ""))
		pname_f := strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(pname_field.FindString(f), "\n", ""), "\"", ""))
		ver_f := strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(ver_field.FindString(f), "\n", ""), "\"", ""))

		name := strings.Split(name_f, "=")[1]

		if pname_f != "" {
			inf.Name = strings.Split(pname_f, "=")[1]
		} else {
			inf.Name = "unknown"
		}

		if ver_f != "" {
			inf.Version = strings.Split(ver_f, "=")[1]
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
