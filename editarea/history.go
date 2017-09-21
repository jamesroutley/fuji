package editarea

import "github.com/jamesroutley/fuji/text"

// state stores the text and cursor position at a particular time
type state struct {
	text       *text.Text
	curX, curY int
	next       *state
	prev       *state
}

// history implements an EditArea's undo and redo functionality.
// It is a doubly linked list of states.
type history struct {
	head *state
	tail *state
}

// newHistory instantiates a new history object
func newHistory(t *text.Text, curX, curY int) *history {
	s := &state{
		text: t,
		curX: curX,
		curY: curY,
		next: nil,
		prev: nil,
	}
	return &history{head: s, tail: s}
}

// add adds a state to the history
func (h *history) add(t *text.Text, curX, curY int) {
	s := &state{
		text: t,
		curX: curX,
		curY: curY,
		next: nil,
		prev: h.head,
	}
	h.head.next = s
	h.head = s
}

// forget forgets the oldest state.
func (h *history) forget() {
	h.tail = h.tail.next
	h.tail.prev = nil
}

// undo puts h.head into a previous state
func (h *history) undo() {
	if h.head.prev == nil {
		return
	}
	h.head = h.head.prev
}

// redo puts h.head into a future state
func (h *history) redo() {
	if h.head.next == nil {
		return
	}
	h.head = h.head.next
}
