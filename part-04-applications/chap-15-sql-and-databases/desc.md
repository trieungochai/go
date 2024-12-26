## Introduction

As a developer, it is impossible to get by without a proper understanding of persistent data storage and databases. Our applications process input and produce output, but most of the time, if not in all cases, a database is involved in the process.

This database can be in-memory (stored in the computer’s RAM) or file-based (a single file in a directory), and it can live on local or remote storage. A database engine can be installed locally, but it is also possible to use cloud providers, which allow you to use a database as a service; some of the cloud providers that offer several different database engine options are Azure, AWS, and Google Cloud.

---

### Connecting to databases

To connect to any database, we need at least 4 things to be in place:

- a host to connect to
- a database to connect to that is running on a port
- a username
- a password

The user needs to have appropriate privileges because we not only want to connect but we would like to perform specific operations, such as query, insert, or remove data, create or delete databases, and manage users and views.

In most cases, the database server supports multiple databases, and the databases hold one or more tables:

![databases-in-a-server](databases-in-a-server.png)

Imagine that the databases are logical containers that belong together.

---

### Create a new project

To connect, we need to get the appropriate module from GitHub, which needs internet connectivity. We need to issue the following command to get the package needed to interact with the Postgres instance:

```go
go get github.com/lib/pq
```

First, we will initialize our script:

```go
package main
import "fmt"
import "database/sql"
import _ "github.com/lib/pq"
// import _ <package name> is a special import statement that tells Go to import a package solely for its side effects.
```

Now that we have initialized our script, we can connect to our database:

```go
db, err := sql.Open("postgres", "user=postgres password=Start!123 host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
```

The API gives us an `Open()` function, which takes a variety of arguments.

- The `postgres` string, used as the first parameter in the `Open` function call, tells the function to use the Postgres driver to make the connection.
- The second argument is a so-called connection string, which holds the `user`, `password`, `host`, `port`, `dbname`, and `sslmode` arguments; these will be used to initialize the connection. In this example, we’re connecting to the local host marked by 127.0.0.1 on the default port of 5432, and we don’t use ssl. For production systems, people tend to change the default port and enforce encrypted traffic via ssl toward the database server; you should always follow the best practices concerning the type of database you’re working with.

The `Open()` function returns 2 values. One is for the database connection and the other is for the error, if one occurred during initialization.

We can check whether there were any errors by writing the following code:

```go
if err != nil {
  panic(err)
} else {
  fmt.Println("The connection to the DB was successfully initialized!")
}
```

The `panic()` function in Go is used to indicate that something went wrong unexpectedly, and we are not prepared to handle it gracefully, thus stopping the execution.

If the connection succeeds, we print out a message stating `The connection to the DB was successfully initialized!`. When you have a long-running application, it is worth incorporating a way to check whether the database is still reachable because due to intermittent network errors, you could lose the connection and fail to execute whatever you want to execute. This can be checked with the following small code snippet:

```go
connectivity := db.Ping()
if connectivity != nil{
  panic(err)
} else {
  fmt.Println("Good to go!")
}
```

You can run this check constantly on a different Go routine every few seconds. It will check if the database is on but also help keep the connection open; otherwise, it will go idle. This is a proactive solution as you check the status of the database connection.

In this case, we used the `panic()` function to indicate that the connection has been lost. Finally, once our job is done, we need to terminate our connection to the database to remove user sessions and free up resources. This can happen either if you are building a script that will run as a job, hence will run and finish, or if you are building a long-running service. In the first case, you can use the following command at the end of the script:

```go
db.Close()
```

This ensures that before terminating the script, the connection will be dropped. If you are building a long-running service, you don’t have a specific point in your code where you know that the script will terminate, but it can happen at any time. You can use the following code to ensure the connections are dropped:

```go
defer db.Close()
```

The difference is the scope. `db.Close()` will terminate the connection to the database once the execution arrives at the specific line, while `defer db.Close()` indicates that the database connection should be executed once the function in which it was called goes out of scope. The idiomatic way to do this is with `defer db.Close()`.

---

### Creating tables

The act of creating tables aims to make logical containers that persistently hold data that belongs together. The common goal is to provide a service for applications that make sense of it.

How do these database engines control who can access what data? There are 2 approaches:

- The 1st one is access control lists (`ACLs`), which is a simple yet powerful approach. ACL security logic tells us which user has which permissions, such as CREATE, UPDATE, and DELETE.
- The 2nd approach involves inheritance and roles. This is more robust and is better suited for big enterprises.

---

### Inserting data

Long ago, when the era of web applications backed by SQL databases started to bloom, some gutsy people invented the SQL injection attack. Here, a type of authentication is done against a database via SQL queries and, for example, after converting the password with mathematical magic into hash functions, the web app executes the query with the username and password coming from the input of the form.

