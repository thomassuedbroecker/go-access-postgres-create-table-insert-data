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

	var sDec []byte

	// ****************************
	// Get file with statements from GitHub

	// Create "github" request
	// =======================
	//
	// https://api.github.com/repos/IBM/multi-tenancy/contents/installapp/postgres-config/create-populate-tenant-a.sql
	//
	// * GitHub:           https://api.github.com/repos/
	// * Name:             "IBM/"
	// * Repo:             "multi-tenancy"
	// * GitHub:           /contents/
	// * Name of the file: "installapp/postgres-config/create-populate-tenant-a.sql"
	//
	req, err := http.NewRequest("GET", "https://api.github.com/repos/IBM/multi-tenancy/contents/installapp/postgres-config/create-populate-tenant-a.sql", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// 1. Define header
	req.Header.Set("Accept", "application/json")

	// 2. Create client
	client := http.Client{}

	// 3. Invoke request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	// 4. Get all response data including the header
	// respContent, err := httputil.DumpResponse(resp, true)
	// if err != nil {
	//		log.Fatalln(err)
	// }
	// fmt.Println("*****RESPONSE********")
	// fmt.Println(string(respContent))

	if resp.StatusCode == http.StatusOK {

		// 5. Get only body from response
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		//bodyString := string(bodyBytes)
		//fmt.Println("*****BODY********")
		//fmt.Println(string(bodyString))

		// 7. Convert to json content
		var dat GitHubFile
		if err := json.Unmarshal(bodyBytes, &dat); err != nil {
			panic(err)
		}

		//fmt.Println("*****Content********")
		//fmt.Println(dat.Content)

		//fmt.Println("*****Decode********")
		// 8. Get and decode the file content
		sDec, _ = b64.StdEncoding.DecodeString(dat.Content)
		//fmt.Println(string(sDec))
		//fmt.Println()
	}

	fmt.Println("**********START********")

	// 9. Connect to a database
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "1. Unable to connect to database: %v\n", err)
		fmt.Println("**********Error********")
		os.Exit(1)
	} else {
		fmt.Printf("1. Connected to the DB: true [" + os.Getenv("DATABASE_URL") + "] \n")
		fmt.Println()
	}

	// Create a sequence
	//statement := "CREATE SEQUENCE product_id_seq START 1;"
	statement := string(sDec)
	fmt.Println(statement)

	_, err = conn.Exec(context.Background(), statement)

	if err != nil {
		fmt.Fprintf(os.Stderr, "2. Sequence create statement: %v\n", err)
		fmt.Println()
	} else {
		fmt.Printf("2. Sequence create statement: true\n")
		fmt.Println()
	}

	//os.Exit(1)

	/*
		// Create a table
		statement = "CREATE TABLE product(id SERIAL PRIMARY KEY,price DECIMAL(14,2) NOT NULL,name TEXT NOT NULL,description TEXT NOT NULL,image TEXT NOT NULL);"
		_, err = conn.Exec(context.Background(), statement)

		if err != nil {
			fmt.Fprintf(os.Stderr, "3. Table create statement: %v\n", err)
			fmt.Println()
		} else {
			fmt.Printf("3. Table create statement: true\n")
			fmt.Println()
		}

		// Insert a value
		statement = "INSERT INTO product VALUES (nextval('product_id_seq'), 29.99, 'Return of the Jedi', 'Episode 6, Luke has the final confrontation with his father!', 'images/Return.jpg');"
		_, err = conn.Exec(context.Background(), statement)

		if err != nil {
			fmt.Fprintf(os.Stderr, "4. Insert statement: %v\n", err)
			fmt.Println()
		} else {
			fmt.Printf("4. Insert statement: true\n")
			fmt.Println()
		}
	*/
	// Query a value
	var name string
	var price float64

	err = conn.QueryRow(context.Background(), "select name, price from product where name='Return of the Jedi'").Scan(&name, &price)
	if err != nil {
		fmt.Printf("Connected to the DB: true\n")
		fmt.Println()
		fmt.Fprintf(os.Stderr, "5 QueryRow failed: %v\n", err)
		fmt.Println("**********Error********")
		os.Exit(1)
	} else {
		fmt.Println("Return values of the Table: ", name, price)
		fmt.Println()
	}

	defer conn.Close(context.Background())
	fmt.Println("**********DONE********")

}
