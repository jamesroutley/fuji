package text

import (
	"strings"
	"testing"

	"github.com/jamesroutley/fuji/line"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()
	source := `hello
world`
	text := newTextFromString(source)
	for i, line := range strings.Split(source, "\n") {
		assert.Equal(t, line, text.buf[i].String())
	}
	sourceLen := len(strings.Split(source, "\n"))
	for i := sourceLen; i < sourceLen+defaultGapSize; i++ {
		assert.Nil(t, text.buf[i])
	}
	assert.Equal(t, 2, text.start)
	assert.Equal(t, 2+defaultGapSize, text.end)
}

func TestRealIndex(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		source string
		gapLoc int
		i      int
		realI  int
	}{
		{"hello\nworld", 0, 1, defaultGapSize + 1},
		{"hello\nworld", 1, 1, defaultGapSize + 1},
	}
	for _, tc := range testCases {
		t.Run(tc.source, func(t *testing.T) {
			text := newTextFromString(tc.source)
			text = text.moveGap(tc.gapLoc)
			index := text.realIndex(tc.i)
			assert.Equal(t, tc.realI, index)
		})
	}
}

func TestRealIndexWithSingleLine(t *testing.T) {
	t.Parallel()
	text := newTextFromString("hello")
	assert.Equal(t, 0, text.realIndex(0))
}

func TestLineString(t *testing.T) {
	t.Parallel()
	source := "hello"
	text := newTextFromString(source)
	assert.Equal(t, source, text.Line(0).String())
}

func TestInsert(t *testing.T) {
	t.Parallel()
	source := `hello
world`
	text := newTextFromString(source)
	text = text.Insert(1, 0, 's')

	assert.Equal(t, "hello", text.buf[0].String())
	assert.Equal(t, "sworld", text.buf[1].String())
}

func TestDelete(t *testing.T) {
	t.Parallel()
	source := `hello
world`
	text := newTextFromString(source)
	text = text.Delete(0, 0)
	text = text.Delete(0, 0)
	text = text.Delete(0, 0)

	assert.Equal(t, "lo\nworld", text.String())
}

func TestInsertLine(t *testing.T) {
	t.Parallel()
	source := `hello
world`
	text := newTextFromString(source)
	l := line.New("this")
	text = text.InsertLine(1, l)
	expected := `hello
this
world`
	assert.Equal(t, expected, text.String())

}

func TestDeleteLine(t *testing.T) {
	t.Parallel()
	source := `hello
world`
	text := newTextFromString(source)
	text = text.DeleteLine(0)
	assert.Equal(t, "world", text.String())
}

func TestSplitLine(t *testing.T) {
	t.Parallel()
	text := newTextFromString("hello")
	text = text.SplitLine(0, 2)
	expected := `he
llo`
	assert.Equal(t, expected, text.String())
}

func TestLength(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		source string
		len    int
	}{
		{"hello\nworld", 2},
		{"1 \n2 \n3 \n4 \n5", 5},
		{"", 0},
	}
	for _, tc := range testCases {
		t.Run(tc.source, func(t *testing.T) {
			text := newTextFromString(tc.source)
			assert.Equal(t, tc.len, text.Length())
		})
	}
}

func TestDeletePreservesLengthCorrectly(t *testing.T) {
	t.Parallel()
	text := newTextFromString("hello\nworld")
	assert.Equal(t, 2, text.Length())
	text = text.DeleteLine(0)
	assert.Equal(t, 1, text.Length())
}

func TestString(t *testing.T) {
	t.Parallel()
	source := `hello
	world`
	text := newTextFromString(source)
	assert.Equal(t, source, text.String())
}

func newTextFromString(source string) *Text {
	return New(strings.NewReader(source))
}

func (t Text) stringTrue() string {
	var lines []string
	for i, line := range t.buf {
		if i == t.start {
			lines = append(lines, "s")
		} else if i == t.end {
			lines = append(lines, "e")
		} else if i >= t.start && i < t.end {
			lines = append(lines, "_")
		} else {
			lines = append(lines, line.String())
		}
	}
	return strings.Join(lines, "\n")
}
