package tuiles

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

const (
	Bold             = "\033[1m"
	Reset            = "\033[0m"
	LineStart        = "\033[G"
	HighlightedField = "\033[3m\033[7m"
)

var (
	ArchIndex   int
	GlobalIndex int = 1
	BgColor     string
	FgColor     string
	SelColor    string
)

func NewApp() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Print("\033[?25l")
}

func QuitApp() {
	fmt.Print("\033[?25h")
	fmt.Print(Reset)
	cmd1 := exec.Command("clear")
	cmd1.Stdout = os.Stdout
	cmd1.Run()
	cmd2 := exec.Command("reset")
	cmd2.Stdout = os.Stdout
	cmd2.Run()
	os.Exit(0)
}

func MoveDown() {
	if GlobalIndex >= 1 && GlobalIndex < ArchIndex {
		GlobalIndex++
	}
}

func MoveUp() {
	if GlobalIndex > 1 && GlobalIndex <= ArchIndex {
		GlobalIndex--
	}
}

func SetColors(bg, fg, sel int) {
	BgColor = "\033[48;5;" + strconv.Itoa(bg) + "m"
	FgColor = "\033[38;5;" + strconv.Itoa(fg) + "m"
	SelColor = "\033[38;5;" + strconv.Itoa(sel) + "m"
	fmt.Print(BgColor)
}
