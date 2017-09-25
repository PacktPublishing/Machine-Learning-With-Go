package main

import (
	"database/sql"
	"log"
	"os"

	// pq is the libary that allows us to connect
	// to postgres with databases/sql.
	_ "github.com/lib/pq"
)

func main() {

	// Get my postgres connection URL. I have it stored in
	// an environmental variable.
	pgURL := os.Getenv("PGURL")
	if pgURL == "" {
		log.Fatal("PGURL empty")
	}

	// Open a database value. Specify the postgres driver
	// for databases/sql.
	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Update some values.
	res, err := db.Exec("UPDATE iris SET species = 'setosa' WHERE species = 'Iris-setosa'")
	if err != nil {
		log.Fatal(err)
	}

	// See how many rows where updated.
	rowCount, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	// Output the number of rows to standard out.
	log.Printf("affected = %d\n", rowCount)
}
