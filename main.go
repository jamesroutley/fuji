package main

import "github.com/jamesroutley/fuji/editor"

func main() {
	e := editor.New(filename())
	registerNormalModeCommands(e)
	registerInsertModeCommands(e)
	e.Start()
}
