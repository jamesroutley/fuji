package text

import (
	"bufio"
	"io"
	"strings"

	"github.com/jamesroutley/fuji/gapbuffer"
)

// Text is the in-memory representation of the file being edited
type Text struct {
	Lines []*gapbuffer.GapBuffer
}

// New returns an initialised Text, filled with the contents of r.
func New(r io.Reader) *Text {
	var lines []*gapbuffer.GapBuffer
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, gapbuffer.New(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return &Text{Lines: lines}
}

// Line returns the contents of line i
func (text *Text) Line(i int) string {
	return text.Lines[i].String()
}

func (text *Text) String() string {
	lineStrings := make([]string, len(text.Lines))
	for i, line := range text.Lines {
		lineStrings[i] = line.String()
	}
	return strings.Join(lineStrings, "\n")
}

// Insert inserts the rune r at position x, y
// Note that x must be relative to the start of the document, not the start of
// the currently displayed view
func (text *Text) Insert(x, y int, r rune) *Text {
	newText := text.duplicate()
	newText.Lines[x] = newText.Lines[x].Insert(r, y)
	return newText
}

// Delete deletes the rune at position x, y
func (text *Text) Delete(x, y int) *Text {
	newText := text.duplicate()
	newText.Lines[x] = newText.Lines[x].Delete(y)
	return newText
}

// duplicate makes a copy of a text object
func (text *Text) duplicate() *Text {
	lines := make([]*gapbuffer.GapBuffer, len(text.Lines))
	copy(lines, text.Lines)
	return &Text{
		Lines: lines,
	}
}
