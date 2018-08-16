package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("cpus:", runtime.NumCPU())
	fmt.Println("goroot:", runtime.GOROOT())
	fmt.Println("os/platform:", runtime.GOOS)
}
