// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

// This example demonstrates parsing range references with sheet names containing exclamation marks.

import (
	"fmt"
	"os"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/spreadsheet/reference"
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
	parse("A1:A3")
	parse("Sheet1!A1:A3")
	parse("Sheet1!!A1:A3")
	parse("Shee!t1!A1:A3")
	parse("!Sheet1!A1:A3")
}

func parse(s string) {
	from, to, err := reference.ParseRangeReference(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(from.RowIdx, from.Column, from.SheetName)
	fmt.Println(to.RowIdx, to.Column, to.SheetName)
}
