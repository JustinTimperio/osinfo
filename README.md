# osinfo

## What is osinfo?
OSinfo is a cross-platform OS Version collection tool. 

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
```

## Example Usage
```go
   package main

   import (
	   "github.com/JustinTimperio/osinfo"
   )

   func main() {
		Release := osinfo.GetVersion()
		Release.PrintInfo()
	 }
}
```

## Example Output
```
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
Kernel: 5.11.2-arch1-1
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
```
