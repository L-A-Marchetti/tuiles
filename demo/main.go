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
			tuiles.ClickButton(buttons)
		case "escape":
			tuiles.QuitApp()
		case "backspace":
			tuiles.TextFieldErase(textFields)
		default:
			tuiles.TextFieldWrite(textFields, c)
		}
		if buttons[0].IsClicked {
			tuiles.ResetTextFields(textFields)
		}
		if buttons[1].IsClicked {
			tuiles.ResetCheckBoxes(checkBoxes)
		}
		if buttons[2].IsClicked {
			fmt.Println("Results :")
			for _, tF := range textFields {
				fmt.Printf("%s: [%s]\n", tF.Label, tF.Text)
			}
			for _, cB := range checkBoxes {
				fmt.Printf("%s: [%v]\n", cB.Label, cB.IsChecked)
			}
			return
		}
		if buttons[3].IsClicked {
			tuiles.QuitApp()
		}
		tuiles.PrintTextFields(textFields)
		tuiles.PrintCheckBoxes(checkBoxes)
		tuiles.PrintButtons(buttons)
		tuiles.ResetButton(buttons)
	}
}
