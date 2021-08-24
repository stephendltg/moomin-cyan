package main

import (
	"fmt"
	"runtime"
)

func main() {

	if runtime.GOOS == "darwin" {
		fmt.Println("mac")
	} else if runtime.GOOS == "window" {
		fmt.Println("window")
	} else if runtime.GOOS == "linux" {
		fmt.Println("linux")
	}
}
