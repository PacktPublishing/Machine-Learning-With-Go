package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	// Open the iris dataset file.
	f, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)

	// Assume we don't know the number of fields per line.  By setting
	// FieldsPerRecord negative, each row may have a variable
	// number of fields.
	reader.FieldsPerRecord = -1

	// Read in all of the CSV records.
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rawCSVData)
}
