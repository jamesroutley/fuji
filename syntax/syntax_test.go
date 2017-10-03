package syntax

import (
	"fmt"
	"testing"

	"github.com/gdamore/tcell"
	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	testCases := []struct {
		styledRunes []StyledRune
		expected    [][]StyledRune
		separator   rune
	}{
		{
			styledRunes: []StyledRune{
				StyledRune{Rune: 'a', Style: tcell.StyleDefault},
				StyledRune{Rune: 'b', Style: tcell.StyleDefault},
				StyledRune{Rune: 'c', Style: tcell.StyleDefault},
				StyledRune{Rune: '\n', Style: tcell.StyleDefault},
				StyledRune{Rune: 'd', Style: tcell.StyleDefault},
				StyledRune{Rune: 'e', Style: tcell.StyleDefault},
			},
			expected: [][]StyledRune{
				[]StyledRune{
					StyledRune{Rune: 'a', Style: tcell.StyleDefault},
					StyledRune{Rune: 'b', Style: tcell.StyleDefault},
					StyledRune{Rune: 'c', Style: tcell.StyleDefault},
				},
				[]StyledRune{
					StyledRune{Rune: 'd', Style: tcell.StyleDefault},
					StyledRune{Rune: 'e', Style: tcell.StyleDefault},
				},
			},
			separator: '\n',
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test: %d", i), func(t *testing.T) {
			result := split(tc.styledRunes, tc.separator)
			assert.Equal(t, tc.expected, result)
		})
	}
}
