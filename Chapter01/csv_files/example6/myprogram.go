package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {

	// Pull in the CSV file.
	irisFile, err := os.Open("../data/iris_labeled.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	irisDF := dataframe.ReadCSV(irisFile)

	// Create a filter for the dataframe.
	filter := dataframe.F{
		Colname:    "species",
		Comparator: "==",
		Comparando: "Iris-versicolor",
	}

	// Filter the dataframe to see only the rows where
	// the iris species is "Iris-versicolor".
	versicolorDF := irisDF.Filter(filter)
	if versicolorDF.Err != nil {
		log.Fatal(versicolorDF.Err)
	}

	// Output the results to standard out.
	fmt.Println(versicolorDF)

	// Filter the dataframe again, but only select out the
	// sepal_width and species columns.
	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"})
	fmt.Println(versicolorDF)

	// Filter and select the dataframe again, but only display
	// the first three results.
	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"}).Subset([]int{0, 1, 2})
	fmt.Println(versicolorDF)

}
