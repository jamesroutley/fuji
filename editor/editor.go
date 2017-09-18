package editor

import (
	"os"

	"github.com/jamesroutley/fuji/editarea"
	termbox "github.com/nsf/termbox-go"
)

// Editor implements the main editor
type Editor struct {
	editarea editarea.EditArea
}

// Start starts the editor
func (e *Editor) Start(filename string) {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	file, err := os.Open(filename)
	if err != nil {
		panic("cannot open file: " + filename)
	}
	defer file.Close()
	editarea := editarea.New(filename, file)

	editarea.Draw()

	for {
		ev := termbox.PollEvent()
		editarea.HandleEvent(ev)
		editarea.Draw()
	}
}
