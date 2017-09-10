package editor

import (
	"io"

	"github.com/jamesroutley/fuji/text"
	termbox "github.com/nsf/termbox-go"
)

// NormalModeCommand is a function that defines the behaviour of a normal mode
// command
type NormalModeCommand func(*Editor)

// InsertModeCommand defines the behaviour of an insert mode command
type InsertModeCommand func(*Editor)

// mode distinguishes between editor modes
type mode uint8

const (
	normalMode mode = iota
	insertMode
)

// Editor exposes the main API of the text editor
type Editor struct {
	text           *text.Text
	curX, curY     int
	normalCommands map[string]NormalModeCommand
	insertCommands map[termbox.Key]InsertModeCommand
	mode           mode
}

// New returns a new Editor
func New(r io.Reader) *Editor {
	return &Editor{
		text.New(r),
		0,
		0,
		map[string]NormalModeCommand{},
		map[termbox.Key]InsertModeCommand{},
		normalMode}
}

// Start the editor
func (e *Editor) Start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	e.Draw()

	for {
		ev := termbox.PollEvent()
		switch e.mode {
		case normalMode:
			e.handleNormalModeEvent(ev)
		case insertMode:
			e.handleInsertModeEvent(ev)
		default:
			panic("Should not reach here")
		}
		e.Draw()
	}
}

// NormalMode puts the editor into normal mode
func (e *Editor) NormalMode() {
	e.mode = normalMode
}

// InsertMode puts the editor into insert mode
func (e *Editor) InsertMode() {
	e.mode = insertMode
}

func (e *Editor) handleNormalModeEvent(ev termbox.Event) {
	if ev.Ch == 0 {
		return
	}
	command := e.normalCommands[string(ev.Ch)]
	if command == nil {
		return
	}
	command(e)
}

func (e *Editor) handleInsertModeEvent(ev termbox.Event) {
	command := e.insertCommands[ev.Key]
	if command != nil {
		command(e)
		return
	}
	if ev.Ch == 0 {
		return
	}
	e.Insert(ev.Ch)
}

// AddNormalModeCommand adds a new command to the editor
func (e *Editor) AddNormalModeCommand(name string, behaviour NormalModeCommand) {
	e.normalCommands[name] = behaviour
}

// AddInsertModeCommand adds a new insert mode command
func (e *Editor) AddInsertModeCommand(key termbox.Key, behaviour InsertModeCommand) {
	e.insertCommands[key] = behaviour
}

// Draw displays the contents of the editor
func (e *Editor) Draw() {
	// Clear screen
	// TODO: this is a naive way of doing this
	w, h := termbox.Size()
	for y := 0; y < w; y++ {
		for x := 0; x < h; x++ {
			termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
		}
	}
	for y, line := range e.text.Lines {
		for x, r := range line.String() {
			termbox.SetCell(x, y, r, termbox.ColorDefault, termbox.ColorDefault)
		}
	}
	e.displayCursor()
	if err := termbox.Flush(); err != nil {
		panic(err)
	}
}

func (e *Editor) displayCursor() {
	x := e.curX
	lineLength := e.text.LineLength(e.curY)
	if x >= lineLength-1 {
		x = lineLength - 1
	}
	// TODO: pretty hacky!
	if x < 0 {
		x = 0
	}
	termbox.SetCursor(x, e.curY)
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
	if e.curY >= e.text.Length()-1 {
		return
	}
	e.curY++
}

// CursorLeft moves the cursor left
func (e *Editor) CursorLeft() {
	if e.curX == 0 {
		return
	}
	if e.curX > e.text.LineLength(e.curY)-1 {
		e.curX = e.text.LineLength(e.curY) - 1
	}
	e.curX--
}

// CursorRight moves the cursor right
func (e *Editor) CursorRight() {
	if e.curX >= e.text.LineLength(e.curY)-1 {
		return
	}
	e.curX++
}

// Insert inserts rune r at the cursor position
func (e *Editor) Insert(r rune) {
	e.text = e.text.Insert(e.curX, e.curY, r)
	e.CursorRight()
}

// Delete deletes the rune under the cursor
func (e *Editor) Delete() {
	e.text = e.text.Delete(e.curX, e.curY)
}
