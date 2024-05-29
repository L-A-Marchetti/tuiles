package tuiles

import (
	"fmt"
	"strings"
)

type Textfield struct {
	index       int
	Label, Text string
	isSelected  bool
}

func CreateTextFields(label ...string) []Textfield {
	textFields := []Textfield{}
	for _, l := range label {
		isSelected := false
		if ArchIndex == 0 {
			isSelected = true
		}
		ArchIndex++
		textFields = append(textFields, Textfield{
			index:      ArchIndex,
			Label:      l,
			isSelected: isSelected,
			Text:       "Write Here...",
		})
	}
	return textFields
}

func PrintTextFields(textFields []Textfield) {
	for _, b := range textFields {
		if b.isSelected {
			fmt.Printf("%s%s%s%s: %s%s%s\n", LineStart, SelColor, Bold, b.Label, HighlightedField, b.Text, Reset)
			fmt.Print(BgColor)
		} else {
			fmt.Printf("%s%s%s: %s\n", LineStart, FgColor, b.Label, b.Text)
			fmt.Print(BgColor)
		}
	}
	fmt.Println()
}

func SelectTextField(textFields []Textfield) {
	for i, b := range textFields {
		if b.isSelected {
			textFields[i].isSelected = false
		}
		if b.index == GlobalIndex {
			textFields[i].isSelected = true
		}
	}
}

func TextFieldErase(textFields []Textfield) {
	for i, b := range textFields {
		if b.isSelected && len(b.Text) > 0 {
			textFields[i].Text = strings.TrimSpace(textFields[i].Text[:len(textFields[i].Text)-1]) // delete last character
			break
		}
	}
}

func TextFieldWrite(textFields []Textfield, c string) {
	for i, b := range textFields {
		if b.Text == "Write Here..." && b.isSelected {
			textFields[i].Text = ""
		}
		if b.isSelected {
			textFields[i].Text = textFields[i].Text + c
			break
		}
	}
}

func ResetTextFields(textFields []Textfield) {
	for r := range textFields {
		textFields[r].Text = "Write Here..."
	}
}
