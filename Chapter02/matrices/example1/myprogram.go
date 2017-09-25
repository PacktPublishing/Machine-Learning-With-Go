package main

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

func main() {

	// Create a flat representation of our matrix.
	data := []float64{1.2, -5.7, -2.4, 7.3}

	// Form our matrix.
	a := mat64.NewDense(2, 2, data)

	// As a sanity check, output the matrix to standard out.
	fa := mat64.Formatted(a, mat64.Prefix("    "))
	fmt.Printf("A = %v\n\n", fa)
}
