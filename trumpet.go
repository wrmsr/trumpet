package main

/*
#cgo pkg-config: libpq
#include <stdlib.h>
#include <libpq-fe.h>
*/
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	C.PQClear(unsafe.Pointer{0})

	fmt.Print("hi")
}
