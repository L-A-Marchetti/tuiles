package tuiles

import (
	"unicode"

	"github.com/nsf/termbox-go"
)

func KeyListener() chan string {
	key := make(chan string)
	go func() {
		err := termbox.Init()
		if err != nil {
			panic(err)
		}
		defer termbox.Close()
		for {
			ev := termbox.PollEvent()
			if ev.Type == termbox.EventKey {
				key <- keyEventToString(ev.Key)
				if ev.Ch != 0 && unicode.IsPrint(rune(ev.Ch)) {
					key <- string(ev.Ch)
				}
			}
		}
	}()
	return key
}

func keyEventToString(ev termbox.Key) string {
	switch ev {
	case termbox.KeyArrowUp:
		return "up"
	case termbox.KeyArrowDown:
		return "down"
	case termbox.KeyArrowLeft:
		return "left"
	case termbox.KeyArrowRight:
		return "right"
	case termbox.KeyBackspace2:
		return "backspace"
	case termbox.KeySpace:
		return " "
	case termbox.KeyTab:
		return "tab"
	case termbox.KeyEnter:
		return "enter"
	case termbox.KeyEsc:
		return "escape"
	}
	return ""
}
