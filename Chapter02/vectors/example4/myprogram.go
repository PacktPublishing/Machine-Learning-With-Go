package main

import (
	"fmt"

	"github.com/gonum/blas/blas64"
	"github.com/gonum/matrix/mat64"
)

func main() {

	// Initialize a couple of "vectors" represented as slices.
	vectorA := mat64.NewVector(3, []float64{11.0, 5.2, -1.3})
	vectorB := mat64.NewVector(3, []float64{-7.2, 4.2, 5.1})

	// Compute the dot product of A and B
	// (https://en.wikipedia.org/wiki/Dot_product).
	dotProduct := mat64.Dot(vectorA, vectorB)
	fmt.Printf("The dot product of A and B is: %0.2f\n", dotProduct)

	// Scale each element of A by 1.5.
	vectorA.ScaleVec(1.5, vectorA)
	fmt.Printf("Scaling A by 1.5 gives: %v\n", vectorA)

	// Compute the norm/length of B.
	normB := blas64.Nrm2(3, vectorB.RawVector())
	fmt.Printf("The norm/length of B is: %0.2f\n", normB)
}
