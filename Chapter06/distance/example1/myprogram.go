package main

import (
	"fmt"

	"github.com/gonum/floats"
)

func main() {

	// Calculate the Euclidean distance, specified here via
	// the last argument in the Distance function.
	distance := floats.Distance([]float64{1, 2}, []float64{3, 4}, 2)

	fmt.Printf("\nDistance: %0.2f\n\n", distance)
}
