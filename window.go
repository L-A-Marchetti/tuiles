package tuiles

import "fmt"

func ResizeWindow(w, h int) {
	fmt.Printf("\033[8;%v;%vt", h, w)
}
