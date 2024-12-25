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

	DeleteStatement := `
	DELETE FROM test
	WHERE id = $1
	`
	DeleteResult, DeleteResultErr := db.Exec(DeleteStatement, 2)
	if DeleteResultErr != nil {
		panic(DeleteResultErr)
	}
	DeleteRecords, DeleteRecordsErr := DeleteResult.RowsAffected()
	if DeleteRecordsErr != nil {
		panic(DeleteRecordsErr)
	}

	fmt.Println("Number of records deleted: ", DeleteRecords)

	db.Close()
}
