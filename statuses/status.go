package status

import (
	"os/exec"
	"path/filepath"

	"github.com/jamesroutley/fuji/editarea"
)

// Mode returns the editor's mode
func Mode(e *editarea.EditArea) string {
	switch e.Mode {
	case editarea.ModeNormal:
		return "Normal"
	case editarea.ModeInsert:
		return "Insert"
	default:
		return "Error: unimplemented"
	}
}

// Filename return the name of the file being edited
func Filename(e *editarea.EditArea) string {
	_, fn := filepath.Split(e.Filename)
	if e.Saved() {
		return fn
	}
	return fn + " *"
}

// GitBranch returns the current git branch
func GitBranch(e *editarea.EditArea) string {
	cmdOut, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return ""
	}
	return string(cmdOut)
}
