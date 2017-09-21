package statusbar

import (
	"strings"

	"github.com/gdamore/tcell"
)

// StatusBar implements the status bar.
// The status bar is intended for global information.
type StatusBar struct {
	screen tcell.Screen
}

// New initialises and returns a new StatusBar
func New(screen tcell.Screen) *StatusBar {
	return &StatusBar{
		screen: screen,
	}
}

// Status is a function which, when called, returns some status string
type Status func() string

var statusItems []Status

// Draw draws the status bar
func (s *StatusBar) Draw() {
	xmax, y := s.screen.Size()
	// statuses := getStatuses(" / ")

	style := tcell.StyleDefault.Background(tcell.ColorDarkGray)

	for x := 0; x < xmax; x++ {
		s.screen.SetContent(x, y-1, ' ', nil, style)
	}

}

// AddStatus adds a status to the bar. Statuses are displayed in the order
// they were added
func AddStatus(s Status) {
	statusItems = append(statusItems, s)
}

func getStatuses(sep string) string {
	statuses := make([]string, len(statusItems))
	for i, item := range statusItems {
		statuses[i] = item()
	}
	return strings.Join(statuses, sep)
}