Many servers executed something like this:

```sql
SELECT password FROM Auth WHERE username=<input from user>
```

Then, the password gets rehashed; if the 2 hashes match, the password is good for the user.

The problem with this came from the `<input from user>` part because if the attacker was smart enough, they could reformulate the query and run additional commands. Here’s an example:

```sql
SELECT password FROM Auth WHERE username=<input from user> OR '1'='1'
```

The problem with this query is that `OR '1' = '1'` always evaluates to `true`, and it does not matter what the username is; the user’s password hash would be returned. This can be further reused to formulate an additional attack. To prevent this, Go uses something called the `Prepare()` statement, which protects against these attacks.

Go has 2 types of substitutions:

- We use `WHERE col = $1` in the case of queries
- We use `VALUES($1,$2)` in the case of inserts or updates

`db.Prepare()` takes a SQL statement and imbues it with protection against SQL injection attacks. It works by restricting the values of the variable substitutions.

---

### Retrieving data

SQL injection does not only concern the data being inserted. It also concerns any data that is manipulated in the database. Retrieving data and, most importantly, retrieving it safely is also something we must prioritize and handle with proper caution.

When we query data, our results depend on the database we connect to and the table we would like to query. However, we must also mention that the security mechanisms that are implemented by the database engine may also prevent a successful query unless the user has appropriate privileges.

We can differentiate between 2 types of queries:

- Some queries do not take an argument, such as `SELECT \* FROM table`
- Some queries require you to specify filter criteria

Go provides 2 functions that allow you to query data. One is called `Query()` and the other is called `QueryRow()`. As a rule of thumb, you should remember that `Query()` is used to return any number of results, while `QueryRow()` is used when you expect to retrieve at most one row. You can also wrap them with the `Prepare()` statement.

---

### Updating existing data

When you are updating a row or multiple rows with Go, you are in trouble. The sql package does not provide any function called `Update()`; however, there is the `Exec()` function, which serves as a universal executor for your queries. You can execute `SELECT`, `UPDATE`, `DELETE`, or whatever you need to execute with this function.

---

### Deleting data

The analogy is the same as for the `UPDATE` statement of our records. We formulate a `DELETE` statement and execute it; we can technically modify the action of our `UPDATE` script to delete it from the database.

---

### Truncating and deleting table

To empty the table, we can simply formulate `DELETE` statements that match every record in our table and thus remove every single record from our table.

However, there is a more elegant way to do this: we can use the `TRUNCATE TABLE` SQL statement. The result of this statement is an empty table. We can use the `Exec()` function from our sql package for this.

The following statement will achieve a full TRUNCATE:

```go
emptyTable, emptyTableErr := db.Exec("TRUNCATE TABLE test")
if emptyTableErr != nil {
  panic(emptyTableErr)
}
```

The result of this is an empty table called test. To get rid of the table completely, we can modify our statement as follows:

```go
dropTable, dropTableErr := db.Exec("DROP TABLE test")
if dropTableErr != nil {
  panic(dropTableErr)
}
```

If you need a table but do not need any more old data, you might want to truncate it and carry on adding new data to the existing table. If you do not need the table anymore because you changed your schema, you might want to just delete it using the DROP command.

If we inspect our database engine, we won’t find any trace of the test table. This eradicated the whole table from the very face of the database.

---

### Adding users with GORM

So far, we’ve interacted with the database by writing some SQL queries directly. What we’ve done is create and run Go code, which was used to then run SQL code. This is perfectly fine, but there is also a way to run just Go code to interact with a SQL database. On top of this, the data that we are storing in the database will then be unwrapped into Go variables, and the content of a row might define the values of an instance of a Go struct. What we can do to improve and simplify the whole process is abstract the database even more and use an object-relational mapper (ORM).

This is a library that matches the tables and their relations as Go structs so that you can insert and retrieve data the same way you would instantiate and delete any instance of a Go struct. An ORM is not generally part of a language, and Go does not provide one by itself. There is, however, a set of third-party libraries, one of which is the de facto ORM for Go, and this is GORM.

To use GORM, we must import it. Here’s how:

```go
import (
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)
```

- The first loads the GORM library
- The second specifies the driver to use.

GORM can be used to interact with a lot of different database engines, including MySQL, Postgres, and SQLite. While the library itself is available from `gorm.io/gorm`, the specific way to interact with the engine is handled by the driver – in this case, the Postgres driver.

The next step will be to define a schema – that is, a Go struct representing what’s inside a table. Let’s define a struct representing a user:

```go
type User struct {
  gorm.Model
  FirstName  string
  LastName   string
  Email      string
}
```

