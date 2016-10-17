package trumpet

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
	"unsafe"
)

func Run() {
	cs := C.CString("host=192.168.99.100 port=9109 user=postgres")
	defer C.free(unsafe.Pointer(cs))
	db := C.PQconnectdb(cs)

	if C.PQstatus(db) != C.CONNECTION_OK {
		fmt.Print(C.GoString(C.PQerrorMessage(db)))
	}

	C.PQfinish(db)

	m := map[string]string{
		"host":                      "192.168.99.100",
		"port":                      "9109",
		"user":                      "postgres",
		"replication":               "database",
		"fallback_application_name": "trumpet",
	}

	keys := make([]unsafe.Pointer, len(m)+1)
	values := make([]unsafe.Pointer, len(m)+1)
	i := 0
	for k, v := range m {
		keys[i] = unsafe.Pointer(C.CString(k))
		defer func(i int) {
			C.free(keys[i])
		}(i)
		values[i] = unsafe.Pointer(C.CString(v))
		defer func(i int) {
			C.free(values[i])
		}(i)
		i += 1
	}
	keys[i] = nil
	values[i] = nil

	db = C.PQconnectdbParams(unsafe.Pointer(&keys[0]), unsafe.Pointer(&values[0]), 1)
	if C.PQstatus(db) != C.CONNECTION_OK {
		fmt.Print(C.GoString(C.PQerrorMessage(db)))
	}
	C.PQfinish(db)
}
