# Fuji

A text editor

Text

```go
func (t *Text) Insert(x, y, r) *Text
func (t *Text) Delete(x, y) *Text
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
