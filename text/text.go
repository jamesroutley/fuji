// Package text implements a data structure which contains the text of a file.
package text

import (
	"bufio"
	"io"
	"strings"

	"github.com/jamesroutley/fuji/line"
)

const defaultGapSize = 50

// Text is the in-memory representation of the file being edited
type Text struct {
	buf   []*line.Line
	start int
	end   int
	size  int
}

// New returns an initialised Text, filled with the contents of r.
func New(r io.Reader) *Text {
	var buf []*line.Line
	scanner := bufio.NewScanner(r)
	lineCount := 0
	for scanner.Scan() {
		buf = append(buf, line.New(scanner.Text()))
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	buf = append(buf, make([]*line.Line, defaultGapSize)...)
	size := lineCount + defaultGapSize
	return &Text{
		buf:   buf,
		start: lineCount,
		end:   size,
		size:  size,
	}
}

// Line returns the contents of line i
func (t *Text) Line(i int) *line.Line {
	return t.buf[t.realIndex(i)]
}

func (t *Text) realIndex(i int) (index int) {
	index = i
	if i >= t.start {
		gapSize := t.end - t.start
		index = i + gapSize
	}
	return
}

// LineLength returns the length of line i
func (t *Text) LineLength(i int) int {
	return t.Line(i).Length()
}

// Length returns the number of lines stored by text
func (t *Text) Length() int {
	gapSize := t.end - t.start
	return t.size - gapSize
}

// String returns the contents of text as a string
func (t *Text) String() string {
	lines := make([]string, t.Length())
	i := 0
	for _, line := range t.buf[:t.start] {
		lines[i] = line.String()
		i++
	}
	for _, line := range t.buf[t.end:] {
		lines[i] = line.String()
		i++
	}
	return strings.Join(lines, "\n")
}

// Insert inserts the rune r at (row, col)
// Note that x must be relative to the start of the document, not the start of
// the currently displayed view
func (t *Text) Insert(row, col int, r rune) *Text {
	new := t.duplicate()
	new.buf[new.realIndex(col)] = new.Line(col).Insert(r, row)
	return new
}

// Delete deletes the rune at (row, col)
func (t *Text) Delete(row, col int) *Text {
	new := t.duplicate()
	new.buf[new.realIndex(col)] = new.buf[col].Delete(row)
	return new
}

// InsertLine inserts l at col
func (t *Text) InsertLine(col int, l *line.Line) *Text {
	new := t.moveGap(col)
	new.buf[col] = l
	new.start++
	return new
}

// DeleteLine deletes the line at col
func (t *Text) DeleteLine(col int) *Text {
	if t.Length() == 0 {
		return t
	}
	new := t.moveGap(col + 1)
	new.start--
	return new
}

// SplitLine splits a line at (col, row)
func (t *Text) SplitLine(row, col int) *Text {
	l := t.Line(col)
	a, b := l.Split(row)
	new := t.InsertLine(col+1, b)
	new.buf[new.realIndex(col)] = a
	return new
}

// duplicate makes a copy of a text object
func (t *Text) duplicate() *Text {
	buf := make([]*line.Line, len(t.buf))
	copy(buf, t.buf)
	return &Text{
		buf:   buf,
		start: t.start,
		end:   t.end,
		size:  t.size,
	}
}

func (t *Text) moveGap(i int) (new *Text) {
	new = t.duplicate()
	for new.start != i {
		if new.start < i {
			if new.end == new.size {
				return
			}
			new.moveGapRight()
		} else {
			if new.start == 0 {
				return
			}
			new.moveGapLeft()
		}
	}
	return
}

func (t *Text) moveGapLeft() {
	if t.start == 0 {
		return
	}
	t.buf[t.end-1] = t.buf[t.start-1]
	t.start--
	t.end--
}

func (t *Text) moveGapRight() {
	if t.end == t.size+1 {
		return
	}
	t.buf[t.start] = t.buf[t.end]
	t.start++
	t.end++
}
