package tuiles

import (
	"fmt"
	"strings"
)

type Textfield struct {
	index       int
	label, text string
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
			label:      l,
			isSelected: isSelected,
			text:       "Write Here...",
		})
	}
	return textFields
}

func PrintTextFields(textFields []Textfield) {
	highlight := "\033[1m"
	reset := "\033[0m"
	lineStart := "\033[G"
	highlightedField := "\033[3m\033[7m"
	for _, b := range textFields {
		if b.isSelected {
			fmt.Printf("%s%s%s%s: %s%s%s\n", lineStart, highlight, b.label, reset, highlightedField, b.text, reset)
		} else {
			fmt.Printf("%s%s: %s\n", lineStart, b.label, b.text)
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
		if b.isSelected && len(b.text) > 0 {
			textFields[i].text = strings.TrimSpace(textFields[i].text[:len(textFields[i].text)-1]) // delete last character
			break
		}
	}
}

func TextFieldWrite(textFields []Textfield, c string) {
	for i, b := range textFields {
		if b.text == "Write Here..." && b.isSelected {
			textFields[i].text = ""
		}
		if b.isSelected {
			textFields[i].text = textFields[i].text + c
			break
		}
	}
}
