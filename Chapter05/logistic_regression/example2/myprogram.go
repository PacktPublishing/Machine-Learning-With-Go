package main

import (
	"image/color"
	"log"
	"math"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
)

func main() {

	// Create a new plot.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.Title.Text = "Logistic Function"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "f(x)"

	// Create the plotter function.
	logisticPlotter := plotter.NewFunction(func(x float64) float64 { return logistic(x) })
	logisticPlotter.Color = color.RGBA{B: 255, A: 255}

	// Add the plotter function to the plot.
	p.Add(logisticPlotter)

	// Set the axis ranges.  Unlike other data sets,
	// functions don't set the axis ranges automatically
	// since functions don't necessarily have a
	// finite range of x and y values.
	p.X.Min = -10
	p.X.Max = 10
	p.Y.Min = -0.1
	p.Y.Max = 1.1

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "logistic.png"); err != nil {
		log.Fatal(err)
	}
}

// logistic implements the logistic function, which
// is used in logistic regression.
func logistic(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}
