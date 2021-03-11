package osinfo_test

import (
	"fmt"
	"testing"

	"github.com/JustinTimperio/osinfo"
)

func TestMain(m *testing.M) {
	info := osinfo.GetVersion()
	fmt.Println("-------------")
	fmt.Println(info.Arch)
	fmt.Println(info.Runtime)
	fmt.Println(info.Name)
	fmt.Println(info.Version)
	fmt.Println("-------------")

	fmt.Printf(info.String())
	fmt.Println("-------------")
}
