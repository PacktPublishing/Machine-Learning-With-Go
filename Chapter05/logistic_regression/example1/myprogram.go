package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(logistic(1.0))
}

// logistic implements the logistic function, which
// is used in logistic regression.
func logistic(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}
