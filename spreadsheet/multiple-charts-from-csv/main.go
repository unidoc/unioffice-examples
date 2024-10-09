// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/spreadsheet"
)

func init() {
	err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully set the Unidoc License Key")
}

type SalesData struct {
	DateOfSale      time.Time
	QuantitySold    float64
	SalePrice       float64
	SalespersonName string
}

func main() {
	filePath := "./data/test-data.csv"
	data, header, err := load_csv(filePath)
	if err != nil {
		panic("failed to load data ")
	}

	ss := spreadsheet.New()
	defer ss.Close()
	sheet := ss.AddSheet()

	// Add header row
	row := sheet.AddRow()
	for _, h := range header {
		row.AddCell().SetString(h)
	}
	row.AddCell().SetString("Cumulative Sum By Time")
	row.AddCell().SetString("Cumulative Sum By Person")
	// Add data rows.
	cumulativeSumsByPerson := make(map[string]float64)

	totalSum := 0.0
	for _, dataRow := range data {
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.SetDate(dataRow.DateOfSale)
		cellStyle := ss.StyleSheet.AddCellStyle()
		cellStyle.SetNumberFormatStandard(spreadsheet.StandardFormat15) // format in d-mmm-yy
		cell.SetStyle(cellStyle)

		row.AddCell().SetNumber(dataRow.QuantitySold)
		row.AddCell().SetNumber(dataRow.SalePrice)
		row.AddCell().SetString(dataRow.SalespersonName)
		totalSum += dataRow.SalePrice
		row.AddCell().SetNumber(totalSum)

		cumulativeSumsByPerson[dataRow.SalespersonName] += dataRow.SalePrice
		row.AddCell().SetNumber(cumulativeSumsByPerson[dataRow.SalespersonName])
	}

	drawing := ss.AddDrawing()

	// Create Line Chart showing daily sales.
	chart, anc := drawing.AddChart(spreadsheet.AnchorTypeTwoCell)
	// Set width to 8 cells
	anc.SetWidthCells(10)
	anc.MoveTo(7, 1)

	lc := chart.AddLineChart()
	dailyPriceSeries := lc.AddSeries()
	dailyPriceSeries.Values().SetReference(`'Sheet 1'!C2:C32`)
	dailyPriceSeries.CategoryAxis().SetLabelReference(`'Sheet 1'!A2:A32`)
	dailyPriceSeries.SetText("Daily sales")

	// Add Total Sales plot
	SumOfSales := lc.AddSeries()
	SumOfSales.SetText("Total Sales")
	SumOfSales.Values().SetReference(`'Sheet 1'!E2:E32`)
	SumOfSales.Values()

	ca := chart.AddCategoryAxis()
	va := chart.AddValueAxis()
	lc.AddAxis(ca)
	lc.AddAxis(va)

	ca.SetCrosses(va)
	va.SetCrosses(ca)
	// add a title and legend
	title := chart.AddTitle()
	title.SetText("Daily Sales Over Time")
	chart.AddLegend()

	// Create Bar Chart to compare total sales by each salesperson to identify who is performing well.
	chart2, anc2 := drawing.AddChart(spreadsheet.AnchorTypeTwoCell)
	anc2.SetWidthCells(8)
	anc2.MoveTo(18, 1)
	bc := chart2.AddBarChart()
	salesByPerson := bc.AddSeries()
	salesByPerson.SetText("Price")
	salesPersons := []string{}
	sales := []float64{}
	Contributions := []float64{}
	for person, totalSale := range cumulativeSumsByPerson {
		salesPersons = append(salesPersons, person)
		sales = append(sales, totalSale)
		Contribution := totalSale / totalSum
		Contributions = append(Contributions, Contribution)
	}

	salesByPerson.CategoryAxis().SetValues(salesPersons)
	salesByPerson.Values().SetValues(sales)

	ca = chart2.AddCategoryAxis()
	va = chart2.AddValueAxis()
	bc.AddAxis(ca)
	bc.AddAxis(va)

	ca.SetCrosses(va)
	va.SetCrosses(ca)
	// add a title and legend
	title2 := chart2.AddTitle()
	title2.SetText("Daily Sales Per Person")
	chart2.AddLegend()

	// Create the Pie Chart to Show percentage contribution of each salesperson to total sales.
	chart3, anc3 := drawing.AddChart(spreadsheet.AnchorTypeTwoCell)
	anc3.SetWidthCells(10)
	anc3.MoveTo(9, 23)
	pc := chart3.AddPieChart()
	totalSalesSeries := pc.AddSeries()
	totalSalesSeries.SetText("Sales Persons Contribution ")

	// Set a category axis's value with salesPersons name.
	totalSalesSeries.CategoryAxis().SetValues(salesPersons)
	totalSalesSeries.Values().SetValues(Contributions) // Set the values with each persons total sum
	totalSalesSeries.SetExplosion(3)

	// add a title and legend
	title3 := chart3.AddTitle()
	title3.SetText("Total Sales Contribution by Sales Person")
	chart3.AddLegend()
	// and finally add the chart to the sheet
	sheet.SetDrawing(drawing)

	ss.SaveToFile("charts-from-csv.xlsx")

}

func load_csv(filePath string) ([]SalesData, []string, error) {
	file, err := os.Open(filePath) // Ensure the filename matches your CSV file
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all rows from the CSV
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	// Create a slice to hold the Sale objects
	var sales []SalesData
	var header []string

	// Loop through the rows (skipping the header)
	for i, row := range rows {
		if i == 0 {
			header = append(header, row...)
			continue
		}

		dateOfSale, err := time.Parse("2006-01-02", row[0])
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing sale price: %v", err)
		}

		quantitySold, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing sale price: %v", err)
		}

		salePrice, err := strconv.ParseFloat(row[2], 64)
		if err != nil {

			return nil, nil, fmt.Errorf("error parsing sale price: %v", err)
		}

		// Create a new Sale object and populate it
		sale := SalesData{
			DateOfSale:      dateOfSale,
			QuantitySold:    quantitySold,
			SalePrice:       salePrice,
			SalespersonName: row[3],
		}

		sales = append(sales, sale)
	}
	return sales, header, nil
}
