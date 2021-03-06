package osinfo

import (
	"fmt"

	"github.com/JusinTimperio/osinfo"
)

func ExampleStruct() {
	release := osinfo.GetVersion()
	fmt.Println(release.Name)
	fmt.Println(release.Version)
}

func ExamplePrint() {
	release := osinfo.GetVersion()
	release.PrintInfo()
}
