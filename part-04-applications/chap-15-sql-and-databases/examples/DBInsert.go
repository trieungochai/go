package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=911225 host=localhost port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connection to the DB was successfully initialized!")
	}

	insert, err := db.Prepare("INSERT INTO test VALUES ($1, $2)")
	if err != nil {
		panic(err)
	}

	_, err = insert.Exec(2, "second")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The table was successfully inserted!")
	}

	db.Close()
}

// db.Prepare() takes a SQL statement and imbues it with protection against SQL injection attacks.
// It works by restricting the values of the variable substitutions.
// In our case, we have two columns, so for the substitution to work, we use $1 and $2.
// You can use any number of substitutions; you only need to make sure they result in a valid SQL statement when evaluated.
// When the insert variable is initialized without errors, it will be responsible for executing the SQL statement.
// It finds out how many arguments the prepared statement expects,
// and its sole purpose is to call the statement and perform the operation.
// insert.Exec(2,"second") inserts a new element with id=2 and name='second'.
// If we were to check what we have in our database, we would see the results.
