package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/gonum/stat"
	"github.com/kniren/gota/dataframe"
)

func main() {

	// Open the CSV file.
	passengersFile, err := os.Open("AirPassengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()

	// Create a dataframe from the CSV file.
	passengersDF := dataframe.ReadCSV(passengersFile)

	// Get the time and passengers as a slice of floats.
	passengers := passengersDF.Col("AirPassengers").Float()

	// Loop over various values of lag in the series.
	fmt.Println("Autocorrelation:")
	for i := 1; i < 11; i++ {

		// Calculate the autocorrelation.
		ac := acf(passengers, i)
		fmt.Printf("Lag %d period: %0.2f\n", i, ac)
	}
}

// acf calculates the autocorrelation for a series
// at the given lag.
func acf(x []float64, lag int) float64 {

	// Shift the series.
	xAdj := x[lag:len(x)]
	xLag := x[0 : len(x)-lag]

	// numerator will hold our accumulated numerator, and
	// denominator will hold our accumulated denominator.
	var numerator float64
	var denominator float64

	// Calculate the mean of our x values, which will be used
	// in each term of the autocorrelation.
	xBar := stat.Mean(x, nil)

	// Calculate the numerator.
	for idx, xVal := range xAdj {
		numerator += ((xVal - xBar) * (xLag[idx] - xBar))
	}

	// Calculate the denominator.
	for _, xVal := range x {
		denominator += math.Pow(xVal-xBar, 2)
	}

	return numerator / denominator
}
