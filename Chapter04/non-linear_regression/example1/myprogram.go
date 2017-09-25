package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gonum/matrix/mat64"
)

func main() {

	// Open the training dataset file.
	f, err := os.Open("training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 4

	// Read in all of the CSV records
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// featureData will hold all the float values that will eventually be
	// used to form our matrix of features.
	featureData := make([]float64, 4*len(rawCSVData))
	yData := make([]float64, len(rawCSVData))

	// featureIndex and yIndex will track the current index of the matrix values.
	var featureIndex int
	var yIndex int

	// Sequentially move the rows into a slice of floats.
	for idx, record := range rawCSVData {

		// Skip the header row.
		if idx == 0 {
			continue
		}

		// Loop over the float columns.
		for i, val := range record {

			// Convert the value to a float.
			valParsed, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal("Could not parse float value")
			}

			if i < 3 {

				// Add an intercept to the model.
				if i == 0 {
					featureData[featureIndex] = 1
					featureIndex++
				}

				// Add the float value to the slice of feature floats.
				featureData[featureIndex] = valParsed
				featureIndex++
			}

			if i == 3 {

				// Add the float value to the slice of y floats.
				yData[yIndex] = valParsed
				yIndex++
			}

		}
	}

	// Form the matrices that will be input to our regression.
	features := mat64.NewDense(len(rawCSVData), 4, featureData)
	y := mat64.NewVector(len(rawCSVData), yData)

	if features != nil && y != nil {
		fmt.Println("Matrices formed for ridge regression")
	}
}
