package trumpet

import (
	"database/sql"
	"fmt"
	"testing"
	_ "./postgres"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/golang/go/src/io/ioutil"
	"os"
	"path"
)

func TestRun(t *testing.T) {
	Run()
}

func chkerr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
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
		for n := range cols {
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

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func TestGorm(t *testing.T) {
	dbdir, err := ioutil.TempDir("", "trumpet-test")
	if err != nil {
		panic("failed to create temp dir")
	}
	defer os.Remove(dbdir)

	db, err := gorm.Open("sqlite3", path.Join(dbdir, "test.db"))
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	db.First(&product, 1) // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	var code string
	var price uint
	row := db.Table("products").Where("code = ?", "L1212").Select("code, price").Row() // (*sql.Row)
	row.Scan(&code, &price)
	fmt.Println(code)
	fmt.Println(price)

	rows, err := db.Model(&Product{}).Where("code = ?", "L1212").Select("code, price").Rows() // (*sql.Rows, error)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&code, &price)
		fmt.Println(code)
		fmt.Println(price)
	}

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)
}
