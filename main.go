package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("number of cores: ", runtime.NumCPU())
}
