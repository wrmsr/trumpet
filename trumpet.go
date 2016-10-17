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
	"log"
	"unsafe"
)

func Run() error {
	cs := C.CString("host=192.168.99.100 port=9109 user=postgres")
	defer C.free(unsafe.Pointer(cs))
	db := C.PQconnectdb(cs)
	defer C.PQfinish(db)

	if C.PQstatus(db) != C.CONNECTION_OK {
		err := fmt.Errorf("Error connecting db: %s", C.GoString(C.PQerrorMessage(db)))
		log.Print(err)
		return err
	}

	m := map[string]string{
		"host":                      "192.168.99.100",
		"port":                      "9109",
		"user":                      "postgres",
		"replication":               "database",
		"fallback_application_name": "trumpet",
	}

	keys := make([]unsafe.Pointer, len(m) + 1)
	values := make([]unsafe.Pointer, len(m) + 1)
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

	repl := C.PQconnectdbParams(unsafe.Pointer(&keys[0]), unsafe.Pointer(&values[0]), 1)
	defer C.PQfinish(repl)
	if C.PQstatus(repl) != C.CONNECTION_OK {
		err := fmt.Errorf("Error connecting repl: %s", C.GoString(C.PQerrorMessage(repl)))
		log.Print(err)
		return err
	}

	identify_system := C.CString("IDENTIFY_SYSTEM")
	defer C.free(unsafe.Pointer(identify_system))
	res := C.PQexec(repl, identify_system)
	defer C.PQclear(res)
	if C.PQresultStatus(res) != C.PGRES_TUPLES_OK {
		err := fmt.Errorf("Error calling identify_system %s", C.GoString(C.PQerrorMessage(repl)))
		log.Print(err)
		return err
	}

	if C.PQntuples(res) != 1 || C.PQnfields(res) < 4 {
		err := fmt.Errorf("Unexpected identify_system result (%d rows, %d fields).", C.PQntuples(res), C.PQnfields(res))
		log.Print(err)
		return err
	}

	/* Check that the database name (fourth column of the result tuple) is non-null,
	 * implying a database-specific connection. */
	if C.PQgetisnull(res, 0, 3) != 0 {
		err := fmt.Errorf("%s", "Not using a database-specific replication connection.")
		log.Print(err)
		return err
	}

	for i := 0; i < 3; i += 1 {
		val := C.GoString(C.PQgetvalue(res, C.int(0), C.int(i)))
		fmt.Println(val)
	}

	return nil
}
