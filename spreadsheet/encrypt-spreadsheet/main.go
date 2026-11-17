/*
 * Copyright 2026 FoxyUtils ehf. All rights reserved.
 *
 * This example demonstrates saving an XLSX encrypted with a password and
 * opening it again with the password.
 */

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/crypto"
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
	wb := spreadsheet.New()
	defer wb.Close()
	sheet := wb.AddSheet()
	sheet.Cell("A1").SetString("This workbook is encrypted with a password.")

	// The zero value of Encryption (EncryptionDefault) is agile AES-256;
	// EncryptionAES192 and EncryptionAES128 select a smaller key if required.
	opts := &crypto.SaveOptions{
		Password:   "unioffice",
		Encryption: crypto.EncryptionAES256,
	}
	if err := wb.SaveToFileWithOptions("encrypted.xlsx", opts); err != nil {
		log.Fatal(err)
	}

	got, err := spreadsheet.OpenWithOptions("encrypted.xlsx", &crypto.OpenOptions{Password: "unioffice"})
	if err != nil {
		log.Fatal(err)
	}
	defer got.Close()
	fmt.Println("A1:", got.Sheets()[0].Cell("A1").GetString())
}
