package commands

import (
	"os"

	"github.com/jamesroutley/fuji/editarea"
	termbox "github.com/nsf/termbox-go"
)

func MoveCursorUp(e *editarea.EditArea)    { e.CursorUp() }
func MoveCursorDown(e *editarea.EditArea)  { e.CursorDown() }
func MoveCursorLeft(e *editarea.EditArea)  { e.CursorLeft() }
func MoveCursorRight(e *editarea.EditArea) { e.CursorRight() }

func Quit(e *editarea.EditArea) {
	// Set cursor to 0, 0 to avoid clear screen on quit.
	termbox.SetCursor(0, 0)
	termbox.Flush()
	os.Exit(0)
}

func Insert(e *editarea.EditArea) { e.Mode = editarea.ModeInsert }
func Append(e *editarea.EditArea) {
	e.CursorRight()
	e.Mode = editarea.ModeInsert
}

func NormalMode(e *editarea.EditArea) { e.Mode = editarea.ModeNormal }

func Space(e *editarea.EditArea) { e.Insert(' ') }

func Backspace(e *editarea.EditArea) {
	e.Backspace()
}

func Save(e *editarea.EditArea) { e.Save() }

func LineBreak(e *editarea.EditArea) { e.LineBreak() }
