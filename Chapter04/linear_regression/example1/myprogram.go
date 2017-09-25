package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {

	// Open the CSV file.
	advertFile, err := os.Open("Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer advertFile.Close()

	// Create a dataframe from the CSV file.
	advertDF := dataframe.ReadCSV(advertFile)

	// Use the Describe method to calculate summary statistics
	// for all of the columns in one shot.
	advertSummary := advertDF.Describe()

	// Output the summary statistics to stdout.
	fmt.Println(advertSummary)
}
