package main

import (
	"github.com/jamesroutley/fuji/commands"
	"github.com/jamesroutley/fuji/editor"
	termbox "github.com/nsf/termbox-go"
)

func registerNormalModeCommands(e *editor.Editor) {
	e.AddNormalModeCommand("j", commands.MoveCursorDown)
	e.AddNormalModeCommand("k", commands.MoveCursorUp)
	e.AddNormalModeCommand("h", commands.MoveCursorLeft)
	e.AddNormalModeCommand("l", commands.MoveCursorRight)
	e.AddNormalModeCommand("Q", commands.Quit)
	e.AddNormalModeCommand("i", commands.Insert)
	e.AddNormalModeCommand("a", commands.Append)
	e.AddNormalModeCommand("W", commands.Save)
}

func registerInsertModeCommands(e *editor.Editor) {
	e.AddInsertModeCommand(termbox.KeyEsc, commands.NormalMode)
	e.AddInsertModeCommand(termbox.KeySpace, commands.Space)
	e.AddInsertModeCommand(termbox.KeyBackspace, commands.Backspace)
	e.AddInsertModeCommand(termbox.KeyBackspace2, commands.Backspace)
	e.AddInsertModeCommand(termbox.KeyEnter, commands.LineBreak)
}
