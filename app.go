package tuiles

import (
	"fmt"
	"os"
	"os/exec"
)

func NewApp() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Print("\033[?25l")
}

func QuitApp() {
	fmt.Print("\033[?25h")
	cmd1 := exec.Command("clear")
	cmd1.Stdout = os.Stdout
	cmd1.Run()
	cmd2 := exec.Command("reset")
	cmd2.Stdout = os.Stdout
	cmd2.Run()
	os.Exit(0)
}
