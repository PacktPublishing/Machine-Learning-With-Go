package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Open the loan dataset file.
	f, err := os.Open("loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 2

	// Read in all of the CSV records
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Create the output file.
	f, err = os.Create("clean_loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a CSV writer.
	w := csv.NewWriter(f)

	// Sequentially move the rows writing out the parsed values.
	for idx, record := range rawCSVData {

		// Skip the header row.
		if idx == 0 {

			// Write the header to the output file.
			if err := w.Write([]string{"FICO_score", "class"}); err != nil {
				log.Fatal(err)
			}
			continue
		}

		// Initialize a slice to hold our parsed values.
		outRecord := make([]string, 2)

		// Parse and normalize the FICO score.
		score, err := strconv.ParseFloat(strings.Split(record[0], "-")[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		outRecord[0] = strconv.FormatFloat((score-640.0)/(830.0-640.0), 'f', 4, 64)

		// Parse the Interest rate class.
		rate, err := strconv.ParseFloat(strings.TrimSuffix(record[1], "%"), 64)
		if err != nil {
			log.Fatal(err)
		}

		if rate <= 12.0 {
			outRecord[1] = "1.0"

			// Write the record to the output file.
			if err := w.Write(outRecord); err != nil {
				log.Fatal(err)
			}
			continue
		}

		outRecord[1] = "0.0"

		// Write the record to the output file.
		if err := w.Write(outRecord); err != nil {
			log.Fatal(err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
