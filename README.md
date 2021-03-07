# osinfo

![GitHub](https://img.shields.io/github/license/JustinTimperio/osinfo)
[![Go Reference](https://pkg.go.dev/badge/github.com/JustinTimperio/osinfo.svg)](https://pkg.go.dev/github.com/JustinTimperio/osinfo)
[![Go Report Card](https://goreportcard.com/badge/github.com/JustinTimperio/osinfo)](https://goreportcard.com/report/github.com/JustinTimperio/osinfo)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/30b01ce21f514d46ab3d47b5c371fa38)](https://www.codacy.com/gh/JustinTimperio/osinfo/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=JustinTimperio/osinfo&amp;utm_campaign=Badge_Grade)

## What is osinfo?
OSinfo is a cross-platform OS Version collection tool. It is designed to unify multiple OS detection methods in a single module that can easily be integrated into other projects. 

### Officially Supported:

| Windows             | MacOS              | Linux               | BSD     | Android         | IOS    |
|---------------------|--------------------|---------------------|---------|-----------------|--------|
| Windows Server 2012 | 10 - Snow Leopard  | Ubuntu              | FreeBSD | 5 - Lollipop    | IOS 10 |
| Windows Server 2016 | 11 - Lion          | Debian              | OpenBSD | 6 - Marshmallow | IOS 11 |
| Windows Server 2019 | 12 - Mountain Lion | MXLinux             |         | 7 - Nougat      | IOS 12 |
| Windows 7           | 13 - Mavericks     | Mint                |         | 8 - Oreo        | IOS 13 |
| Windows 8           | 14 - Yosemite      | Kali                |         | 9 - Pie         | IOS 14 |
| Windows 10          | 15 - El Capitan    | ParrotOS            |         |                 |        |
|                     | 16 - Sierra        | OpenSUSE Leap       |         |                 |        |
|                     | 17 - High Sierra   | OpenSUSE TumbleWeed |         |                 |        |
|                     | 18 - Mojave        | OpenSUSE SLES       |         |                 |        |
|                     | 19 - Catalina      | Arch                |         |                 |        |
|                     | 20 - Big Sur       | Manjaro             |         |                 |        |
|                     |                    | Alpine              |         |                 |        |
|                     |                    | Fedora              |         |                 |        |
|                     |                    | RHEL                |         |                 |        |
|                     |                    | CentOS              |         |                 |        |
|                     |                    | Oracle              |         |                 |        |


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
