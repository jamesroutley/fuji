package syntax

import (
	"os"

	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

// // Highlight implements a syntax highlighter
// type Highlight struct {
// 	lexer chroma.Lexer
// }

// // New returns an intialised Highlight object
// func New(filename string) *Highlight {
// 	lexer := lexers.Match(filename)
// 	return &Highlight{
// 		lexer: lexer,
// 	}
// }

func Highlight(stylename, filename, body string) {
	lexer := lexers.Match(filename)
	if lexer == nil {
		lexer = lexers.Fallback
	}
	style := styles.Get(stylename)
	if style == nil {
		style = styles.Fallback
	}
	iterator, err := lexer.Tokenise(nil, body)
	if err != nil {
		// TODO: this probably shouldn't panic - maybe return a
		// non-highlighted string?
		panic(err)
	}
	formatter := formatters.Get("html")

	f, _ := os.Create("/tmp/shout")
	if err := formatter.Format(f, style, iterator); err != nil {
		panic(err)
	}
}