package tuiles

import "fmt"

type Checkbox struct {
	index      int
	label      string
	isChecked  bool
	isSelected bool
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

func CreateCheckBoxes(label ...string) []Checkbox {
	checkboxes := []Checkbox{}
	for _, l := range label {
		isSelected := false
		if ArchIndex == 0 {
			isSelected = true
		}
		ArchIndex++
		checkboxes = append(checkboxes, Checkbox{
			index:      ArchIndex,
			label:      l,
			isChecked:  false,
			isSelected: isSelected,
		})
	}
	return checkboxes
}

func PrintCheckBoxes(checkBoxes []Checkbox) {
	for _, b := range checkBoxes {
		if b.isChecked {
			if b.isSelected {
				fmt.Printf("%s%s%s%c%c%c %s%s\n", LineStart, SelColor, Bold, '[', 'X', ']', b.label, Reset)
				fmt.Print(BgColor)
			} else {
				fmt.Printf("%s%s%c%c%c %s\n", LineStart, FgColor, '[', 'X', ']', b.label)
				fmt.Print(BgColor)
			}
		} else {
			if b.isSelected {
				fmt.Printf("%s%s%s%c %c %s%s\n", LineStart, SelColor, Bold, '[', ']', b.label, Reset)
				fmt.Print(BgColor)
			} else {
				fmt.Printf("%s%s%c %c %s\n", LineStart, FgColor, '[', ']', b.label)
				fmt.Print(BgColor)
			}
		}
	}
	fmt.Println()
}
