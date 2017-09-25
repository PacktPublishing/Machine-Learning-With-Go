package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/mash/gokmeans"
)

func main() {

	// Open the driver dataset file.
	f, err := os.Open("fleet_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader.
	r := csv.NewReader(f)
	r.FieldsPerRecord = 3

	// Initialize a slice of gokmeans.Node's to
	// hold our input data.
	var data []gokmeans.Node

	// Loop over the records creating our slice of
	// gokmeans.Node's.
	for {

		// Read in our record and check for errors.
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// Skip the header.
		if record[0] == "Driver_ID" {
			continue
		}

		// Initialize a point.
		var point []float64

		// Fill in our point.
		for i := 1; i < 3; i++ {

			// Parse the float value.
			val, err := strconv.ParseFloat(record[i], 64)
			if err != nil {
				log.Fatal(err)
			}

			// Append this value to our point.
			point = append(point, val)
		}

		// Append our point to the data.
		data = append(data, gokmeans.Node{point[0], point[1]})
	}

	// Generate our clusters with k-means.
	success, centroids := gokmeans.Train(data, 2, 50)
	if !success {
		log.Fatal("Could not generate clusters")
	}

	// Output the centroids to stdout.
	fmt.Println("The centroids for our clusters are:")
	for _, centroid := range centroids {
		fmt.Println(centroid)
	}
}
