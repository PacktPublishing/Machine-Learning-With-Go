package main

import (
	"fmt"
	"log"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/filters"
	"github.com/sjwhitworth/golearn/naive"
)

func main() {

	// Read in the loan training data set into golearn "instances".
	trainingData, err := base.ParseCSVToInstances("training.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize a new Naive Bayes classifier.
	nb := naive.NewBernoulliNBClassifier()

	// Fit the Naive Bayes classifier.
	nb.Fit(convertToBinary(trainingData))

	// Read in the loan test data set into golearn "instances".
	testData, err := base.ParseCSVToInstances("test.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	// Make our predictions.
	predictions := nb.Predict(convertToBinary(testData))

	// Generate a Confusion Matrix.
	cm, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the accuracy.
	accuracy := evaluation.GetAccuracy(cm)
	fmt.Printf("\nAccuracy: %0.2f\n\n", accuracy)
}

// convertToBinary utilizes built in golearn functionality to
// convert our labels to a binary label format.
func convertToBinary(src base.FixedDataGrid) base.FixedDataGrid {
	b := filters.NewBinaryConvertFilter()
	attrs := base.NonClassAttributes(src)
	for _, a := range attrs {
		b.AddAttribute(a)
	}
	b.Train()
	ret := base.NewLazilyFilteredInstances(src, b)
	return ret
}
