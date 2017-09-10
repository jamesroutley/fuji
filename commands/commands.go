package commands

import (
	"os"

	"github.com/jamesroutley/fuji/editor"
)

func MoveCursorUp(e *editor.Editor)    { e.CursorUp() }
func MoveCursorDown(e *editor.Editor)  { e.CursorDown() }
func MoveCursorLeft(e *editor.Editor)  { e.CursorLeft() }
func MoveCursorRight(e *editor.Editor) { e.CursorRight() }

func Quit(e *editor.Editor) {

	os.Exit(0)
}

func Insert(e *editor.Editor) { e.Mode = editor.ModeInsert }
func Append(e *editor.Editor) {
	e.CursorRight()
	e.Mode = editor.ModeInsert
}

func NormalMode(e *editor.Editor) { e.Mode = editor.ModeNormal }

func Space(e *editor.Editor) { e.Insert(' ') }

func Backspace(e *editor.Editor) {
	e.CursorLeft()
	e.Delete()
}

func Save(e *editor.Editor) { e.Save() }