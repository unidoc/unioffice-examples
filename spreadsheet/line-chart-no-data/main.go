// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"log"
	"os"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/spreadsheet"
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
	ss := spreadsheet.New()
	defer ss.Close()
	sheet := ss.AddSheet()

	dwng := ss.AddDrawing()
	chart, anc := dwng.AddChart(spreadsheet.AnchorTypeTwoCell)
	anc.MoveTo(0, 0)
	anc.SetWidth(10)

	// No cell data needed, we can supply data directly to the chart
	lc := chart.AddLineChart()
	priceSeries := lc.AddSeries()
	priceSeries.SetText("Price")
	priceSeries.CategoryAxis().SetValues([]string{"Prod 1", "Prod 2", "Prod 3", "Prod 4", "Prod 5"})
	priceSeries.Values().SetValues([]float64{5, 4, 3, 9, 2})

	soldSeries := lc.AddSeries()
	soldSeries.SetText("Sold")
	soldSeries.Values().SetValues([]float64{1, 2, 3, 4, 5})

	totalSeries := lc.AddSeries()
	totalSeries.SetText("Total")
	totalSeries.Values().SetValues([]float64{9, 2, 1, 8, 1})

	// the line chart accepts up to two axes
	ca := chart.AddCategoryAxis()
	va := chart.AddValueAxis()
	lc.AddAxis(ca)
	lc.AddAxis(va)

	ca.SetCrosses(va)
	va.SetCrosses(ca)

	// add a title and legend
	title := chart.AddTitle()
	title.SetText("Items Sold")
	chart.AddLegend()

	// and finally add the chart to the sheet
	sheet.SetDrawing(dwng)

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}
	ss.SaveToFile("line-chart-no-data.xlsx")
}
