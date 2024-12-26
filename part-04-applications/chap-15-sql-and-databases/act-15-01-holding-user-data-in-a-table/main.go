// we are going to create a table that is going to hold user information
// such as ID, Name, and Email.
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Users struct {
	id    int
	name  string
	email string
}

func main() {
	users := []Users{
		{1, "Jack Daniel's", "jack.daniels@27.com"},
		{2, "Spike Spiegel", "spike.spiegel@bebop.com"},
	}

	db, err := sql.Open("postgres", "user=postgres password=911225 host=localhost port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connection to the DB was successfully initialized!")
	}

	connectivity := db.Ping()
	if connectivity != nil {
		panic(connectivity)
	} else {
		fmt.Println("Good to go!")
	}

	TableCreate := `
	create table users
(
	id integer not null,
	name text collate pg_catalog."default" not null,
	email text collate  pg_catalog."default" not null .
	constraint "users_pkey" primary key (id)
)
with (
	oids = false
)
tablespace pg_default;
alter table users
	owner to postgres;
`

	_, err = db.Exec(TableCreate)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The table called users was successfully created!")
	}

	insert, insertErr := db.Prepare("INSERT INTO users VALUES($1, $2, $3)")
	if insertErr != nil {
		panic(insertErr)
	}

	for _, u := range users {
		_, err := insert.Exec(u.id, u.name, u.email)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("The user with name: ", u.name, "and email: ", u.email, "was successfully added!")
		}
	}

	insert.Close()

	TableUpdate := `UPDATE users SET email=$1 WHERE id=$2`

	update, updateErr := db.Prepare(TableUpdate)
	if updateErr {
		panic(updateErr)
	}
	_, err = update.Exec("jack.daniels@bebop.com")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The user's email address was successfully updated!")
	}

	update.Close()

	remove, removeErr := db.Prepare("DELETE FROM users WHERE id=$1")
	if removeErr != nil {
		panic(removeErr)
	}

	_, err = remove.Exec(2)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The second user was successfully removed!")
	}

	remove.Close()
	db.Close()
}
