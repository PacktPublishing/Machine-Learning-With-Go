package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/kniren/gota/dataframe"
	"github.com/sajari/regression"
)

func main() {

	// Open the CSV file.
	passengersFile, err := os.Open("log_diff_series.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()

	// Create a dataframe from the CSV file.
	passengersDF := dataframe.ReadCSV(passengersFile)

	// Get the time and passengers as a slice of floats.
	passengers := passengersDF.Col("log_differenced_passengers").Float()

	// Calculate the coefficients for lag 1 and 2 and
	// our error.
	coeffs, intercept := autoregressive(passengers, 2)

	// Output the AR(2) model to stdout.
	fmt.Printf("\nlog(x(t)) - log(x(t-1)) = %0.6f + lag1*%0.6f + lag2*%0.6f\n\n", intercept, coeffs[0], coeffs[1])
}

// autoregressive calculates an AR model for a series
// at a given order.
func autoregressive(x []float64, lag int) ([]float64, float64) {

	// Create a regresssion.Regression value needed to train
	// a model using github.com/sajari/regression.
	var r regression.Regression
	r.SetObserved("x")

	// Define the current lag and all of the intermediate lags.
	for i := 0; i < lag; i++ {
		r.SetVar(i, "x"+strconv.Itoa(i))
	}

	// Shift the series.
	xAdj := x[lag:len(x)]

	// Loop over the series creating the data set
	// for the regression.
	for i, xVal := range xAdj {

		// Loop over the intermediate lags to build up
		// our independent variables.
		laggedVariables := make([]float64, lag)
		for idx := 1; idx <= lag; idx++ {

			// Get the lagged series variables.
			laggedVariables[idx-1] = x[lag+i-idx]
		}

		// Add these points to the regression value.
		r.Train(regression.DataPoint(xVal, laggedVariables))
	}

	// Fit the regression.
	r.Run()

	// coeff hold the coefficients for our lags.
	var coeff []float64
	for i := 1; i <= lag; i++ {
		coeff = append(coeff, r.Coeff(i))
	}

	return coeff, r.Coeff(0)
}
