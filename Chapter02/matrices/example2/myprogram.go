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

	// Get a single value from the matrix.
	val := a.At(0, 1)
	fmt.Printf("The value of a at (0,1) is: %.2f\n\n", val)

	// Get the values in a specific column.
	col := mat64.Col(nil, 0, a)
	fmt.Printf("The values in the 1st column are: %v\n\n", col)

	// Get the values in a kspecific row.
	row := mat64.Row(nil, 1, a)
	fmt.Printf("The values in the 2nd row are: %v\n\n", row)

	// Modify a single element.
	a.Set(0, 1, 11.2)

	// Modify an entire row.
	a.SetRow(0, []float64{14.3, -4.2})

	// Modify an entire column.
	a.SetCol(0, []float64{1.7, -0.3})

	// As a sanity check, output the matrix to standard out.
	fa := mat64.Formatted(a, mat64.Prefix("    "))
	fmt.Printf("A = %v\n\n", fa)
}
