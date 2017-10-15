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
	e.Mode = editarea.ModeInsert
	e.CursorRight()
}

// NormalMode switches the EditArea into normal mode
func NormalMode(e *editarea.EditArea) {
	e.CursorLeft()
	e.Mode = editarea.ModeNormal
}

// Backspace deletes the previous rune
func Backspace(e *editarea.EditArea) {
	e.Backspace()
}

// JmpToLineEnd moves cursor to line end
func JmpToLineEnd(e *editarea.EditArea) {
	for !e.CursorAtLineEnd() {
		e.CursorRight()
	}
}

// JmpToLineStart moves the cursor to line start
func JmpToLineStart(e *editarea.EditArea) {
	for !e.CursorAtLineStart() {
		e.CursorLeft()
	}
}

// JmpToWordStart moves the cursor to the beginning of the word
// TODO: this won't stop at non-space chars like . or (
func JmpToWordStart(e *editarea.EditArea) {
	// This is a hack to allow the command to be run repeatedly to jump back
	// multiple words
	e.CursorLeft()
	e.CursorLeft()
	for e.Peek() != ' ' {
		if e.CursorAtLineStart() {
			return
		}
		e.CursorLeft()
	}
	e.CursorRight()
}

// JmpToWordEnd moves the cursor the the beginning of the next word
func JmpToWordEnd(e *editarea.EditArea) {
	for e.Peek() != ' ' {
		if e.CursorAtTextEnd() && e.CursorAtLineEnd() {
			return
		}
		e.CursorRight()
	}
	e.CursorRight()
}

// atEmptyLine returns a bool indicating whether the line that the cursor is at
// is empty.
func atEmptyLine(e *editarea.EditArea) bool {
	return e.CursorAtLineStart() && e.CursorAtLineEnd()
}

// JmpToParagraphEnd moves the cursor to the next empty line
func JmpToParagraphEnd(e *editarea.EditArea) {
	// If we are already on an empty line, skip to the next non-empty line
	for atEmptyLine(e) {
		if e.CursorAtTextEnd() {
			return
		}
		e.CursorDown()
	}

	for !atEmptyLine(e) {
		if e.CursorAtTextEnd() {
			return
		}
		e.CursorDown()
	}
}

// JmpToParagraphStart moves the cursor to the previous empty line
func JmpToParagraphStart(e *editarea.EditArea) {
	// If we are already on an empty line, skip to the next non-empty line
	for atEmptyLine(e) {
		if e.CursorAtTextStart() {
			return
		}
		e.CursorUp()
	}

	for !atEmptyLine(e) {
		if e.CursorAtTextStart() {
			return
		}
		e.CursorUp()
	}

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
