package main

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

func main() {

	// Create a new vector value.
	myvector := mat64.NewVector(2, []float64{11.0, 5.2})

	fmt.Println(myvector)
}
