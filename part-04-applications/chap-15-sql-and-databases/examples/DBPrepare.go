// prepare a query that is dynamic in the sense that it accepts a parameter that will be the ID we are looking for.
// Then, qryrow is used to execute the QueryRow() function,
// which, in turn, takes the id variable we specified previously and returns the result in the name variable.
// Then, we output the string with an explanation that the value of the column is based on the id variable that was specified.
// In the end, the qryrow and db resources are closed.

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	var id int
	var name string

	id = 2

	db, err := sql.Open("postgres", "user=postgres password=911225 host=localhost port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connection to the DB was successfully initialized!")
	}

	qryrow, err := db.Prepare("SELECT name from test WHERE id=$1")
	if err != nil {
		panic(err)
	}
	err = qryrow.QueryRow(id).Scan(&name)
	if err != nil {
		panic(err)
	}

	fmt.Printf("The name with id %d is %s", id, name)
	err = qryrow.Close()
	if err != nil {
		panic(err)
	}
	db.Close()
}
