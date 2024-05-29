package tuiles

import (
	"fmt"
)

type Button struct {
	index      int
	label      string
	isSelected bool
}

func CreateButtons(label ...string) []Button {
	buttons := []Button{}
	for _, l := range label {
		isSelected := false
		if ArchIndex == 0 {
			isSelected = true
		}
		ArchIndex++
		buttons = append(buttons, Button{
			index:      ArchIndex,
			label:      l,
			isSelected: isSelected,
		})
	}
	return buttons
}

func PrintButtons(buttons []Button) {
	for _, b := range buttons {
		if b.isSelected {
			fmt.Printf("%s%s%s%s %s %s\n\n", LineStart, SelColor, Bold, ButtonField, b.label, Reset)
			fmt.Print(BgColor)
		} else {
			fmt.Printf("%s%s%s %s %s\n\n", LineStart, FgColor, ButtonField, b.label, Reset)
			fmt.Print(BgColor)
		}
	}
	fmt.Println()
}

func SelectButtons(buttons []Button) {
	for i, b := range buttons {
		if b.isSelected {
			buttons[i].isSelected = false
		}
		if b.index == GlobalIndex {
			buttons[i].isSelected = true
		}
	}
}

func ClickButton(buttons Button, f func()) {
	if buttons.isSelected {
		f()
	}
}
