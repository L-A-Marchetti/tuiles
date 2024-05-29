package tuiles

import "fmt"

type Checkbox struct {
	index             int
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

func SelectCheckBox(checkboxes []Checkbox) {
	for i, b := range checkboxes {
		if b.isSelected {
			checkboxes[i].isSelected = false
		}
		if b.index == GlobalIndex {
			checkboxes[i].isSelected = true
		}
	}
}

func CreateCheckBoxes(boxL, boxR, check byte, label ...string) []Checkbox {
	checkboxes := []Checkbox{}
	for _, l := range label {
		isSelected := false
		if ArchIndex == 0 {
			isSelected = true
		}
		ArchIndex++
		checkboxes = append(checkboxes, Checkbox{
			index:      ArchIndex,
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
	fmt.Println()
}
