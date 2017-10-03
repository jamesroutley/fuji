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

- scrolling
- read in file if it exists
- think about how to compose commands - currying???
- don't panic on empty file
- Running 'a' at the end of a line should move cursor one over
- expand buffers
- maybe text.Line() should panic if int is bigger than text.Length?
- Line nums
- Git diff
- Quit sensibly -- Q should call EditArea.Quit. Editor should check for active
  buffers and close everything if there are none
- Copy/paste
