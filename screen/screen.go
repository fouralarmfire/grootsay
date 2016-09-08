package screen

import (
	"fmt"

	"golang.org/x/crypto/ssh/terminal"
)

type Screen struct {
	width  int
	height int
}

func NewScreen() *Screen {
	w, h, err := terminal.GetSize(0)
	if err != nil {
		fmt.Errorf("failed to get terminal size: %s", err)
		return &Screen{}
	}
	return &Screen{
		width:  w,
		height: h,
	}
}

func (s *Screen) Width() int {
	return s.width
}
