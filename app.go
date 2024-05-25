package tuiles

import "os"

func NewApp() {
	ClearConsole()
	HideCursor()
}

func QuitApp() {
	ShowCursor()
	ClearConsole()
	ResetConsole()
	os.Exit(0)
}
