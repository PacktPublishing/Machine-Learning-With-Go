package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

type centroid []float64

func main() {

	// Pull in the CSV file.
	irisFile, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Create a dataframe from the CSV file.
	irisDF := dataframe.ReadCSV(irisFile)

	// Define the names of the three separate species contained in the CSV file.
	speciesNames := []string{
		"Iris-setosa",
		"Iris-versicolor",
		"Iris-virginica",
	}

	// Create a map to hold our centroid information.
	centroids := make(map[string]centroid)

	// Filter the dataset into three separate dataframes,
	// each corresponding to one of the Iris species.
	for _, species := range speciesNames {

		// Filer the original dataset.
		filter := dataframe.F{
			Colname:    "species",
			Comparator: "==",
			Comparando: species,
		}
		filtered := irisDF.Filter(filter)

		// Calculate the mean of features.
		summaryDF := filtered.Describe()

		// Put each dimension's mean into the corresponding centroid.
		var c centroid
		for _, feature := range summaryDF.Names() {

			// Skip the irrelevant columns.
			if feature == "column" || feature == "species" {
				continue
			}
			c = append(c, summaryDF.Col(feature).Float()[0])
		}

		// Add this centroid to our map.
		centroids[species] = c
	}

	// As a sanity check, output our centroids.
	for _, species := range speciesNames {
		fmt.Printf("%s centroid: %v\n", species, centroids[species])
	}
}
