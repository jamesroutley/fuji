package main

import (
	"os"

	"github.com/jamesroutley/fuji/editor"
)

func main() {
	f, err := os.Open("README.md")
	if err != nil {
		panic(err)
	}
	e := editor.New(f)
	e.Start()
}
