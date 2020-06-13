package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("os: ", runtime.GOOS)
	fmt.Println("number of cores: ", runtime.NumCPU())
}
