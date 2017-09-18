package logger

import (
	"io"
	"log"
)

// L is a logger
var L *log.Logger

// Init initialises the logger
func Init(out io.Writer) {
	L = log.New(out, "", 0)
}
