# Connect to postgresSQL db and create tables/insert data

The example is related to [pgx - PostgreSQL Driver and Toolkit](https://github.com/jackc/pgx)

### Objective

Connect and create tables / insert data to PostgresSQL database.

### Basic flow 

1. Connect to a database
2. Get a file from a GitHub project with the SQL statements to create tables and insert data 
3. Execute the SQL statement
4. Verify one value with a query 

### Basic programming steps

1. Create HTTP request "github"
2. Define header
3. Create client
4. Invoke HTTP request
5. Verify the request status
6. Get only body from response
7. Convert body to json content
8. Extract and decode file content from json
9. Connect to a database
10. Create a sql statements from file content
11. Verify the created tables with a query	

### Some useful resources:

* [How to convert an HTTP response body to a string in Go](https://freshman.tech/snippets/go/http-response-to-string/)
* [Go by Example: Base64 Encoding](https://gobyexample.com/base64-encoding)
* [How to get a file via GitHub APIs](https://stackoverflow.com/questions/9272535/how-to-get-a-file-via-github-apis)
* [Go by Example: JSON](https://gobyexample.com/json)
* [Access HTTP response as string in Go](https://stackoverflow.com/questions/38673673/access-http-response-as-string-in-go)


### Run the example application

This are the step you need to follow to run the example.

> Note: You need a running PostgresSQL database somewhere

### Step 1: Git clone

```sh
git clone https://github.com/thomassuedbroecker/go-access-postgres-example.git
cd go-access-postgres-example
```

### Step 2: Create a mod file (that file exists)

```sh
cd gopostgressql
```

### Step 3: Set the enviornment variable

```sh
export DATABASE_URL="postgres://username:password@localhost:5432/database_name"
```

### Step 5: Execute the go program

```sh
go run  .
```


