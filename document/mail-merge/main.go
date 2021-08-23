package main

import (
	"log"
	"os"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
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
	d, err := document.Open("mm.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer d.Close()
	for _, v := range d.MergeFields() {
		log.Println("replacing", v)
	}
	rep := map[string]string{}
	rep["Title"] = "mr."      // has a \* Upper attribute on the field
	rep["FirstName"] = "JOHN" // has a \* Lower attribute on the field
	d.MailMerge(rep)
	d.SaveToFile("merged.docx")
}
