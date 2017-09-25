package main

import (
	"image/color"
	"log"
	"os"

	"github.com/gonum/floats"
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
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
	yVals := driverDF.Col("Distance_Feature").Float()

	// clusterOne and clusterTwo will hold the values for plotting.
	var clusterOne [][]float64
	var clusterTwo [][]float64

	// Fill the clusters with data.
	for i, xVal := range driverDF.Col("Speeding_Feature").Float() {
		distanceOne := floats.Distance([]float64{yVals[i], xVal}, []float64{50.05, 8.83}, 2)
		distanceTwo := floats.Distance([]float64{yVals[i], xVal}, []float64{180.02, 18.29}, 2)
		if distanceOne < distanceTwo {
			clusterOne = append(clusterOne, []float64{xVal, yVals[i]})
			continue
		}
		clusterTwo = append(clusterTwo, []float64{xVal, yVals[i]})
	}

	// pts* will hold the values for plotting
	ptsOne := make(plotter.XYs, len(clusterOne))
	ptsTwo := make(plotter.XYs, len(clusterTwo))

	// Fill pts with data.
	for i, point := range clusterOne {
		ptsOne[i].X = point[0]
		ptsOne[i].Y = point[1]
	}

	for i, point := range clusterTwo {
		ptsTwo[i].X = point[0]
		ptsTwo[i].Y = point[1]
	}

	// Create the plot.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "Speeding"
	p.Y.Label.Text = "Distance"
	p.Add(plotter.NewGrid())

	sOne, err := plotter.NewScatter(ptsOne)
	if err != nil {
		log.Fatal(err)
	}
	sOne.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	sOne.GlyphStyle.Radius = vg.Points(3)

	sTwo, err := plotter.NewScatter(ptsTwo)
	if err != nil {
		log.Fatal(err)
	}
	sTwo.GlyphStyle.Color = color.RGBA{B: 255, A: 255}
	sTwo.GlyphStyle.Radius = vg.Points(3)

	// Save the plot to a PNG file.
	p.Add(sOne, sTwo)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "fleet_data_clusters.png"); err != nil {
		log.Fatal(err)
	}
}
