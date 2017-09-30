package statusbar

import (
	"strings"

	"github.com/gdamore/tcell"
	"github.com/jamesroutley/fuji/area"
	"github.com/jamesroutley/fuji/editarea"
)

var statuses []func(*editarea.EditArea) string

// AddStatus adds a new status function
func AddStatus(status func(*editarea.EditArea) string) {
	statuses = append(statuses, status)
}

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
func (s *StatusBar) Draw(e *editarea.EditArea, area area.Area) {
	statuses := getStatuses(e)
	content := strings.Join(statuses, " / ")

	style := tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorBlack)

	for x := area.Start.X; x < area.End.X; x++ {
		s.screen.SetContent(x, area.Start.Y, ' ', nil, style)
	}

	offset := 1
	for x, r := range content {
		s.screen.SetContent(x+offset, area.Start.Y, r, nil, style)
	}
}

func getStatuses(e *editarea.EditArea) []string {
	content := make([]string, len(statuses))
	for i, status := range statuses {
		content[i] = status(e)
	}
	return content
}
