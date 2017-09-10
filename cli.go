package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	flag.Parse()
	if flag.NArg() != 1 {
		usage()
	}
}

func filename() string {
	filename := flag.Arg(0)
	// if fileDoesNotExist(filename) {
	// 	fmt.Printf("file '%s' does not exist", filename)
	// 	os.Exit(1)
	// }
	return filename
}

func usage() {
	fmt.Println("usage: fuji [file]")
	os.Exit(1)
}

func fileDoesNotExist(fn string) bool {
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return false
	}
	return true
}
