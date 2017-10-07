package editarea

import (
	"github.com/jamesroutley/fuji/area"
	"github.com/jamesroutley/fuji/text"
)

// state stores the text and cursor position at a particular time
type state struct {
	text   *text.Text
	cursor area.Point
	next   *state
	prev   *state
}

// history implements an EditArea's undo and redo functionality.
// It is a doubly linked list of states.
type history struct {
	head        *state
	tail        *state
	len, maxLen int
}

// newHistory instantiates a new history object
func newHistory(t *text.Text, cursor area.Point, maxLen int) *history {
	s := &state{
		text:   t,
		cursor: cursor,
		next:   nil,
		prev:   nil,
	}
	return &history{head: s, tail: s, len: 0, maxLen: maxLen}
}

// add adds a state to the history
func (h *history) add(t *text.Text, cursor area.Point) {
	if h.len > h.maxLen {
		h.forget()
	}
	s := &state{
		text:   t,
		cursor: cursor,
		next:   nil,
		prev:   h.head,
	}
	h.head.next = s
	h.head = s
	h.len++
}

// forget forgets the oldest state.
func (h *history) forget() {
	h.tail = h.tail.next
	h.tail.prev = nil
	h.len--
}

// undo puts h.head into a previous state
func (h *history) undo() {
	if h.head.prev == nil {
		return
	}
	h.head = h.head.prev
	h.len--
}

// redo puts h.head into a future state
func (h *history) redo() {
	if h.head.next == nil {
		return
	}
	h.head = h.head.next
	h.len++
}
