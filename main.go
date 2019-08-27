package main

import (
	"fmt"
	"os"
	"runtime/debug"
)

func main() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		panic("couldn't read build info")
	}

	fmt.Println(bi.Path)

	fmt.Println(bi.Main.Path)
	fmt.Println(bi.Main.Version)
	fmt.Println(bi.Main.Sum)
	fmt.Println()

	for _, mod := range bi.Deps {
		fmt.Println(mod.Path)
		fmt.Println(mod.Version)
		fmt.Println(mod.Sum)
		fmt.Println()
	}

	os.Exit(0)
}
