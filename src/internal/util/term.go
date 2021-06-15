package util

import (
	"github.com/c-bata/go-prompt/term"
	"golang.org/x/sys/unix"
)

func GetWindowWidth() uint16 {
	ws, err := unix.IoctlGetWinsize(term.GetSaveTermiosFD(), unix.TIOCGWINSZ)
	if err != nil {
		return 80
	}
	return ws.Col
}

func GetWindowHeight() uint16 {
	ws, err := unix.IoctlGetWinsize(term.GetSaveTermiosFD(), unix.TIOCGWINSZ)
	if err != nil {
		return 25
	}
	return ws.Row
}
