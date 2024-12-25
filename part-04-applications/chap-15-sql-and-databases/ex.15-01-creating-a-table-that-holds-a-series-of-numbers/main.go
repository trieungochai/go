package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	var property string

	db, err := sql.Open("postgres", "user=postgres password=911225 host=localhost port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connection to the DB was successfully initialized!")
	}

	TableCreate := `
	CREATE TABLE numbers	
	(
		number integer NOT NULL,
		property text COLLATE pg_catalog."default" NOT NULL
	)
	WITH (
		OIDS = FALSE
	)
	TABLESPACE pg_default;

	ALTER TABLE numbers
		OWNER to postgres;
	`

	_, err = db.Exec(TableCreate)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The table called numbers was successfully created!")
	}

	insert, insertErr := db.Prepare("INSERT INTO numbers VALUES($1,$2)")
	if insertErr != nil {
		panic(insertErr)
	}

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			property = "Even"
		} else {
			property = "Odd"
		}
		_, err = insert.Exec(i, property)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("The number: ", i, "is", property)
		}
	}

	insert.Close()

	fmt.Println("The numbers are ready.")
	db.Close()
}
