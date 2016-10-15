package main

// https://github.com/jbarham/gopgsqldriver
// https://github.com/kardianos/govendor
// https://github.com/Shopify/sarama
// https://github.com/aws/aws-sdk-go

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

	cs := C.CString("host=192.168.99.100 port=9109 user=postgres")
	defer C.free(unsafe.Pointer(cs))
	db := C.PQconnectdb(cs)

	if C.PQstatus(db) != C.CONNECTION_OK {
		fmt.Print(C.GoString(C.PQerrorMessage(db)))
	}

	C.PQfinish(db)
}
