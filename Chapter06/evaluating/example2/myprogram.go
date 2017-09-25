package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gonum/floats"
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

	// Create a map to hold the filtered dataframe for each cluster.
	clusters := make(map[string]dataframe.DataFrame)

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

		// Add the filtered dataframe to the map of clusters.
		clusters[species] = filtered

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

	// Convert our labels into a slice of strings and create a slice
	// of float column names for convenience.
	labels := irisDF.Col("species").Records()
	floatColumns := []string{
		"sepal_length",
		"sepal_width",
		"petal_length",
		"petal_width",
	}

	// Loop over the records accumulating the average silhouette coefficient.
	var silhouette float64

	for idx, label := range labels {

		// a will store our accumulated value for a.
		var a float64

		// Loop over the data points in the same cluster.
		for i := 0; i < clusters[label].Nrow(); i++ {

			// Get the data point for comparison.
			current := dfFloatRow(irisDF, floatColumns, idx)
			other := dfFloatRow(clusters[label], floatColumns, i)

			// Add to a.
			a += floats.Distance(current, other, 2) / float64(clusters[label].Nrow())
		}

		// Determine the nearest other cluster.
		var otherCluster string
		var distanceToCluster float64
		for _, species := range speciesNames {

			// Skip the cluster containing the data point.
			if species == label {
				continue
			}

			// Calculate the distance to the cluster from the current cluster.
			distanceForThisCluster := floats.Distance(centroids[label], centroids[species], 2)

			// Replace the current cluster if relevant.
			if distanceToCluster == 0.0 || distanceForThisCluster < distanceToCluster {
				otherCluster = species
				distanceToCluster = distanceForThisCluster
			}
		}

		// b will store our accumulated value for b.
		var b float64

		// Loop over the data points in the nearest other cluster.
		for i := 0; i < clusters[otherCluster].Nrow(); i++ {

			// Get the data point for comparison.
			current := dfFloatRow(irisDF, floatColumns, idx)
			other := dfFloatRow(clusters[otherCluster], floatColumns, i)

			// Add to b.
			b += floats.Distance(current, other, 2) / float64(clusters[otherCluster].Nrow())
		}

		// Add to the average silhouette coefficient.
		if a > b {
			silhouette += ((b - a) / a) / float64(len(labels))
		}
		silhouette += ((b - a) / b) / float64(len(labels))
	}

	// Output the final average silhouette coeffcient to stdout.
	fmt.Printf("\nAverage Silhouette Coefficient: %0.2f\n\n", silhouette)
}

// dfFloatRow retrieves a slice of float values from a DataFrame
// at the given index and for the given column names.
func dfFloatRow(df dataframe.DataFrame, names []string, idx int) []float64 {
	var row []float64
	for _, name := range names {
		row = append(row, df.Col(name).Float()[idx])
	}
	return row
}
