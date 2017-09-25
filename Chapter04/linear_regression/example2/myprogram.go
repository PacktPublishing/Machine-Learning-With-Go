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

	// Open the advertising dataset file.
	f, err := os.Open("Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a dataframe from the CSV file.
	advertDF := dataframe.ReadCSV(f)

	// Create a histogram for each of the columns in the dataset.
	for _, colName := range advertDF.Names() {

		// Create a plotter.Values value and fill it with the
		// values from the respective column of the dataframe.
		plotVals := make(plotter.Values, advertDF.Nrow())
		for i, floatVal := range advertDF.Col(colName).Float() {
			plotVals[i] = floatVal
		}

		// Make a plot and set its title.
		p, err := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

		// Create a histogram of our values drawn
		// from the standard normal.
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
