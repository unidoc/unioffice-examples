// Copyright 2024 FoxyUtils ehf. All rights reserved.
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/spreadsheet"
)

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		panic(err)
	}
}

func main() {
	sliceData, err := readCsv("example-data.csv")
	if err != nil {
		log.Fatalf("error reading csv: %s", err)
	}

	kmSlice := strings.Split(sliceData[1][0], ";")
	dateSlice := strings.Split(sliceData[2][0], ";")

	ss := spreadsheet.New()
	defer ss.Close()
	sheet := ss.AddSheet()

	// Create all of our data
	row := sheet.AddRow()
	row.AddCell().SetString("Date")
	row.AddCell().SetString("Length")

	startRowNumber := row.RowNumber() + 1
	endRowNumber := startRowNumber + 1
	for r := 1; r < len(dateSlice); r++ {
		km, err := strconv.Atoi(kmSlice[r])
		if err != nil {
			log.Fatalf("unable to convert data into integer: %v", err)
		}

		row := sheet.AddRow()
		row.AddCell().SetString(dateSlice[r])
		row.AddCell().SetNumber(float64(km))

		endRowNumber = row.RowNumber()
	}

	// Charts need to reside in a drawing
	dwng := ss.AddDrawing()
	chart, anc := dwng.AddChart(spreadsheet.AnchorTypeTwoCell)
	// make it a bit wider than the default
	anc.SetWidthCells(10)

	lc := chart.AddLineChart()
	kmSeries := lc.AddSeries()
	kmSeries.SetText("KM")

	// Set a category axis reference on the first series to pull the dates
	kmSeries.CategoryAxis().SetLabelReference(fmt.Sprintf(`'Sheet 1'!A%d:A%d`, startRowNumber, endRowNumber))
	kmSeries.Values().SetReference(fmt.Sprintf(`'Sheet 1'!B%d:B%d`, startRowNumber, endRowNumber))

	// the line chart accepts up to two axes
	ca := chart.AddCategoryAxis()
	va := chart.AddValueAxis()
	lc.AddAxis(ca)
	lc.AddAxis(va)

	ca.SetCrosses(va)
	va.SetCrosses(ca)

	// add a title and legend
	title := chart.AddTitle()
	title.SetText("Length in KM")
	chart.AddLegend()

	// and finally add the chart to the sheet
	sheet.SetDrawing(dwng)

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	if err := ss.SaveToFile("output.xlsx"); err != nil {
		log.Fatalf("error saving: %s", err)
	}
}

// readCsv reads a csv file and returns a slice of slices of strings.
func readCsv(path string) ([][]string, error) {
	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.File
	file, err := os.Open(path)

	// Checks for the error
	if err != nil {
		return nil, err
	}

	// Closes the file
	defer file.Close()

	// The csv.NewReader() function is called in
	// which the object os.File passed as its parameter
	// and this creates a new csv.Reader that reads
	// from the file
	reader := csv.NewReader(file)

	return reader.ReadAll()
}
