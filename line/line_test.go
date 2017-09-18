package line

import (
	"fmt"
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
		{
			&Line{[]rune{72, 72, 0, 0}, 2, 4, 4},
			3,
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
	assert.Equal(t, l, newl)
}

func TestSplit(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		line                 string
		split                int
		expectedA, expectedB string
	}{
		{"hello", 2, "he", "llo"},
		{"", 0, "", ""},
	}

	for _, tc := range testCases {
		t.Run(
			fmt.Sprintf("split '%s' at position %d", tc.line, tc.split),
			func(t *testing.T) {
				l := New(tc.line)
				a, b := l.Split(tc.split)
				assert.Equal(t, tc.expectedA, a.String())
				assert.Equal(t, tc.expectedB, b.String())
			},
		)
	}
}

func TestAppend(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		l1, l2, l3 string
	}{
		{"hello ", "world", "hello world"},
		{"", "hello", "hello"},
		{"hello", "", "hello"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("'%s', '%s'", tc.l1, tc.l2), func(t *testing.T) {
			l := New(tc.l1)
			l = l.Append(New(tc.l2))
			assert.Equal(t, tc.l3, l.String())
		})
	}
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
