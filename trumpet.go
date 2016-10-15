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

	cs := C.CString("barf")
	defer C.free(unsafe.Pointer(cs))
	db := C.PQconnectdb(cs)

	if C.PQstatus(db) != C.CONNECTION_OK {
		fmt.Print(C.GoString(C.PQerrorMessage(db)))
	}

	C.PQfinish(db)
}
