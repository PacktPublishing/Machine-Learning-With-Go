package main

import (
	"encoding/csv"
	"image/color"
	"log"
	"os"
	"strconv"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"github.com/kniren/gota/dataframe"
)

func main() {

	// Open the CSV file.
	passengersFile, err := os.Open("AirPassengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()

	// Create a dataframe from the CSV file.
	passengersDF := dataframe.ReadCSV(passengersFile)

	// Extract the number of passengers and time columns
	// as slices of floats.
	passengerVals := passengersDF.Col("AirPassengers").Float()
	timeVals := passengersDF.Col("time").Float()

	// pts will hold the values for plotting.
	pts := make(plotter.XYs, passengersDF.Nrow()-1)

	// differenced will hold our differenced values
	// that will be output to a new CSV file.
	var differenced [][]string
	differenced = append(differenced, []string{"time", "differenced_passengers"})

	// Fill pts with data.
	for i := 1; i < len(passengerVals); i++ {
		pts[i-1].X = timeVals[i]
		pts[i-1].Y = passengerVals[i] - passengerVals[i-1]
		differenced = append(differenced, []string{
			strconv.FormatFloat(timeVals[i], 'f', -1, 64),
			strconv.FormatFloat(passengerVals[i]-passengerVals[i-1], 'f', -1, 64),
		})
	}

	// Create the plot.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "time"
	p.Y.Label.Text = "differenced passengers"
	p.Add(plotter.NewGrid())

	// Add the line plot points for the time series.
	l, err := plotter.NewLine(pts)
	if err != nil {
		log.Fatal(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	// Save the plot to a PNG file.
	p.Add(l)
	if err := p.Save(10*vg.Inch, 4*vg.Inch, "diff_passengers_ts.png"); err != nil {
		log.Fatal(err)
	}

	// Save the differenced data out to a new CSV.
	f, err := os.Create("diff_series.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.WriteAll(differenced)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
