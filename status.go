package main

import (
	"github.com/jamesroutley/fuji/statusbar"
	status "github.com/jamesroutley/fuji/statuses"
)

func registerStatuses() {
	statusbar.AddStatus(status.Mode)
	statusbar.AddStatus(status.Filename)
	statusbar.AddStatus(status.GitBranch)
}
