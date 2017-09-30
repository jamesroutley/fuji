package pane

import (
	"os"

	"github.com/gdamore/tcell"
	"github.com/jamesroutley/fuji/area"
	"github.com/jamesroutley/fuji/editarea"
	"github.com/jamesroutley/fuji/statusbar"
)

// EditPane is a pane which contains editable text
type EditPane struct {
	screen    tcell.Screen
	area      area.Area
	statusbar *statusbar.StatusBar
	editarea  *editarea.EditArea
}

// NewEditPane initialises and returns a new EditPane
func NewEditPane(filename string, screen tcell.Screen, area area.Area) *EditPane {
	file, err := os.Open(filename)
	if err != nil {
		panic("cannot open file: " + filename)
	}
	defer file.Close()
	editarea := editarea.New(screen, filename, file)
	statusbar := statusbar.New(screen)
	return &EditPane{
		screen:    screen,
		area:      area,
		editarea:  editarea,
		statusbar: statusbar,
	}
}

// Draw draws the contents of the edit pane
func (ep *EditPane) Draw() {
	ep.editarea.Draw(area.Area{
		Start: ep.area.Start,
		End:   area.Point{X: ep.area.End.X, Y: ep.area.End.Y},
	})
	ep.statusbar.Draw(
		ep.editarea,
		area.Area{
			Start: area.Point{X: ep.area.Start.X, Y: ep.area.End.Y},
			End:   ep.area.End,
		},
	)
}

// HandleEvent handles the tcell event ev
func (ep *EditPane) HandleEvent(ev *tcell.EventKey) {
	ep.editarea.HandleEvent(ev)
}
