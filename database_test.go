package godatabase

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql" // run init function of package
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/db_godatabase")
	if err != nil {
		panic(err)
	}
	defer db.Close() // ensure database closed if apps done using db
}
