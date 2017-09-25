package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

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
	reader.FieldsPerRecord = 2

	// Read in all of the CSV records
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// featureData and labels will hold all the float values that
	// will eventually be used in our training.
	featureData := make([]float64, 2*(len(rawCSVData)-1))
	labels := make([]float64, len(rawCSVData)-1)

	// featureIndex will track the current index of the features
	// matrix values.
	var featureIndex int

	// Sequentially move the rows into the slices of floats.
	for idx, record := range rawCSVData {

		// Skip the header row.
		if idx == 0 {
			continue
		}

		// Add the FICO score feature.
		featureVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		featureData[featureIndex] = featureVal

		// Add an intercept.
		featureData[featureIndex+1] = 1.0

		// Increment our feature row.
		featureIndex += 2

		// Add the class label.
		labelVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		labels[idx-1] = labelVal
	}

	// Form a matrix from the features.
	features := mat64.NewDense(len(rawCSVData)-1, 2, featureData)

	// Train the logistic regression model.
	weights := logisticRegression(features, labels, 1000, 0.3)

	// Output the Logistic Regression model formula to stdout.
	formula := "p = 1 / ( 1 + exp(- m1 * FICO.score - m2) )"
	fmt.Printf("\n%s\n\nm1 = %0.2f\nm2 = %0.2f\n\n", formula, weights[0], weights[1])
}

// logistic implements the logistic function, which
// is used in logistic regression.
func logistic(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

// logisticRegression fits a logistic regression model
// for the given data.
func logisticRegression(features *mat64.Dense, labels []float64, numSteps int, learningRate float64) []float64 {

	// Initialize random weights.
	_, numWeights := features.Dims()
	weights := make([]float64, numWeights)

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for idx, _ := range weights {
		weights[idx] = r.Float64()
	}

	// Iteratively optimize the weights.
	for i := 0; i < numSteps; i++ {

		// Initialize a variable to accumulate error for this iteration.
		var sumError float64

		// Make predictions for each label and accumlate error.
		for idx, label := range labels {

			// Get the features corresponding to this label.
			featureRow := mat64.Row(nil, idx, features)

			// Calculate the error for this iteration's weights.
			pred := logistic(featureRow[0]*weights[0] + featureRow[1]*weights[1])
			predError := label - pred
			sumError += math.Pow(predError, 2)

			// Update the feature weights.
			for j := 0; j < len(featureRow); j++ {
				weights[j] += learningRate * predError * pred * (1 - pred) * featureRow[j]
			}
		}
	}

	return weights
}
