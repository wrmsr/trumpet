package main

import (
	"runtime"

	"github.com/wrmsr/trumpet"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	trumpet.Run()
}