We define a struct called `User` and we add some fields that will hold the first and last name of a user, together with their email address. The first important thing, however, is that we embed the `gorm.Model` struct into our struct, making it effectively a GORM model. This struct will add some fields, such as an ID, and set it as a primary key, as well as some other fields, such as creation and update date, and will also add some methods that will be used by the library to make it interact with a database.

To interact with the database, we must connect to it. Earlier, we saw how to connect to PostgreSQL; we will do something similar here:

```go
connection_string = "user=postgres password=Start!123 host=127.0.0.1 port=5432 dbname=postgres sslmode=disable"
db, err := gorm.Open(postgres.Open(connection_string), &gorm.Config{})
if err != nil {
   panic("failed to connect database")
}
```

We can use the same connection string as earlier, but we will do so inside the gorm.Open call, which allows GORM to interact with the underlying database engine.

So far, we haven’t created a table for the users, and we’ve seen how to create one using SQL and call it via Go. With GORM, we do not need to do that. After defining the type that will go inside the table that will hold users, we can have GORM create that table for us, if it does not exist already. We can do this with the following code:

```go
db.AutoMigrate(&User{})
```

This call ensures that there is a table holding users that contains all the required columns, and by default will call it users. There are ways to change the name of the table, but in general, it is better to follow the conventions. So, a table holding users’ data will be called users, while a struct holding the details of a user will be called User.

---

### Finding Users with GORM

Once we’ve added users, we would like to retrieve them. Let’s add a few other users using what we learned in the previous section:

```go
db.Create(&User{FirstName: "John", LastName: "Doe", Email: "john.doe@gmail.com"
db.Create(&User{FirstName: "James", LastName: "Smith", Email: "james.smith@gmail.com"})
```

Let’s assume that we had already inserted the record for John Smith. So, starting from a clean database and clean table, we should have users with IDs of 1, 2, and 3, respectively.
Now, we want to retrieve details about the first user we inserted. We can do that with the following command:

```go
var user User
db.First(&user, 1)
```

This will return the first user matching the condition where the user’s ID is equal to 1. The returned record is un-marshaled into the user variable, which is an instance of the User struct. We can search for every other user via their ID and substitute the number 1 with 2 or 3. This, however, is not very interesting, as we might not know the user’s ID but only their name or surname. Let’s see how to retrieve John Doe from his surname:

```go
db.First(&user, "last_name = ?", "Doe")
```

Note that we did not use `LastName` but last_name as GORM automatically transforms every attribute of the struct that’s camel case into snake case; this is the usual convention for database column names. The other important thing to notice is that we use two parameters:

```go
"last_name = ?" and "Doe”
```

The first one represents the column we want to search in, and we have a question mark after the equals sign. The question mark is a placeholder and will be replaced by the next parameter, which is Doe. As we have two people with the surname Smith, the function we just used will retrieve the first person with that surname, but this is not necessarily the one we are looking for. We could use the Last function, which returns the last result that matches the query, but we could have more users with the same surname. The solution for this is as follows:

```go
db.First(&user, "last_name = ? AND first_name= ?", "Smith", "James")
```

Here, we created a query that includes more conditions – the first few parameters express the condition, while the following parameters fill the values with placeholders.

The issue we could face here is that we might get confused with the names of the struct’s attributes and the actual column names. If we need to do a simple matching query, we can substitute the previous code with the following:

```go
db.First(&user, &User{FirstName: "James", LastName: "Smith"})”
```

Here, we just pass an instance of the User struct with a few attributes set, leaving the other ones to the default values.

These examples allow us to search for a specific record, but often, we need a list of objects. Of course, the First and Last functions return only one item, but GORM also gives us a function to return all the records that match our criteria. If the criteria is simply an ID, or if the field we search for is unique, we are better off sticking with First, but if our criteria are not unique, we should use the following function:

```go
var users []User
db.Find(&users, &User{LastName: "Smith"})”
```

The Find function returns all the matching records, but we cannot just un-marshal it into a single user instance. So, we must define a users variable, which is a slice of User instances, rather than using the previously seen user, which was an instance of a User struct.

This gives us an idea of how to use GORM to insert and retrieve data, but we’ve forgotten one important thing: errors. These functions are contacting the database, but the queries might somehow error for several reasons, and we need to control that. The previously seen function does not return an error but a pointer to the database struct, which we can use to get the errors:

```go
tx := db.Find(&users, &User{LastName: "Smith"})
if tx.Error != nil {
  fmt.Println(tx.Err”
```

Here, the tx variable stands for transaction and returns a set of values with a potential error among them. We can check if there is an error by comparing the tx.Error value with nil. When we use a transaction, whatever we do to the database is not definitive; it does not affect the state of the database that’s accessed by any other client, so any change is temporary. To make any change effective, we need to commit the transaction. In this case, we are just returning results, and not modifying the database, so we do not need to commit. “We are using the transactions because GORM returns a transaction from the Find call.
