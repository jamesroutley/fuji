// Package line implements a Gap Buffer, which stores a line of text and
// offers faster insert and delete times than a simple string.
// Insert and Delete operations are immutable and will return a new gap buffer.
package line

import "unicode/utf8"

const defaultBufferSize = 200

// Line implements a gap buffer
type Line struct {
	buf   []rune
	start int
	end   int
	size  int
}

// New returns an initialised Line
func New(text string) *Line {
	length := utf8.RuneCountInString(text)
	buf := make([]rune, defaultBufferSize)
	copy(buf, []rune(text))

	return &Line{
		buf:   buf,
		start: length,
		end:   defaultBufferSize,
		size:  defaultBufferSize,
	}
}

// Insert inserts rune r at position y
func (l *Line) Insert(r rune, i int) *Line {
	new := l.moveGap(i)
	new.buf[i] = r
	new.start++
	return new
}

// Delete deletes the rune at location i
func (l *Line) Delete(i int) *Line {
	if l.Length() == 0 {
		return l
	}
	new := l.moveGap(i + 1)
	new.start--
	return new
}

// String returns the string that the gap buffer contains
// TODO: bytes.buffer this?
func (l *Line) String() string {
	return string(l.buf[:l.start]) + string(l.buf[l.end:])
}

// Length returns the length of the gap buffer text
// TODO: rewrite this to count runes
// TODO: does this support unicode?
func (l *Line) Length() int {
	return len(l.String())
}

// Split splits the line at position i
func (l *Line) Split(i int) (*Line, *Line) {
	new := l.moveGap(i)
	return New(string(new.buf[:new.start])), New(string(new.buf[new.end:]))
}

// Append appends l2 to l
func (l *Line) Append(l2 *Line) *Line {
	return New(l.String() + l2.String())
}

// duplicate deep copies a Line
func (l *Line) duplicate() *Line {
	if l == nil {
		panic(nil)
	}
	newBuf := make([]rune, l.size)
	copy(newBuf, l.buf)
	return &Line{
		newBuf,
		l.start,
		l.end,
		l.size,
	}
}

// moveGap moves l's gap so it starts at position i.
// moveGap will move the gap to the furthest left or right position possible if
// i is out of array bounds.
func (l *Line) moveGap(i int) (newl *Line) {
	newl = l.duplicate()
	for newl.start != i {
		if newl.start < i {
			if newl.end == newl.size {
				return
			}
			newl.moveGapRight()
		} else {
			if newl.start == 0 {
				return
			}
			newl.moveGapLeft()
		}
	}
	return
}

// moveGapLeft moves l's gap one space left.
// It is not immutable, and should not be called directly
func (l *Line) moveGapLeft() {
	if l.start == 0 {
		return
	}
	l.buf[l.end-1] = l.buf[l.start-1]
	l.start--
	l.end--
}

// moveGapRight moves l's gap one space right.
// It is not immutable, and should not be called directly
func (l *Line) moveGapRight() {
	if l.end == l.size+1 {
		return
	}
	l.buf[l.start] = l.buf[l.end]
	l.start++
	l.end++
}
