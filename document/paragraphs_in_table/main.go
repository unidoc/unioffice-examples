package main

import (
	"os"

	"github.com/unidoc/unioffice/v2/color"
	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/document"
	"github.com/unidoc/unioffice/v2/measurement"
	"github.com/unidoc/unioffice/v2/schema/soo/wml"
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
	doc := document.New()

	paraBeforeTable := doc.AddParagraph()
	paraBeforeTable.AddRun().AddText("before table")

	table := doc.InsertTableAfter(paraBeforeTable)
	table.Properties().Borders().SetAll(wml.ST_BorderBasicBlackDots, color.AliceBlue, measurement.Point*2)
	tablePara1 := table.AddRow().AddCell().AddParagraph()
	tablePara1.AddRun().AddText("table paragraph 1")

	paraAfterTable := doc.AddParagraph()
	paraAfterTable.AddRun().AddText("after table")

	tablePara2 := doc.InsertParagraphAfter(tablePara1)
	tablePara2.AddRun().AddText("table paragraph after table paragraph 1")

	tablePara3 := doc.InsertParagraphBefore(tablePara1)
	tablePara3.AddRun().AddText("table paragraph before table paragraph 1")

	tableInTable := doc.InsertTableAfter(tablePara3)
	tableInTable.Properties().Borders().SetAll(wml.ST_BorderBasicBlackDots, color.DarkGreen, measurement.Point*2)
	tableInTablePara := tableInTable.AddRow().AddCell().AddParagraph()
	tableInTablePara.AddRun().AddText("table in table paragraph 1")

	tableInTableInTable := doc.InsertTableBefore(tableInTablePara)
	tableInTableInTable.Properties().Borders().SetAll(wml.ST_BorderBasicBlackDots, color.OrangeRed, measurement.Point*2)
	tableInTableInTablePara := tableInTableInTable.AddRow().AddCell().AddParagraph()
	tableInTableInTablePara.AddRun().AddText("table in table in table paragraph 1")

	doc.SaveToFile("out.docx")
}
