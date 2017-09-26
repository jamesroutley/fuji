package editor

import (
	"os"

	"github.com/gdamore/tcell"
	"github.com/jamesroutley/fuji/editarea"
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

	file, err := os.Open(filename)
	if err != nil {
		panic("cannot open file: " + filename)
	}
	defer file.Close()
	editarea := editarea.New(screen, filename, file)

	editarea.Draw()

	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			editarea.HandleEvent(ev)
		default:
			// do something
		}
		screen.Clear()
		editarea.Draw()
		screen.Show()
	}
}
