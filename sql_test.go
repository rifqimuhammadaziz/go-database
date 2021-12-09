package godatabase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id, name) VALUES('2', 'Rifqi')"
	_, err := db.ExecContext(ctx, query) // INSERT DATA
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert New Customer!")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, query) // READ DATA
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// iterate rows until next is false
	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		// output data
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}
}

func TestQuerySqlMultipleData(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"

	rows, err := db.QueryContext(ctx, query) // used to read data
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// iterate rows until next is false
	for rows.Next() {
		// mapping data
		var id, name string        // varchar
		var email sql.NullString   // nullable varchar
		var balance int32          // int(int32), bigint(int64)
		var rating float64         // double
		var birthDate sql.NullTime // nullable date
		var createdAt time.Time    // date, timestamp
		var married bool           // boolean

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		// output data
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)

		// check data if valid (nullable field table)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}

		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)

		// check data if valid (nullable field table)
		if birthDate.Valid {
			fmt.Println("Birth Date:", birthDate.Time)
		}

		fmt.Println("Married:", married)
		fmt.Println("Created At:", createdAt)
		fmt.Println("===========================")
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// testing using username password = admin
	username := "admin'; #" // sql injection, next code after input username become comments (#)
	password := "password"

	query := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	fmt.Println(query)

	rows, err := db.QueryContext(ctx, query) // READ DATA
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// iterate rows until next is false
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success login:", username)
	} else {
		fmt.Println("Failed login, username", username, "wrong password / not found.")
	}
}
