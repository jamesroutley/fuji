package editarea

import (
	"io"
	"math"
	"os"

	"github.com/gdamore/tcell"
	"github.com/jamesroutley/fuji/area"
	"github.com/jamesroutley/fuji/logger"
	"github.com/jamesroutley/fuji/syntax"
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
	cursor     area.Point
	beenEdited bool
	beenSaved  bool
	screen     tcell.Screen
	lineno     int
	displayLen int
}

// New returns a new EditArea
func New(screen tcell.Screen, filename string, r io.Reader) *EditArea {
	t := text.New(r)
	return &EditArea{
		Filename:   filename,
		Mode:       ModeNormal,
		history:    newHistory(t, area.Point{X: 0, Y: 0}, 50),
		text:       t,
		cursor:     area.Point{X: 0, Y: 0},
		beenEdited: false,
		beenSaved:  true,
		screen:     screen,
		lineno:     0,
		displayLen: 0,
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
func (e *EditArea) Draw(a area.Area) {
	e.displayLen = a.End.Y - a.Start.Y
	if e.beenEdited {
		e.history.add(e.text, e.cursor)
		e.beenEdited = false
	}

	// TODO: Need to do this in case the iterator panics
	// defer func() {
	// 	if perr := recover(); perr != nil {
	// 		err = perr.(error)
	// 	}
	// }()

	styledRunes := syntax.Highlight(e.Filename, e.text.String())

	yMin := math.MaxInt32
	for _, i := range []int{a.End.Y, e.text.Length()} {
		if i < yMin {
			yMin = i
		}
	}
	for y := a.Start.Y; y < yMin; y++ {
		for x, sr := range styledRunes[y+e.lineno] {
			switch sr.Rune {
			case '\t':
				for i := 0; i < 4; i++ {
					e.screen.SetContent(x, y, ' ', nil, sr.Style)
					x++
				}
			default:
				e.screen.SetContent(x, y, sr.Rune, nil, sr.Style)
				x++
			}
		}
	}

	e.displayCursor()
}

func (e *EditArea) displayCursor() {
	x := e.cursor.X
	maxX := e.cursorMaxX()

	// If the cursor x is greater than the number of characters on that line,
	// display the cursor at the end of the line
	if x > maxX {
		x = maxX
	}
	// TODO: pretty hacky!
	// if x < 0 {
	// 	x = 0
	// }
	logger.L.Print("cursorMaxX: ", maxX)
	logger.L.Print("cursor: ", e.cursor.X, e.cursor.Y)
	e.screen.ShowCursor(x, e.cursor.Y)
}

// cursorMaxX returns the maximum x that the cursor can be at for the current
// line. This x value depends on the mode of the editor. In insert mode, the
// cursor can be one square further to the right, like in vim.
func (e *EditArea) cursorMaxX() (x int) {
	x = e.text.LineLength(e.cursor.Y)
	if x == 0 {
		return
	}
	if e.Mode == ModeNormal {
		x--
	}
	return
}

// CursorUp moves the cursor up
func (e *EditArea) CursorUp() {
	if e.cursor.Y+e.lineno == 0 {
		return
	}
	if e.cursor.Y < 10 && e.lineno > 0 {
		e.lineno--
		return
	}
	e.cursor.Y--
}

// CursorDown moves the cursor down
func (e *EditArea) CursorDown() {
	// Return early if the cursor is at the end of the document
	if e.cursor.Y+e.lineno >= e.text.Length()-1 {
		return
	}
	// If the cursor is under 10 lines from the end of the displayed document
	// and there is room to scroll
	if e.displayLen-e.cursor.Y < 10 && e.lineno+e.displayLen < e.text.Length() {
		e.lineno++
		return
	}
	e.cursor.Y++
}

// CursorLeft moves the cursor left
func (e *EditArea) CursorLeft() {
	// If the cursor x is greater than the maximum x for that line, decrement
	// x until the cursor is at the end of the line.
	xMax := e.cursorMaxX()
	if e.cursor.X > xMax {
		e.cursor.X = xMax
	}
	// Don't do anything if the cursor is at (0, 0)
	if e.cursor.X <= 0 && e.cursor.Y <= 0 {
		return
	}
	// If the x == 0, move the cursor to the end of the line above
	if e.cursor.X <= 0 {
		e.cursor.Y--
		e.cursor.X = e.cursorMaxX()
		return
	}
	e.cursor.X--
}

// CursorRight moves the cursor right
func (e *EditArea) CursorRight() {
	// Don'e move cursor past end of document
	if e.cursor.X >= e.text.LineLength(e.cursor.Y)-1 && e.cursor.Y >= e.text.Length()-1 {
		return
	}

	// If at the end of a line, move cursor to beginning of next
	if e.cursor.X >= e.cursorMaxX() {
		e.cursor.Y++
		e.cursor.X = 0
		return
	}
	e.cursor.X++
}

// CursorAtLineStart returns whether the cursor is at the beginning of a line
func (e *EditArea) CursorAtLineStart() bool {
	return e.cursor.X == 0
}

// CursorAtLineEnd returns whether the cursor is at the end of a line
func (e *EditArea) CursorAtLineEnd() bool {
	return e.cursor.X == e.cursorMaxX()
}

// CursorAtTextStart returns whether the cursor is at the beginning of the text
func (e *EditArea) CursorAtTextStart() bool {
	return e.cursor.Y == 0
}

// CursorAtTextEnd returns whether the cursor is at the end of the text
func (e *EditArea) CursorAtTextEnd() bool {
	return e.cursor.Y == e.text.Length()-1
}

// Insert inserts rune r at the cursor position
func (e *EditArea) Insert(r rune) {
	e.beenEdited = true
	e.text = e.text.Insert(e.cursor.Y, e.cursor.X, r)
	e.CursorRight()
	e.beenSaved = false
}

// Delete deletes the rune under the cursor
func (e *EditArea) Delete() {
	e.beenEdited = true
	e.text = e.text.Delete(e.cursor.Y, e.cursor.X)
	e.beenSaved = false
}

// LineBreak inserts a line break at the cursor position
func (e *EditArea) LineBreak() {
	e.beenEdited = true
	e.text = e.text.SplitLine(e.cursor.Y, e.cursor.X)
	e.CursorDown()
	for e.cursor.X > 0 {
		e.CursorLeft()
	}
	e.beenSaved = false
}

// Backspace handles the backspace event
func (e *EditArea) Backspace() {
	e.beenEdited = true
	if e.cursor.X == 0 && e.cursor.Y == 0 {
		return
	}
	if e.cursor.X == 0 {
		lineAboveLen := e.text.LineLength(e.cursor.Y - 1)
		e.text = e.text.AppendLine(e.cursor.Y-1, e.text.Line(e.cursor.Y))
		e.text = e.text.DeleteLine(e.cursor.Y)
		e.CursorUp()
		e.cursor.X = lineAboveLen
		return
	}
	e.CursorLeft()
	e.Delete()
	e.beenSaved = false
}

// Peek returns the rune under the cursor
func (e *EditArea) Peek() (r rune) {
	r, _, _, _ = e.screen.GetContent(e.cursor.X, e.cursor.Y)
	return
}

// Undo undoes the last action
func (e *EditArea) Undo() {
	e.history.undo()
	e.text = e.history.head.text
	e.cursor.X = e.history.head.cursor.X
	e.cursor.Y = e.history.head.cursor.Y
}

// Redo undoes the last undo
func (e *EditArea) Redo() {
	e.history.redo()
	e.text = e.history.head.text
	e.cursor.X = e.history.head.cursor.X
	e.cursor.Y = e.history.head.cursor.Y
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
	e.beenSaved = true
}

// Saved returns a boolean indicating whether the current file has been saved
func (e *EditArea) Saved() bool {
	return e.beenSaved
}
