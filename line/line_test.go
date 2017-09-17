package line

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	t.Parallel()
	testCases := []string{
		"Hello",
		"Goodbye goodbye",
		"",
		"\n",
		"012345678901234567890",
	}
	for _, tc := range testCases {
		l := New(tc)
		assert.Equal(t, tc, l.String())
	}
}

func TestInsert(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		l        *Line
		r        rune
		loc      int
		expected *Line
	}{
		{
			&Line{[]rune{72, 72, 0, 0}, 2, 4, 4},
			'H', // H in ascii/unicode == 72
			2,
			&Line{[]rune{72, 72, 72, 0}, 3, 4, 4},
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expected, tc.l.Insert(tc.r, tc.loc))
	}
}

func TestDelete(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		l        *Line
		loc      int
		expected *Line
	}{
		{
			&Line{[]rune{72, 72, 0, 0}, 2, 4, 4},
			1,
			&Line{[]rune{72, 72, 0, 0}, 1, 4, 4},
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expected, tc.l.Delete(tc.loc))
	}
}

func TestDeleteWithMoreDeletesThanRunes(t *testing.T) {
	t.Parallel()
	l := New("cat")
	l = l.Delete(0)
	l = l.Delete(0)
	l = l.Delete(0)
	l = l.Delete(0)
	assert.Equal(t, "", l.String())
}

func TestDeleteAndInsertAndString(t *testing.T) {
	t.Parallel()
	l := New("Hello")
	l = l.Delete(0)
	l = l.Insert('B', 0)
	assert.Equal(t, "Bello", l.String())
}

func TestDuplicate(t *testing.T) {
	t.Parallel()
	l := New("Hello")
	newl := l.duplicate()
	assert.Equal(t, l.buf, newl.buf)
	assert.Equal(t, l.start, newl.start)
	assert.Equal(t, l.end, newl.end)
	assert.Equal(t, l.size, newl.size)
}

func TestSplit(t *testing.T) {
	t.Parallel()
	l := New("hello")
	a, b := l.Split(2)
	assert.Equal(t, "he", a.String())
	assert.Equal(t, "llo", b.String())
}

func TestAppend(t *testing.T) {
	t.Parallel()
	l := New("hello ")
	l = l.Append(New("world"))
	assert.Equal(t, "hello world", l.String())
}

func TestMoveGap(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		l    *Line
		loc  int
		repr string
	}{
		{
			&Line{[]rune{72, 72, 0, 0}, 2, 4, 4},
			1,
			"H__H",
		},
		{
			&Line{[]rune{72, 72, 0, 0}, 2, 4, 4},
			0,
			"__HH",
		},
		{
			&Line{[]rune{72, 72, 0, 0}, 2, 4, 4},
			4,
			"HH__",
		},
		{
			&Line{[]rune{72, 72, 0, 0}, 2, 4, 4},
			-1,
			"__HH",
		},
		{
			&Line{[]rune{72, 72, 72, 72, 0, 0}, 4, 6, 6},
			-1,
			"__HHHH",
		},
	}
	for _, tc := range testCases {
		l := tc.l.moveGap(tc.loc)
		assert.Equal(t, tc.repr, l.stringTrue())
	}
}

func TestMoveGapLeft(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		l        *Line
		expected string
	}{
		{
			New("Hi"),
			"H" + strings.Repeat("_", defaultBufferSize-len("Hi")) + "i",
		},
		{
			&Line{[]rune{72, 0, 0, 72}, 1, 3, 4},
			"__HH",
		},
	}
	for _, tc := range testCases {
		tc.l.moveGapLeft()
		assert.Equal(t, tc.expected, tc.l.stringTrue())
	}
}

func TestMoveGapRight(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		l        *Line
		expected string
	}{
		{
			&Line{[]rune{0, 0, 72, 72}, 0, 2, 4},
			"H__H",
		},
	}
	for _, tc := range testCases {
		tc.l.moveGapRight()
		assert.Equal(t, tc.expected, tc.l.stringTrue())
	}
}

// stringTrue returns a textual representation of the Line. Runes are
// displayed as runes, gaps are displayed as the rune "_"
func (l *Line) stringTrue() string {
	s := ""
	for i, r := range l.buf {
		if i >= l.start && i < l.end {
			s += "_"
		} else {
			s += string(r)
		}
	}
	return s
}
