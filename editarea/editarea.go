package editarea

import (
	"io"
	"os"

	"github.com/gdamore/tcell"
	"github.com/jamesroutley/fuji/text"
)

// Mode distinguishes between editor modes
type Mode uint8

const (
	// ModeNormal indicates the editor is in normal mode
	ModeNormal Mode = iota
	// ModeInsert indicates the editor is in insert mode
	ModeInsert
)

// NormalModeCommand is a function that defines the behaviour of a normal mode
// command
// TODO: this type is the same as InsertModeCommand - there really isn't
// a difference between them. They should be combined. Maybe the type should
// be called Behaviour?
type NormalModeCommand func(*EditArea)

// InsertModeCommand defines the behaviour of an insert mode command
type InsertModeCommand func(*EditArea)

var normalModeCommands = make(map[string]NormalModeCommand)
var insertModeCommands = make(map[tcell.Key]InsertModeCommand)

// EditArea exposes the main API of the text editor
type EditArea struct {
	Filename   string
	Mode       Mode
	history    *history
	text       *text.Text
	curX, curY int
	beenEdited bool
	screen     tcell.Screen
}

// New returns a new EditArea
func New(screen tcell.Screen, filename string, r io.ReadWriter) *EditArea {
	t := text.New(r)
	return &EditArea{
		Filename: filename,
		Mode:     ModeNormal,
		history:  newHistory(t, 0, 0, 50),
		text:     t,
		curX:     0,
		curY:     0,
		screen:   screen,
	}
}

// HandleEvent handles the tcell event ev
func (e *EditArea) HandleEvent(ev *tcell.EventKey) {
	switch e.Mode {
	case ModeNormal:
		e.handleNormalModeEvent(ev)
	case ModeInsert:
		e.handleInsertModeEvent(ev)
	default:
		panic("Should not reach here")
	}
}

func (e *EditArea) handleNormalModeEvent(ev *tcell.EventKey) {
	if ev.Key() != tcell.KeyRune {
		return
	}
	command := normalModeCommands[string(ev.Rune())]
	if command == nil {
		return
	}
	command(e)
}

func (e *EditArea) handleInsertModeEvent(ev *tcell.EventKey) {
	command := insertModeCommands[ev.Key()]
	if command != nil {
		command(e)
		return
	}
	if ev.Key() != tcell.KeyRune {
		return
	}
	e.Insert(ev.Rune())
}

// AddNormalModeCommand adds a new command to the editor
func AddNormalModeCommand(name string, behaviour NormalModeCommand) {
	normalModeCommands[name] = behaviour
}

// AddInsertModeCommand adds a new insert mode command
func AddInsertModeCommand(key tcell.Key, behaviour InsertModeCommand) {
	insertModeCommands[key] = behaviour
}

// Draw writes the contents of the EditArea to tcell's internal buffer.
// screen.Show() should be called after Draw() to write the contents to
// the screen
func (e *EditArea) Draw() {
	if e.beenEdited {
		e.history.add(e.text, e.curX, e.curY)
		e.beenEdited = false
	}
	// Clear screen
	e.screen.Clear()
	for y := 0; y < e.text.Length(); y++ {
		line := e.text.Line(y)
		for x, r := range line.String() {
			e.screen.SetContent(x, y, r, nil, tcell.StyleDefault)
		}
	}
	e.displayCursor()
}

func (e *EditArea) displayCursor() {
	x := e.curX
	lineLength := e.text.LineLength(e.curY)
	if x >= lineLength-1 {
		x = lineLength - 1
	}
	// TODO: pretty hacky!
	if x < 0 {
		x = 0
	}
	e.screen.ShowCursor(x, e.curY)
}

// CursorUp moves the cursor up
func (e *EditArea) CursorUp() {
	if e.curY == 0 {
		return
	}
	e.curY--
}

// CursorDown moves the cursor down
func (e *EditArea) CursorDown() {
	if e.curY >= e.text.Length()-1 {
		return
	}
	e.curY++
}

// CursorLeft moves the cursor left
func (e *EditArea) CursorLeft() {
	if e.curX > e.text.LineLength(e.curY)-1 {
		e.curX = e.text.LineLength(e.curY) - 1
	}
	if e.curX <= 0 && e.curY <= 0 {
		return
	}
	if e.curX <= 0 {
		e.curY--
		e.curX = e.text.LineLength(e.curY) - 1
		return
	}
	e.curX--
}

// CursorRight moves the cursor right
func (e *EditArea) CursorRight() {
	// Don'e move cursor past end of document
	if e.curX >= e.text.LineLength(e.curY)-1 && e.curY >= e.text.Length()-1 {
		return
	}
	// If at the end of a line, move cursor to beginning of next
	if e.curX >= e.text.LineLength(e.curY)-1 {
		e.curY++
		e.curX = 0
		return
	}
	e.curX++
}

// CursorAtLineStart returns whether the cursor is at the beginning of a line
func (e *EditArea) CursorAtLineStart() bool {
	return e.curX == 0
}

// CursorAtLineEnd returns whether the cursor is at the end of a line
func (e *EditArea) CursorAtLineEnd() bool {
	return e.curX == e.text.LineLength(e.curY)-1
}

// Insert inserts rune r at the cursor position
func (e *EditArea) Insert(r rune) {
	e.beenEdited = true
	e.text = e.text.Insert(e.curY, e.curX, r)
	e.CursorRight()
}

// Delete deletes the rune under the cursor
func (e *EditArea) Delete() {
	e.beenEdited = true
	e.text = e.text.Delete(e.curY, e.curX)
}

// LineBreak inserts a line break at the cursor position
func (e *EditArea) LineBreak() {
	e.beenEdited = true
	e.text = e.text.SplitLine(e.curY, e.curX)
	e.CursorDown()
	for e.curX > 0 {
		e.CursorLeft()
	}
}

// Backspace handles the backspace event
func (e *EditArea) Backspace() {
	e.beenEdited = true
	if e.curX == 0 && e.curY == 0 {
		return
	}
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

// Undo undoes the last action
func (e *EditArea) Undo() {
	e.history.undo()
	e.text = e.history.head.text
	e.curX = e.history.head.curX
	e.curY = e.history.head.curY
}

// Redo undoes the last undo
func (e *EditArea) Redo() {
	e.history.redo()
	e.text = e.history.head.text
	e.curX = e.history.head.curX
	e.curY = e.history.head.curY
}

// Save saves the file
func (e *EditArea) Save() {
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
