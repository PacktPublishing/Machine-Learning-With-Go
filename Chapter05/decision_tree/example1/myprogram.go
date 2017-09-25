package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/trees"
)

func main() {

	// Read in the iris data set into golearn "instances".
	irisData, err := base.ParseCSVToInstances("iris.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	// This is to seed the random processes involved in building the
	// decision tree.
	rand.Seed(44111342)

	// We will use the ID3 algorithm to build our decision tree.  Also, we
	// will start with a parameter of 0.6 that controls the train-prune split.
	tree := trees.NewID3DecisionTree(0.6)

	// Use cross-fold validation to successively train and evaluate the model
	// on 5 folds of the data set.
	cv, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(irisData, tree, 5)
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
