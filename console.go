package tuiles

import (
	"os"
	"os/exec"
)

func ClearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ResetConsole() {
	cmd := exec.Command("reset")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
