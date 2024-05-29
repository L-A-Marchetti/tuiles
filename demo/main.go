package main

import (
	"fmt"
	"time"

	"github.com/L-A-Marchetti/tuiles"
)

func main() {
	tuiles.NewApp()
	tuiles.SetColors(233, 89, 24)
	key := tuiles.KeyListener()
	time.Sleep(10 * time.Millisecond)
	textFields := tuiles.CreateTextFields("First field", "Second field", "and anotherone")
	checkBoxes := tuiles.CreateCheckBoxes("This is the first box", "then the second one", "and another one")
	buttons := tuiles.CreateButtons("Reset text fields", "Reset check boxes", "Ok", "Quit")
	tuiles.PrintTextFields(textFields)
	tuiles.PrintCheckBoxes(checkBoxes)
	tuiles.PrintButtons(buttons)
	for c := range key {
		tuiles.ClearConsole()
		switch c {
		case "down":
			tuiles.MoveDown()
			tuiles.SelectTextField(textFields)
			tuiles.SelectCheckBox(checkBoxes)
			tuiles.SelectButtons(buttons)
		case "up":
			tuiles.MoveUp()
			tuiles.SelectTextField(textFields)
			tuiles.SelectCheckBox(checkBoxes)
			tuiles.SelectButtons(buttons)
		case "enter":
			tuiles.DrawCheckBox(checkBoxes)
			tuiles.ClickButton(buttons[0], func() { tuiles.ResetTextFields(textFields) })
			tuiles.ClickButton(buttons[1], func() { tuiles.ResetCheckBoxes(checkBoxes) })
			tuiles.ClickButton(buttons[2], func() { PrintTheResult(textFields, checkBoxes) })
			tuiles.ClickButton(buttons[3], func() { tuiles.QuitApp() })
		case "escape":
			tuiles.QuitApp()
		case "backspace":
			tuiles.TextFieldErase(textFields)
		default:
			tuiles.TextFieldWrite(textFields, c)
		}
		tuiles.PrintTextFields(textFields)
		tuiles.PrintCheckBoxes(checkBoxes)
		tuiles.PrintButtons(buttons)
	}
}

func PrintTheResult(textFields []tuiles.Textfield, checkBoxes []tuiles.Checkbox) {
	fmt.Println("Results :")
	for _, tF := range textFields {
		fmt.Printf("%s: [%s]\n", tF.Label, tF.Text)
	}
	for _, cB := range checkBoxes {
		fmt.Printf("%s: [%v]\n", cB.Label, cB.IsChecked)
	}
}
