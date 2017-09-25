package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gonum/floats"
	"github.com/kniren/gota/dataframe"
)

func main() {

	// Open the driver dataset file.
	f, err := os.Open("fleet_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a dataframe from the CSV file.
	driverDF := dataframe.ReadCSV(f)

	// Extract the distance column.
	distances := driverDF.Col("Distance_Feature").Float()

	// clusterOne and clusterTwo will hold the values for plotting.
	var clusterOne [][]float64
	var clusterTwo [][]float64

	// Fill the clusters with data.
	for i, speed := range driverDF.Col("Speeding_Feature").Float() {
		distanceOne := floats.Distance([]float64{distances[i], speed}, []float64{50.05, 8.83}, 2)
		distanceTwo := floats.Distance([]float64{distances[i], speed}, []float64{180.02, 18.29}, 2)
		if distanceOne < distanceTwo {
			clusterOne = append(clusterOne, []float64{distances[i], speed})
			continue
		}
		clusterTwo = append(clusterTwo, []float64{distances[i], speed})
	}

	// Output our within cluster metrics.
	fmt.Printf("\nCluster 1 Metric: %0.2f\n", withinClusterMean(clusterOne, []float64{50.05, 8.83}))
	fmt.Printf("\nCluster 2 Metric: %0.2f\n", withinClusterMean(clusterTwo, []float64{180.02, 18.29}))
}

// withinClusterMean calculates the mean distance between
// points in a cluster and the centroid of the cluster.
func withinClusterMean(cluster [][]float64, centroid []float64) float64 {

	// meanDistance will hold our result.
	var meanDistance float64

	// Loop over the points in the cluster.
	for _, point := range cluster {
		meanDistance += floats.Distance(point, centroid, 2) / float64(len(cluster))
	}

	return meanDistance
}
