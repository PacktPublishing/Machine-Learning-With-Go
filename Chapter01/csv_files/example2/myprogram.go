package main

import (
	"encoding/csv"
	"fmt"
	"io"
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
	reader.FieldsPerRecord = -1

	// rawCSVData will hold our successfully parsed rows.
	var rawCSVData [][]string

	// Read in the records one by one.
	for {

		// Read in a row. Check if we are at the end of the file.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Append the record to our data set.
		rawCSVData = append(rawCSVData, record)
	}

	fmt.Println(rawCSVData)
}
