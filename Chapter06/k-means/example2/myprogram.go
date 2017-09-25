package main

import (
	"image/color"
	"log"
	"os"

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

	// pts will hold the values for plotting
	pts := make(plotter.XYs, driverDF.Nrow())

	// Fill pts with data.
	for i, floatVal := range driverDF.Col("Speeding_Feature").Float() {
		pts[i].X = floatVal
		pts[i].Y = yVals[i]
	}

	// Create the plot.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "Speeding"
	p.Y.Label.Text = "Distance"
	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(3)

	// Save the plot to a PNG file.
	p.Add(s)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "fleet_data_scatter.png"); err != nil {
		log.Fatal(err)
	}
}
