package trumpet

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	Run()
}

func TestInformationSchema(t *testing.T) {
	db, err := sql.Open("postgres", "host=192.168.99.100 port=9109 user=postgres")
	chkerr(t, err)
	defer db.Close()

	chkerr(t, db.Ping())

	rows, err := db.Query("select * from information_schema.tables")
	chkerr(t, err)

	cols, err := rows.Columns()
	fmt.Println(cols)
	count := len(cols)
	vals := make([]interface{}, count)
	valPtrs := make([]interface{}, count)

	for i := 0; rows.Next(); i++ {
		for n, _ := range cols {
			valPtrs[n] = &vals[n]
		}
		rows.Scan(valPtrs...)
		for n, col := range cols {
			var v interface{}
			val := vals[n]
			b, ok := val.([]byte)
			if (ok) {
				v = string(b)
			} else {
				v = val
			}
			fmt.Println(col, v)
		}
		fmt.Println()
	}
}
