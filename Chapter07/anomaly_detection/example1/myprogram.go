package main

import (
	"fmt"
	"log"

	"github.com/lytics/anomalyzer"
)

func main() {

	// Initialize an AnomalyzerConf value with
	// configurations such as which anomaly detection
	// methods we want to use.
	conf := &anomalyzer.AnomalyzerConf{
		Sensitivity: 0.1,
		UpperBound:  5,
		LowerBound:  anomalyzer.NA, // ignore the lower bound
		ActiveSize:  1,
		NSeasons:    4,
		Methods:     []string{"diff", "fence", "highrank", "lowrank", "magnitude"},
	}

	// Create a time series of periodic observations
	// as a slice of floats.  This could come from a
	// database or file, as utilized in earlier examples.
	ts := []float64{0.1, 0.2, 0.5, 0.12, 0.38, 0.9, 0.74}

	// Create a new anomalyzer based on the existing
	// time series values and configuration.
	anom, err := anomalyzer.NewAnomalyzer(conf, ts)
	if err != nil {
		log.Fatal(err)
	}

	// Supply a new observed value to the Anomalyzer.
	// The Anomalyzer will analyze the value in reference
	// to pre-existing values in the series and output
	// a probability of the value being anomalous.
	prob := anom.Push(15.2)
	fmt.Printf("Probability of 15.2 being anomalous: %0.2f\n", prob)

	prob = anom.Push(0.43)
	fmt.Printf("Probability of 0.33 being anomalous: %0.2f\n", prob)
}
