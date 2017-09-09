package text

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()
	source := `hello
world`
	text := newTextFromString(source)
	for i, line := range strings.Split(source, "\n") {
		assert.Equal(t, line, text.Lines[i].String())
	}
}

func TestLine(t *testing.T) {
	t.Parallel()
	source := "hello"
	text := newTextFromString(source)
	assert.Equal(t, source, text.Line(0))
}

func TestInsert(t *testing.T) {
	t.Parallel()
	source := `hello
world`
	text := newTextFromString(source)
	text = text.Insert(0, 0, 's')

	assert.Equal(t, "shello", text.Lines[0].String())
	assert.Equal(t, "world", text.Lines[1].String())
}

func TestDelete(t *testing.T) {
	t.Parallel()
	source := `hello
world`
	text := newTextFromString(source)
	text = text.Delete(0, 0)

	assert.Equal(t, "ello\nworld", text.String())
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
