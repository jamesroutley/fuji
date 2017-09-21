package main

import (
	"github.com/jamesroutley/fuji/commands"
	"github.com/jamesroutley/fuji/editarea"
	termbox "github.com/nsf/termbox-go"
)

func registerNormalModeCommands() {
	editarea.AddNormalModeCommand("j", commands.MoveCursorDown)
	editarea.AddNormalModeCommand("k", commands.MoveCursorUp)
	editarea.AddNormalModeCommand("h", commands.MoveCursorLeft)
	editarea.AddNormalModeCommand("l", commands.MoveCursorRight)
	editarea.AddNormalModeCommand("Q", commands.Quit)
	editarea.AddNormalModeCommand("i", commands.Insert)
	editarea.AddNormalModeCommand("a", commands.Append)
	editarea.AddNormalModeCommand("W", commands.Save)
	editarea.AddNormalModeCommand("x", commands.Delete)
	editarea.AddNormalModeCommand("u", commands.Undo)
	editarea.AddNormalModeCommand("R", commands.Redo)
}

func registerInsertModeCommands() {
	editarea.AddInsertModeCommand(termbox.KeyEsc, commands.NormalMode)
	editarea.AddInsertModeCommand(termbox.KeySpace, commands.Space)
	editarea.AddInsertModeCommand(termbox.KeyBackspace, commands.Backspace)
	editarea.AddInsertModeCommand(termbox.KeyBackspace2, commands.Backspace)
	editarea.AddInsertModeCommand(termbox.KeyEnter, commands.LineBreak)
	editarea.AddInsertModeCommand(termbox.KeyArrowDown, commands.MoveCursorDown)
	editarea.AddInsertModeCommand(termbox.KeyArrowUp, commands.MoveCursorUp)
	editarea.AddInsertModeCommand(termbox.KeyArrowLeft, commands.MoveCursorLeft)
	editarea.AddInsertModeCommand(termbox.KeyArrowRight, commands.MoveCursorRight)
}
