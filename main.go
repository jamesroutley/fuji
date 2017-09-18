package main

import (
	"os"

	"github.com/jamesroutley/fuji/editor"
	"github.com/jamesroutley/fuji/logger"
)

func init() {
}

func main() {
	logfile, err := os.Create("/tmp/fujilog")
	if err != nil {
		panic("can't log")
	}
	defer logfile.Close()
	logger.Init(logfile)

	e := editor.Editor{}
	registerNormalModeCommands()
	registerInsertModeCommands()
	e.Start(filename())
}
