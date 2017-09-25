package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"github.com/kniren/gota/dataframe"
)

func main() {

	// Open the CSV file.
	driverDataFile, err := os.Open("fleet_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer driverDataFile.Close()

	// Create a dataframe from the CSV file.
	driverDF := dataframe.ReadCSV(driverDataFile)

	// Use the Describe method to calculate summary statistics
	// for all of the columns in one shot.
	driverSummary := driverDF.Describe()

	// Output the summary statistics to stdout.
	fmt.Println(driverSummary)

	// Create a histogram for each of the columns in the dataset.
	for _, colName := range driverDF.Names() {

		// Create a plotter.Values value and fill it with the
		// values from the respective column of the dataframe.
		plotVals := make(plotter.Values, driverDF.Nrow())
		for i, floatVal := range driverDF.Col(colName).Float() {
			plotVals[i] = floatVal
		}

		// Make a plot and set its title.
		p, err := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.Title.Text = fmt.Sprintf("Histogram of %s", colName)

		// Create a histogram of our values.
		h, err := plotter.NewHist(plotVals, 16)
		if err != nil {
			log.Fatal(err)
		}

		// Normalize the histogram.
		h.Normalize(1)

		// Add the histogram to the plot.
		p.Add(h)

		// Save the plot to a PNG file.
		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
			log.Fatal(err)
		}
	}
}
