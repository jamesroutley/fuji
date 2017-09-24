package statusbar

import (
	"strings"

	"github.com/gdamore/tcell"
)

// StatusBar implements the status bar.
// The status bar is intended for information relevant to an edit area.
type StatusBar struct {
	screen tcell.Screen
}

// New initialises and returns a new StatusBar
func New(screen tcell.Screen) *StatusBar {
	return &StatusBar{screen: screen}
}

// Draw draws the status bar
func (s *StatusBar) Draw(statuses []string) {
	offset := 1
	xmax, y := s.screen.Size()
	content := strings.Join(statuses, " / ")

	style := tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorBlack)

	for x := 0; x < xmax; x++ {
		s.screen.SetContent(x, y-2, ' ', nil, style)
	}

	for x, r := range content {
		s.screen.SetContent(x+offset, y-2, r, nil, style)
	}
}
