package osinfo

import "fmt"

func fetchAndPrint() {
	x := GetVersion()
	fmt.Println("-------------")
	fmt.Println(x.Arch)
	fmt.Println(x.Runtime)
	fmt.Println(x.Name)
	fmt.Println(x.Version)
	fmt.Println("-------------")
	x.PrintInfo()
}
