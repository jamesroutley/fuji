package main

import (
	"github.com/gdamore/tcell"
	"github.com/jamesroutley/fuji/commands"
	"github.com/jamesroutley/fuji/editarea"
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
	editarea.AddInsertModeCommand(tcell.KeyESC, commands.NormalMode)
	// editarea.AddInsertModeCommand(tcell.KeySpace, commands.Space)
	editarea.AddInsertModeCommand(tcell.KeyBackspace, commands.Backspace)
	editarea.AddInsertModeCommand(tcell.KeyBackspace2, commands.Backspace)
	editarea.AddInsertModeCommand(tcell.KeyEnter, commands.LineBreak)
	editarea.AddInsertModeCommand(tcell.KeyDown, commands.MoveCursorDown)
	editarea.AddInsertModeCommand(tcell.KeyUp, commands.MoveCursorUp)
	editarea.AddInsertModeCommand(tcell.KeyLeft, commands.MoveCursorLeft)
	editarea.AddInsertModeCommand(tcell.KeyRight, commands.MoveCursorRight)
}
