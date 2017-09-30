package area

// Point represents a point (x, y) in the screen.
// X is horizontal. Y is vertial
type Point struct {
	X, Y int
}

// Area represents an area of the screen
type Area struct {
	Start, End Point
}
