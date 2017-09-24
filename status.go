package main

import (
	"github.com/jamesroutley/fuji/editarea"
	status "github.com/jamesroutley/fuji/statuses"
)

func registerStatuses() {
	editarea.AddStatus(status.Mode)
	editarea.AddStatus(status.Filename)
	editarea.AddStatus(status.GitBranch)
}
