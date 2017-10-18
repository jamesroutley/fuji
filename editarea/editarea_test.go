package editarea

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gdamore/tcell"
	"github.com/jamesroutley/fuji/area"
)

func TestThing(t *testing.T) {
	input := "hello"
	screen := tcell.NewSimulationScreen("UTF-8")
	defer screen.Fini()
	e := New(screen, "my_file.txt", strings.NewReader(input))
	e.Draw(area.Area{
		Start: area.Point{0, 0},
		End:   area.Point{50, 50},
	})
	cells, _, _ := screen.GetContents()
	for i, c := range cells {
		// if i < len(input) {
		// 	assert.Equal()
		// }
		fmt.Println(i, c.Runes)
	}

}
