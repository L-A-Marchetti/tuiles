package tuiles

import "fmt"

type Checkbox struct {
	boxL, boxR, check byte
	label             string
	isChecked         bool
	isSelected        bool
}

func DrawCheckBox(checkboxes []Checkbox) {
	for i, b := range checkboxes {
		if b.isSelected {
			if !b.isChecked {
				checkboxes[i].isChecked = true
			} else {
				checkboxes[i].isChecked = false
			}
			break
		}
	}
}

func SelectCheckBox(checkboxes []Checkbox, vector int) {
	selector := 0
	for i, b := range checkboxes {
		if b.isSelected {
			checkboxes[i].isSelected = false
			selector = i + vector
			break
		}
	}
	for i := range checkboxes {
		if i == selector {
			checkboxes[i].isSelected = true
		}
	}
}

func CreateCheckBoxes(boxL, boxR, check byte, label ...string) []Checkbox {
	checkboxes := []Checkbox{}
	for i, l := range label {
		isSelected := false
		if i == 0 {
			isSelected = true
		}
		checkboxes = append(checkboxes, Checkbox{
			boxL:       boxL,
			boxR:       boxR,
			check:      check,
			label:      l,
			isChecked:  false,
			isSelected: isSelected,
		})
	}
	return checkboxes
}

func PrintCheckBoxes(checkBoxes []Checkbox) {
	highlight := "\033[1m"
	reset := "\033[0m"
	lineStart := "\033[G"
	for _, b := range checkBoxes {
		if b.isChecked {
			if b.isSelected {
				fmt.Printf("%s%s%c%c%c %s%s\n", lineStart, highlight, b.boxL, b.check, b.boxR, b.label, reset)
			} else {
				fmt.Printf("%s%c%c%c %s\n", lineStart, b.boxL, b.check, b.boxR, b.label)
			}
		} else {
			if b.isSelected {
				fmt.Printf("%s%s%c %c %s%s\n", lineStart, highlight, b.boxL, b.boxR, b.label, reset)
			} else {
				fmt.Printf("%s%c %c %s\n", lineStart, b.boxL, b.boxR, b.label)
			}
		}
	}
}
