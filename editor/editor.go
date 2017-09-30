package editor

import (
	"github.com/gdamore/tcell"
	"github.com/jamesroutley/fuji/area"
	"github.com/jamesroutley/fuji/pane"
)

// Editor implements the main editor
type Editor struct{}

// Start starts the editor
func (e *Editor) Start(filename string) {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}
	defer screen.Fini()

	xMax, yMax := screen.Size()
	// TODO: yMax - 1 because of tmux - find some way to fix this
	area := area.Area{
		area.Point{X: 0, Y: 0}, area.Point{X: xMax, Y: yMax - 1},
	}
	editpane := pane.NewEditPane(filename, screen, area)

	editpane.Draw()

	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			editpane.HandleEvent(ev)
		default:
			// do something
		}
		screen.Clear()
		editpane.Draw()
		screen.Show()
	}
}
