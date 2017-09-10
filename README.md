# Fuji

A text editor

Editor

```go
CursorUp()
CursorDown()
CursorLeft()
CursorRight()
Get() (rune)
Set(r rune)
```

Text

```go
func (t *Text) Insert(x, y, r) *Text
func (t *Text) Delete(x, y) *Text
func (t *Text) String()
```

GapBuffer

TODO:
- Add automatic expansion when the gap becomes 0
- [Low priority] Add automatic downsizing when the gap becomes large

API

```go
func (gb *GapBuffer) Insert(y, r) *GapBuffer
func (gb *GapBuffer) Delete(y) *GapBuffer
func (gb *GapBuffer) Length() int
func (gb *GapBuffer) String() int
```

Todo:

- Sensible cursor movement
- insert at position 0 is broken
- read in file if it exists
- think about how to compose commands - currying???
- don't panic on empty file
- Running 'a' at the end of a line should move cursor one over
