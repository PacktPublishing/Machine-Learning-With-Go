package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {

	// Open the CSV file.
	irisFile, err := os.Open("../data/iris_labeled.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	irisDF := dataframe.ReadCSV(irisFile)

	// As a sanity check, display the records to stdout.
	// Gota will format the dataframe for pretty printing.
	fmt.Println(irisDF)
}
