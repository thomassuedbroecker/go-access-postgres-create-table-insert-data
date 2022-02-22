package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	b64 "encoding/base64"

	"github.com/jackc/pgx/v4"
)

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

func main() {

	var sDecFileContent []byte

	// ****************************
	// Get file with statements from GitHub

	// 1. Create HTTP request "github"
	// =======================
	//

	req, err := http.NewRequest("GET", "https://api.github.com/repos/IBM/multi-tenancy/contents/installapp/postgres-config/create-populate-tenant-a.sql", nil)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	// 2. Define header
	req.Header.Set("Accept", "application/json")

	// 3. Create client
	client := http.Client{}

	// 4. Invoke HTTP request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	// 5. Verify the request status
	if resp.StatusCode == http.StatusOK {

		// 6. Get only body from response
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		// 7. Convert body to json content
		var dat GitHubFile
		if err := json.Unmarshal(bodyBytes, &dat); err != nil {
			panic(err)
			os.Exit(1)
		}

		// 8. Extract and decode file content from json
		sDecFileContent, _ = b64.StdEncoding.DecodeString(dat.Content)
	}

	// 9. Connect to a database
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Connected to the DB: true [" + os.Getenv("DATABASE_URL") + "] \n")
		fmt.Println()
	}

	// 10. Create a sql statements from file content
	statement := string(sDecFileContent)
	_, err = conn.Exec(context.Background(), statement)

	if err != nil {
		fmt.Fprintf(os.Stderr, "File content for the statement: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("File content for the statement: true\n")
		fmt.Println()
	}

	// 11. Verify the created tables with a query
	var name string
	var price float64

	err = conn.QueryRow(context.Background(), "select name, price from product where name='Return of the Jedi'").Scan(&name, &price)
	if err != nil {
		fmt.Printf("Connected to the DB: true\n")
		fmt.Println()
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Return values of the Table: ", name, price)
		fmt.Println()
	}

	defer conn.Close(context.Background())
	fmt.Println("**********DONE********")
}
