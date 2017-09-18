package editor

import (
	"os"

	"github.com/jamesroutley/fuji/text"
	termbox "github.com/nsf/termbox-go"
)

// NormalModeCommand is a function that defines the behaviour of a normal mode
// command
type NormalModeCommand func(*Editor)

// InsertModeCommand defines the behaviour of an insert mode command
type InsertModeCommand func(*Editor)

// Mode distinguishes between editor modes
type Mode uint8

const (
	// ModeNormal indicates the editor is in normal mode
	ModeNormal Mode = iota
	// ModeInsert indicates the editor is in insert mode
	ModeInsert
)

// Editor exposes the main API of the text editor
type Editor struct {
	Filename       string
	text           *text.Text
	curX, curY     int
	normalCommands map[string]NormalModeCommand
	insertCommands map[termbox.Key]InsertModeCommand
	Mode           Mode
}

// New returns a new Editor
func New(filename string) *Editor {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return &Editor{
		filename,
		text.New(f),
		0,
		0,
		map[string]NormalModeCommand{},
		map[termbox.Key]InsertModeCommand{},
		ModeNormal}
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
		switch e.Mode {
		case ModeNormal:
			e.handleNormalModeEvent(ev)
		case ModeInsert:
			e.handleInsertModeEvent(ev)
		default:
			panic("Should not reach here")
		}
		e.Draw()
	}
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
	for y := 0; y < e.text.Length(); y++ {
		line := e.text.Line(y)
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

// CursorAtLineStart returns whether the cursor is at the beginning of a line
func (e *Editor) CursorAtLineStart() bool {
	return e.curX == 0
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
	e.text = e.text.Insert(e.curY, e.curX, r)
	e.CursorRight()
}

// Delete deletes the rune under the cursor
func (e *Editor) Delete() {
	e.text = e.text.Delete(e.curY, e.curX)
}

// LineBreak inserts a line break at the cursor position
func (e *Editor) LineBreak() {
	e.text = e.text.SplitLine(e.curY, e.curX)
	e.CursorDown()
	for e.curX > 0 {
		e.CursorLeft()
	}
}

// Backspace handles the backspace event
func (e *Editor) Backspace() {
	if e.curX == 0 {
		lineAboveLen := e.text.LineLength(e.curY - 1)
		e.text = e.text.AppendLine(e.curY-1, e.text.Line(e.curY))
		e.text = e.text.DeleteLine(e.curY)
		e.CursorUp()
		e.curX = lineAboveLen
		return
	}
	e.CursorLeft()
	e.Delete()
}

// Save saves the file
func (e *Editor) Save() {
	f, err := os.Create(e.Filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(e.text.String())
	if err != nil {
		panic(err)
	}
	f.Sync()
}
