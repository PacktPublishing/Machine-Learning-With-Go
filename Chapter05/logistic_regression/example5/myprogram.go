package main

import (
	"bufio"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {

	// Open the clean loan dataset file.
	f, err := os.Open("clean_loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	loanDF := dataframe.ReadCSV(f)

	// Calculate the number of elements in each set.
	trainingNum := (4 * loanDF.Nrow()) / 5
	testNum := loanDF.Nrow() / 5
	if trainingNum+testNum < loanDF.Nrow() {
		trainingNum++
	}

	// Create the subset indices.
	trainingIdx := make([]int, trainingNum)
	testIdx := make([]int, testNum)

	// Enumerate the training indices.
	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}

	// Enumerate the test indices.
	for i := 0; i < testNum; i++ {
		testIdx[i] = trainingNum + i
	}

	// Create the subset dataframes.
	trainingDF := loanDF.Subset(trainingIdx)
	testDF := loanDF.Subset(testIdx)

	// Create a map that will be used in writing the data
	// to files.
	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
	}

	// Create the respective files.
	for idx, setName := range []string{"training.csv", "test.csv"} {

		// Save the filtered dataset file.
		f, err := os.Create(setName)
		if err != nil {
			log.Fatal(err)
		}

		// Create a buffered writer.
		w := bufio.NewWriter(f)

		// Write the dataframe out as a CSV.
		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}
}
