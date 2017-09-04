// Package gapbuffer implements a Gap Buffer, which stores a line of text and
// offers faster insert and delete times than a simple string.
package gapbuffer

import "unicode/utf8"

const defaultBufferSize = 20

// GapBuffer implements a gap buffer
type GapBuffer struct {
	buf   []rune
	start int
	end   int
	size  int
}

// New returns an initialised GapBuffer
func New(text string) *GapBuffer {
	length := utf8.RuneCountInString(text)
	buf := make([]rune, defaultBufferSize)
	copy(buf, []rune(text))

	return &GapBuffer{
		buf:   buf,
		start: length,
		end:   defaultBufferSize,
		size:  defaultBufferSize,
	}
}

// Insert inserts rune r at position y
func (gb *GapBuffer) Insert(r rune, y int) *GapBuffer {
	newGb := gb.moveGap(y)
	newGb.buf[y] = r
	newGb.start++

	return newGb
}

// Delete deletes the rune at location loc
func (gb *GapBuffer) Delete(loc int) *GapBuffer {
	newGb := gb.moveGap(loc + 1)
	newGb.start--
	return newGb
}

// String returns the string that the gap buffer contains
func (gb *GapBuffer) String() string {
	return string(gb.buf[:gb.start]) + string(gb.buf[gb.end:])
}

// duplicate deep copies a GapBuffer
func (gb *GapBuffer) duplicate() *GapBuffer {
	newBuf := make([]rune, gb.size)
	copy(newBuf, gb.buf)
	return &GapBuffer{
		newBuf,
		gb.start,
		gb.end,
		gb.size,
	}
}

// moveGap moves gb's gap so it starts at position i.
// moveGap will move the gap to the furthest left or right position possible if
// i is out of array bounds.
func (gb *GapBuffer) moveGap(i int) (newGb *GapBuffer) {
	newGb = gb.duplicate()
	for newGb.start != i {
		if newGb.start < i {
			if newGb.end == newGb.size {
				return
			}
			newGb.moveGapRight()
		} else {
			if newGb.start == 0 {
				return
			}
			newGb.moveGapLeft()
		}
	}
	return
}

// moveGapLeft moves gb's gap one space left.
// It is not immutable, does not perform bounds checking and should never be
// called directly
func (gb *GapBuffer) moveGapLeft() {
	tmp := gb.buf[gb.start-1]
	gb.buf[gb.start-1] = gb.buf[gb.end-1]
	gb.buf[gb.end-1] = tmp
	gb.start--
	gb.end--
}

// moveGapRight moves gb's gap one space right.
// It is not immutable, does not perform bounds checking and should never be
// called directly
func (gb *GapBuffer) moveGapRight() {
	tmp := gb.buf[gb.start]
	gb.buf[gb.start] = gb.buf[gb.end]
	gb.buf[gb.end] = tmp
	gb.start++
	gb.end++
}
