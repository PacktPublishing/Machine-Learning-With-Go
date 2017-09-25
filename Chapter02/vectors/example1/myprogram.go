package main

import "fmt"

func main() {

	// Initialize a "vector" via a slice.
	var myvector []float64

	// Add a couple of components to the vector.
	myvector = append(myvector, 11.0)
	myvector = append(myvector, 5.2)

	fmt.Println(myvector)
}
