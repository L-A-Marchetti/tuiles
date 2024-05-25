package tuiles

import (
	"os"

	"golang.org/x/term"
)

func KeyListener() chan string {
	key := make(chan string)
	go func() {
		fd := int(os.Stdin.Fd())
		term.MakeRaw(fd)
		for {
			b := make([]byte, 1)
			os.Stdin.Read(b)
			key <- string(b)
		}
	}()
	return key
}
