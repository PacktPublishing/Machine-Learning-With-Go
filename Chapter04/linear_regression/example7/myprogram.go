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

	// Open the advertising dataset file.
	f, err := os.Open("Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a dataframe from the CSV file.
	advertDF := dataframe.ReadCSV(f)

	// Extract the target column.
	yVals := advertDF.Col("Sales").Float()

	// pts will hold the values for plotting.
	pts := make(plotter.XYs, advertDF.Nrow())

	// ptsPred will hold the predicted values for plotting.
	ptsPred := make(plotter.XYs, advertDF.Nrow())

	// Fill pts with data.
	for i, floatVal := range advertDF.Col("TV").Float() {
		pts[i].X = floatVal
		pts[i].Y = yVals[i]
		ptsPred[i].X = floatVal
		ptsPred[i].Y = predict(floatVal)
	}

	// Create the plot.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "TV"
	p.Y.Label.Text = "Sales"
	p.Add(plotter.NewGrid())

	// Add the scatter plot points for the observations.
	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(3)

	// Add the line plot points for the predictions.
	l, err := plotter.NewLine(ptsPred)
	if err != nil {
		log.Fatal(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	// Save the plot to a PNG file.
	p.Add(s, l)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "regression_line.png"); err != nil {
		log.Fatal(err)
	}
}

// predict uses our trained regression model to made a prediction.
func predict(tv float64) float64 {
	return 7.07 + tv*0.05
}
