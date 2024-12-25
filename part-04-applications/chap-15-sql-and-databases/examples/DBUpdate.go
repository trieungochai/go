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

	UpdateStatement := `
	UPDATE test
	SET name = $1
	WHERE id = $2
	`

	UpdateResult, UpdateResultErr := db.Exec(UpdateStatement, "well", 2)
	if UpdateResultErr != nil {
		panic(UpdateResultErr)
	}

	// To this end, we could use RowsAffected().
	// It will return the number of rows that were updated and any errors that were faced along the way.
	UpdatedRecords, UpdatedRecordsErr := UpdateResult.RowsAffected()
	if UpdatedRecordsErr != nil {
		panic(UpdatedRecordsErr)
	}

	fmt.Println("Number of records updated: ", UpdatedRecords)

	db.Close()
}
