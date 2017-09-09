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
