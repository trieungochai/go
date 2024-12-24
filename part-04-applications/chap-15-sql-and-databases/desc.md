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
