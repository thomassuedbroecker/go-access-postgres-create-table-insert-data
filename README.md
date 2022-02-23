# Connect to postgresSQL db and create tables/insert data

That example is related to [pgx - PostgreSQL Driver and Toolkit](https://github.com/jackc/pgx)

### Objective

Connect to database, create tables, insert data.

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


### Understand the GitHub url format

For more details visit the [GitHub public APIs ](https://github.com/public-apis/public-apis)

* Example url we use: `https://api.github.com/repos/IBM/multi-tenancy/contents/installapp/postgres-config/create-populate-tenant-a.sql`

Mapping to the GitHub API endpoint: `https://api.github.com/repos/$NAME/$REPO/contents/$FILENAME`

These are the related values to example endpoint above:

* GitHub API:       https://api.github.com/repos/
* Name:             "IBM/"
* Repo:             "multi-tenancy"
* GitHub API:           /contents/
* Filename: "installapp/postgres-config/create-populate-tenant-a.sql"

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

### Step 3: Set the environment variable

```sh
export DATABASE_URL="postgres://username:password@localhost:5432/database_name"
```

### Step 5: Execute the go program

```sh
go run  .
```

### Additional information

#### GitHub configration

* The GO struct which does reflect the GitHub response JSON structure

```go
type GitHubFile struct {
	Name         string
	Path         string
	Sha          string
	Size         uint64
	Url          string
	Html_url     string
	Git_url      string
	Download_url string
	Type         string
	Content      string
}
```

* The JSON structure of the GitHub response

```json
{"name":"create-populate-tenant-a.sql",
 "path":"installapp/postgres-config/create-populate-tenant-a.sql",
 "sha":"95d9bd66b1d9bb18c6cb2e36f84bc91396713710",
 "size":4296,
 "url":"https://api.github.com/repos/IBM/multi-tenancy/contents/installapp/postgres-config/create-populate-tenant-a.sql?ref=main",
 "html_url":"https://github.com/IBM/multi-tenancy/blob/main/installapp/postgres-config/create-populate-tenant-a.sql",
 "git_url":"https://api.github.com/repos/IBM/multi-tenancy/git/blobs/95d9bd66b1d9bb18c6cb2e36f84bc91396713710",
 "download_url":"https://raw.githubusercontent.com/IBM/multi-tenancy/main/installapp/postgres-config/create-populate-tenant-a.sql",
 "type":"file",
 "content":"Q1JFQVRFIFNFUVVFTkNFI...",
 "encoding":"base64",
 "_links":{"self":"https://api.github.com/repos/IBM/multi-tenancy/contents/installapp/postgres-config/create-populate-tenant-a.sql?ref=main",
 "git":"https://api.github.com/repos/IBM/multi-tenancy/git/blobs/95d9bd66b1d9bb18c6cb2e36f84bc91396713710",
 "html":"https://github.com/IBM/multi-tenancy/blob/main/installapp/postgres-config/create-populate-tenant-a.sql"
}
```

