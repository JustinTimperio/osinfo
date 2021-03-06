# osinfo

[![Go Reference](https://pkg.go.dev/badge/github.com/JustinTimperio/osinfo.svg)](https://pkg.go.dev/github.com/JustinTimperio/osinfo)
[![Go Report Card](https://goreportcard.com/badge/github.com/JustinTimperio/osinfo)](https://goreportcard.com/report/github.com/JustinTimperio/osinfo)
![GitHub](https://img.shields.io/github/license/JustinTimperio/osinfo)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/30b01ce21f514d46ab3d47b5c371fa38)](https://www.codacy.com/gh/JustinTimperio/osinfo/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=JustinTimperio/osinfo&amp;utm_campaign=Badge_Grade)

## What is osinfo?
OSinfo is a cross-platform OS Version collection tool. It is designed to unify multiple OS detection methods in a single module that can easily be integrated into other projects. 

### Offically Supported

Windows Based: 
  - Windows Server 2016
  - Windows Server 2019
  - Windows 7
  - Windows 8
  - Windows 10

Debian Linux Based:
  - Ubuntu
  - Debian
  - Kali Rolling
  - Parrot OS

Red Hat Linux Based:
  - Fedora
  - CentOS
  - RHEL
  - Oracle

OpenSUSE Linux Based:
  - Leap
  - TumbleWeed
  - SLES

Arch Linux Based:
  - Arch
  - Manjaro

Alpine Linux Based:
 - Alpine

BSD Based:
  - FreeBSD

## Release Struct
```go
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
```

## Example Usage
 1. Create`fetchinfo.go`
```go
   package main

   import (
	   "github.com/JustinTimperio/osinfo"
   )

   func main() {
		release := osinfo.GetVersion()
		release.PrintInfo()
	 }
```
 2. `go mod init`
 3. `go mod tidy`
 4. `go run fetchinfo.go`

## Example Outputs
```sh
--------------------

Runtime: linux
Architecture: amd64
OS Name: Arch Linux
Version: rolling
Kernel: 5.11.2-arch1-1
Distro: arch
Package Manager: pacman

--------------------

Runtime: linux
Architecture: amd64
OS Name: Debian GNU/Linux 10 (buster)
Version: 10
Kernel: 4.19.0-13-amd64
Distro: debian
Package Manager: apt

--------------------

Runtime: windows
Architecture: amd64
OS Name: Windows Server 2016 Standard Evaluation
Version: 10
Build: 14393

--------------------

Runtime: freebsd
Architecture: amd64
OS Name: FreeBSD 12.1-RELEASE
Version: 12.1-RELEASE
Kernel: 1201000

--------------------
```
