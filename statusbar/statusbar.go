package statusbar

import (
	"strings"

	termbox "github.com/nsf/termbox-go"
)

// StatusBar implements the status bar.
// The status bar is intended for global information.
type StatusBar struct{}

// Status is a function which, when called, returns some status string
type Status func() string

var statusItems []Status

// Draw draws the status bar
func (s *StatusBar) Draw() {
	xmax, y := termbox.Size()
	// statuses := getStatuses(" / ")

	for x := 0; x < xmax; x++ {
		termbox.SetCell(x, y-1, ' ', termbox.ColorDefault, termbox.ColorCyan)
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
