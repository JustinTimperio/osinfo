# osinfo

```
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
