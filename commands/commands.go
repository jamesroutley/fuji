// Package commands implements fuji's built-in commands.
package commands

import (
	"os"

	"github.com/jamesroutley/fuji/editarea"
)

// MoveCursorUp moves the cursor up
func MoveCursorUp(e *editarea.EditArea) { e.CursorUp() }

// MoveCursorDown moves the cursor down
func MoveCursorDown(e *editarea.EditArea) { e.CursorDown() }

// MoveCursorLeft moves the cursor left
func MoveCursorLeft(e *editarea.EditArea) { e.CursorLeft() }

// MoveCursorRight moves the cursor right
func MoveCursorRight(e *editarea.EditArea) { e.CursorRight() }

// Quit quits the editor
func Quit(e *editarea.EditArea) {
	// Set cursor to 0, 0 to avoid clear screen on quit.
	// e.screen.ShowCursor(0, 0)
	// e.screen.Show()
	os.Exit(0)
}

// Insert switches the EditArea into insert mode
func Insert(e *editarea.EditArea) { e.Mode = editarea.ModeInsert }

// Append moves the cursor right and switches the EditArea into insert mode
func Append(e *editarea.EditArea) {
	e.CursorRight()
	e.Mode = editarea.ModeInsert
}

// NormalMode switches the EditArea into normal mode
func NormalMode(e *editarea.EditArea) { e.Mode = editarea.ModeNormal }

// Backspace deletes the previous rune
func Backspace(e *editarea.EditArea) {
	e.Backspace()
}

// Save saves the file being edited
func Save(e *editarea.EditArea) { e.Save() }

// LineBreak inserts a line break
func LineBreak(e *editarea.EditArea) { e.LineBreak() }

// Delete deletes the character under the cursor
func Delete(e *editarea.EditArea) { e.Delete() }

// Undo undoes the last action
func Undo(e *editarea.EditArea) { e.Undo() }

// Redo undoes the last undo
func Redo(e *editarea.EditArea) { e.Redo() }
