package tuiles

import "fmt"

type Textfield struct {
	label, text string
	isSelected  bool
}

func CreateTextFields(label ...string) []Textfield {
	textFields := []Textfield{}
	for i, l := range label {
		isSelected := false
		if i == 0 {
			isSelected = true
		}
		textFields = append(textFields, Textfield{
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
}

func SelectTextField(textFields []Textfield, vector int) {
	selector := 0
	for i, b := range textFields {
		if b.isSelected {
			textFields[i].isSelected = false
			selector = i + vector
			break
		}
	}
	for i := range textFields {
		if i == selector {
			textFields[i].isSelected = true
		}
	}
}
