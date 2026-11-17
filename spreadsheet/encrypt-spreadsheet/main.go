package main

import (
	"fmt"
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
	// EncryptionAES192 and EncryptionAES128 select weaker keys if needed.
	opts := &crypto.SaveOptions{
		Password:   "unioffice",
		Encryption: crypto.EncryptionAES256,
	}
	if err := wb.SaveToFileWithOptions("encrypted.xlsx", opts); err != nil {
		fmt.Println(err)
		return
	}

	got, err := spreadsheet.OpenWithOptions("encrypted.xlsx", &crypto.OpenOptions{Password: "unioffice"})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer got.Close()
	fmt.Println("opened encrypted workbook")
}
