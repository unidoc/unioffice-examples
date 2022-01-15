/*
 * This example showcases getting values from ActiveX forms and changing them.
 */

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
)

func init() {
	// This example requires both for unioffice and unipdf.
	if err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`)); err != nil {
		fmt.Printf("ERROR: Failed to set metered key: %v\n", err)
		fmt.Printf("Make sure to get a valid key from https://cloud.unidoc.io\n")
		fmt.Printf("If you don't have one - Grab one in the Free Tier at https://cloud.unidoc.io\n")
		panic(err)
	}
}

func main() {
	doc, err := document.Open("activex_filled.docm")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc.Close()
	for i, p := range doc.Paragraphs() {
		for _, r := range p.Runs() {
			ctrl := r.Control()
			if ctrl != nil {
				if ctrl.Choice != nil {
					if checkBox := ctrl.Choice.CheckBox; checkBox != nil {
						fmt.Println("found checkbox:", checkBox.GetValue(), checkBox.GetCaption())
						checkBox.SetValue(true)
						checkBox.SetCaption(fmt.Sprintf("CheckBox caption %d", i))
					} else if textBox := ctrl.Choice.TextBox; textBox != nil {
						fmt.Println("found textbox:", textBox.GetValue(), textBox.GetCaption())
						textBox.SetValue(fmt.Sprintf("New textbox value %d", i))
						textBox.SetCaption(fmt.Sprintf("TextBox caption %d", i))
					} else if comboBox := ctrl.Choice.ComboBox; comboBox != nil {
						fmt.Println("found combo box:", comboBox.GetValue())
						comboBox.SetValue(fmt.Sprintf("New combobox value %d", i))
					} else if optionButton := ctrl.Choice.OptionButton; optionButton != nil {
						fmt.Println("found option button:", optionButton.GetValue(), optionButton.GetCaption())
						optionButton.SetValue(!optionButton.GetValue())
						optionButton.SetCaption(fmt.Sprintf("Option button %d", i))
					} else if toggleButton := ctrl.Choice.ToggleButton; toggleButton != nil {
						fmt.Println("found toggle button:", toggleButton.GetValue(), toggleButton.GetCaption())
						toggleButton.SetValue(true)
						toggleButton.SetCaption(fmt.Sprintf("Toggle button %d", i))
					} else if label := ctrl.Choice.Label; label != nil {
						fmt.Println("found label:", label.GetCaption())
						label.SetCaption(fmt.Sprintf("New label %d", i))
						label.SetForeColor(uint32(0x02ff0000))
						label.SetBackColor(uint32(0x020044ff))
					} else if spinButton := ctrl.Choice.SpinButton; spinButton != nil {
						fmt.Println("found spin button:", spinButton.GetMin(), spinButton.GetMax(), spinButton.GetPosition(), spinButton.GetWidth(), spinButton.GetHeight())
						spinButton.SetPosition(1 - spinButton.GetPosition())
						spinButton.SetForeColor(uint32(0x020044ff))
						spinButton.SetBackColor(uint32(0x02ff0000))
					} else if commandButton := ctrl.Choice.CommandButton; commandButton != nil {
						fmt.Println("found command button:", commandButton.GetCaption())
						commandButton.SetCaption(fmt.Sprintf("Command button %d", i))
						commandButton.SetForeColor(uint32(0x02ffffff))
						commandButton.SetBackColor(uint32(0x0200ff00))
					} else if scrollBar := ctrl.Choice.ScrollBar; scrollBar != nil {
						fmt.Println("found scroll bar:", scrollBar.GetMin(), scrollBar.GetMax(), scrollBar.GetPosition(), scrollBar.GetWidth(), scrollBar.GetHeight())
						scrollBar.SetMax(100)
						scrollBar.SetPosition(20)
					}
				}
			}
		}
	}
	doc.SaveToFile("new_activex_filled.docm")
}
