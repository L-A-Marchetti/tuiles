package tuiles

import "fmt"

type Checkbox struct {
	index      int
	Label      string
	IsChecked  bool
	isSelected bool
}

func DrawCheckBox(checkboxes []Checkbox) {
	for i, b := range checkboxes {
		if b.isSelected {
			if !b.IsChecked {
				checkboxes[i].IsChecked = true
			} else {
				checkboxes[i].IsChecked = false
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
			Label:      l,
			IsChecked:  false,
			isSelected: isSelected,
		})
	}
	return checkboxes
}

func PrintCheckBoxes(checkBoxes []Checkbox) {
	for _, b := range checkBoxes {
		if b.IsChecked {
			if b.isSelected {
				fmt.Printf("%s%s%s%c%c%c %s%s\n", LineStart, SelColor, Bold, '[', 'X', ']', b.Label, Reset)
				fmt.Print(BgColor)
			} else {
				fmt.Printf("%s%s%c%c%c %s\n", LineStart, FgColor, '[', 'X', ']', b.Label)
				fmt.Print(BgColor)
			}
		} else {
			if b.isSelected {
				fmt.Printf("%s%s%s%c %c %s%s\n", LineStart, SelColor, Bold, '[', ']', b.Label, Reset)
				fmt.Print(BgColor)
			} else {
				fmt.Printf("%s%s%c %c %s\n", LineStart, FgColor, '[', ']', b.Label)
				fmt.Print(BgColor)
			}
		}
	}
	fmt.Println()
}

func ResetCheckBoxes(checkBoxes []Checkbox) {
	for r := range checkBoxes {
		checkBoxes[r].IsChecked = false
	}
}
