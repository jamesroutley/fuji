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
	text = text.Insert(0, 0, 's')

	assert.Equal(t, "shello", text.buf[0].String())
	assert.Equal(t, "world", text.buf[1].String())
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
	source := `hello
world`
	text := newTextFromString(source)
	assert.Equal(t, 2, text.Length())
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
