package main

import (
	"time"

	"github.com/L-A-Marchetti/tuiles"
)

func main() {
	tuiles.NewApp()
	tuiles.SetColors(233, 89, 57)
	key := tuiles.KeyListener()
	time.Sleep(10 * time.Millisecond)
	textFields := tuiles.CreateTextFields("First field", "Second field", "and anotherone")
	tuiles.PrintTextFields(textFields)
	listOfCheckBoxes := tuiles.CreateCheckBoxes("This is the first box", "then the second one", "and another one")
	tuiles.PrintCheckBoxes(listOfCheckBoxes)
	for c := range key {
		tuiles.ClearConsole()
		switch c {
		case "down":
			tuiles.MoveDown()
			tuiles.SelectTextField(textFields)
			tuiles.SelectCheckBox(listOfCheckBoxes)
		case "up":
			tuiles.MoveUp()
			tuiles.SelectTextField(textFields)
			tuiles.SelectCheckBox(listOfCheckBoxes)
		case "enter":
			tuiles.DrawCheckBox(listOfCheckBoxes)
		case "escape":
			tuiles.QuitApp()
		case "backspace":
			tuiles.TextFieldErase(textFields)
		default:
			tuiles.TextFieldWrite(textFields, c)
		}
		tuiles.PrintTextFields(textFields)
		tuiles.PrintCheckBoxes(listOfCheckBoxes)
	}
}
