package osinfo_test

import (
	"fmt"
	"testing"

	"github.com/JustinTimperio/osinfo"
)


func TestMain(m *testing.M) {
	x := osinfo.GetVersion()
	fmt.Println("-------------")
	fmt.Println(x.Arch)
	fmt.Println(x.Runtime)
	fmt.Println(x.Name)
	fmt.Println(x.Version)
	fmt.Println("-------------")
	x.PrintInfo()
}
