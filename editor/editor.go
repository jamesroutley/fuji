package editor

import (
	"io"

	"github.com/jamesroutley/fuji/text"
	termbox "github.com/nsf/termbox-go"
)

// Editor exposes the main API of the text editor
type Editor struct {
	text       *text.Text
	curX, curY int
}

// New returns a new Editor
func New(r io.Reader) *Editor {
	return &Editor{text.New(r), 0, 0}
}

// Start the editor
func (e *Editor) Start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	e.Draw()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				panic("esc")
			}
		}
		e.Draw()
	}
}

// Draw displays the contents of the editor
func (e *Editor) Draw() {
	for y, line := range e.text.Lines {
		for x, r := range line.String() {
			termbox.SetCell(x, y, r, termbox.ColorDefault, termbox.ColorDefault)
		}
	}
	termbox.SetCursor(e.curX, e.curY)
	termbox.Flush()
}

// CursorUp moves the cursor up
func (e *Editor) CursorUp() {
	if e.curY == 0 {
		return
	}
	e.curY--
}

// CursorDown moves the cursor down
func (e *Editor) CursorDown() {
	e.curY++
}

// CursorLeft moves the cursor left
func (e *Editor) CursorLeft() {
	if e.curX == 0 {
		return
	}
	e.curX--
}

// CursorRight moves the cursor right
func (e *Editor) CursorRight() {
	e.curX++
}
