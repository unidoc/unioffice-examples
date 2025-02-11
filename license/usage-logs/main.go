package main

import (
	"log"
	"os"

	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/common/logger"
	"github.com/unidoc/unioffice/v2/document"
)

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		panic(err)
	}

	// Set the log level to info or higher
	logger.SetLogger(logger.NewConsoleLogger(logger.LogLevelInfo))

	// Enable the verbose mode usage logging to log the usage info.
	license.SetMeteredKeyUsageLogVerboseMode(true)
}

func main() {
	additional_items := []string{"Yogurt", "Cheese", "Cereal", "Pasta", "Olive Oil"}
	doc, err := document.Open("grocery_list.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc.Close()

	var itemStyle string

	// get the style of the list item
	paras := doc.Paragraphs()
	if len(paras) > 1 {
		p := paras[1]
		itemStyle = p.Style()
	}

	for _, item := range additional_items {
		newP := doc.AddParagraph()
		r := newP.AddRun()
		r.AddText(item)
		newP.SetStyle(itemStyle)
	}

	doc.SaveToFile("updated_list.docx")
}
