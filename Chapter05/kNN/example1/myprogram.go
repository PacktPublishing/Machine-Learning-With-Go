package main

import (
	"fmt"
	"log"
	"math"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func main() {

	// Read in the iris data set into golearn "instances".
	irisData, err := base.ParseCSVToInstances("iris.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize a new KNN classifier.  We will use a simple
	// Euclidean distance measure and k=2.
	knn := knn.NewKnnClassifier("euclidean", "linear", 2)

	// Use cross-fold validation to successively train and evalute the model
	// on 5 folds of the data set.
	cv, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(irisData, knn, 5)
	if err != nil {
		log.Fatal(err)
	}

	// Get the mean, variance and standard deviation of the accuracy for the
	// cross validation.
	mean, variance := evaluation.GetCrossValidatedMetric(cv, evaluation.GetAccuracy)
	stdev := math.Sqrt(variance)

	// Output the cross metrics to standard out.
	fmt.Printf("\nAccuracy\n%.2f (+/- %.2f)\n\n", mean, stdev*2)
}
