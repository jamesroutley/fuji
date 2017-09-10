package gapbuffer

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	testCases := []string{
		"Hello",
		"Goodbye goodbye",
		"",
		"\n",
		"012345678901234567890",
	}
	for _, tc := range testCases {
		gb := New(tc)
		assert.Equal(t, tc, gb.String())
	}
}

func TestInsert(t *testing.T) {
	testCases := []struct {
		gb       *GapBuffer
		r        rune
		loc      int
		expected *GapBuffer
	}{
		{
			&GapBuffer{[]rune{72, 72, 0, 0}, 2, 4, 4},
			'H', // H in ascii/unicode == 72
			2,
			&GapBuffer{[]rune{72, 72, 72, 0}, 3, 4, 4},
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expected, tc.gb.Insert(tc.r, tc.loc))
	}
}

func TestDelete(t *testing.T) {
	testCases := []struct {
		gb       *GapBuffer
		loc      int
		expected *GapBuffer
	}{
		{
			&GapBuffer{[]rune{72, 72, 0, 0}, 2, 4, 4},
			1,
			&GapBuffer{[]rune{72, 72, 0, 0}, 1, 4, 4},
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expected, tc.gb.Delete(tc.loc))
	}
}

func TestDeleteWithMoreDeletesThanRunes(t *testing.T) {
	gb := New("cat")
	gb = gb.Delete(0)
	gb = gb.Delete(0)
	gb = gb.Delete(0)
	gb = gb.Delete(0)
	assert.Equal(t, "", gb.String())
}

func TestDeleteAndInsertAndString(t *testing.T) {
	gb := New("Hello")
	gb = gb.Delete(0)
	gb = gb.Insert('B', 0)
	assert.Equal(t, "Bello", gb.String())
}

func TestDuplicate(t *testing.T) {
	gb := New("Hello")
	newGb := gb.duplicate()
	assert.Equal(t, gb.buf, newGb.buf)
	assert.Equal(t, gb.start, newGb.start)
	assert.Equal(t, gb.end, newGb.end)
	assert.Equal(t, gb.size, newGb.size)
}

func TestMoveGap(t *testing.T) {
	testCases := []struct {
		gb   *GapBuffer
		loc  int
		repr string
	}{
		{
			&GapBuffer{[]rune{72, 72, 0, 0}, 2, 4, 4},
			1,
			"H__H",
		},
		{
			&GapBuffer{[]rune{72, 72, 0, 0}, 2, 4, 4},
			0,
			"__HH",
		},
		{
			&GapBuffer{[]rune{72, 72, 0, 0}, 2, 4, 4},
			4,
			"HH__",
		},
		{
			&GapBuffer{[]rune{72, 72, 0, 0}, 2, 4, 4},
			-1,
			"__HH",
		},
		{
			&GapBuffer{[]rune{72, 72, 72, 72, 0, 0}, 4, 6, 6},
			-1,
			"__HHHH",
		},
	}
	for _, tc := range testCases {
		gb := tc.gb.moveGap(tc.loc)
		assert.Equal(t, tc.repr, gb.stringTrue())
	}
}

func TestMoveGapLeft(t *testing.T) {
	testCases := []struct {
		gb       *GapBuffer
		expected string
	}{
		{
			New("Hi"),
			"H" + strings.Repeat("_", defaultBufferSize-len("Hi")) + "i",
		},
		{
			&GapBuffer{[]rune{72, 0, 0, 72}, 1, 3, 4},
			"__HH",
		},
	}
	for _, tc := range testCases {
		tc.gb.moveGapLeft()
		assert.Equal(t, tc.expected, tc.gb.stringTrue())
	}
}

func TestMoveGapRight(t *testing.T) {
	testCases := []struct {
		gb       *GapBuffer
		expected string
	}{
		{
			&GapBuffer{[]rune{0, 0, 72, 72}, 0, 2, 4},
			"H__H",
		},
	}
	for _, tc := range testCases {
		tc.gb.moveGapRight()
		assert.Equal(t, tc.expected, tc.gb.stringTrue())
	}
}

// stringTrue returns a textual representation of the GapBuffer. Runes are
// displayed as runes, gaps are displayed as the rune "_"
func (gb *GapBuffer) stringTrue() string {
	s := ""
	for i, r := range gb.buf {
		if i >= gb.start && i < gb.end {
			s += "_"
		} else {
			s += string(r)
		}
	}
	return s
}
