package syntax

import (
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/gdamore/tcell"
)

var style = styles.Get("dracula")

// StyledRune reprenents a rune and its syntax highlighting style
type StyledRune struct {
	Rune  rune
	Style tcell.Style
}

// Background returns the background style
func Background() tcell.Style {
	styleEntry := style.Get(chroma.Background)
	return chromaStyleToTcellStyle(styleEntry)
}

// Highlight highlights a string
func Highlight(filename, text string) [][]StyledRune {
	lexer := lexers.Match(filename)
	if lexer == nil {
		lexer = lexers.Fallback
	}
	style := styles.Get("dracula")
	if style == nil {
		style = styles.Fallback
	}
	iterator, err := lexer.Tokenise(nil, text)
	if err != nil {
		// TODO: this probably shouldn't panic - maybe return a
		// non-highlighted string?
		panic(err)
	}

	tokens := iterator.Tokens()
	var styledRunes []StyledRune
	for _, token := range tokens {
		styledRunes = append(styledRunes, tokenToStyledRunes(token, style)...)
	}

	return split(styledRunes, '\n')
}

func split(styledRunes []StyledRune, sep rune) [][]StyledRune {
	var word []StyledRune
	var result [][]StyledRune
	for _, sr := range styledRunes {
		if sr.Rune == sep {
			result = append(result, word)
			word = []StyledRune{}
		} else {
			word = append(word, sr)
		}
	}
	result = append(result, word)

	return result
}

func tokenToStyledRunes(t *chroma.Token, s *chroma.Style) []StyledRune {
	styleEntry := s.Get(t.Type)
	style := chromaStyleToTcellStyle(styleEntry)
	// hacky
	// TODO: getting the len of t.Value may not work for unicode
	runes := make([]StyledRune, len(t.Value))
	for i, r := range t.Value {
		runes[i] = StyledRune{
			Rune:  r,
			Style: style,
		}
	}
	return runes
}

func chromaStyleToTcellStyle(se chroma.StyleEntry) (s tcell.Style) {
	s = tcell.StyleDefault
	s = s.Background(chromaToTcellColour(se.Background))
	s = s.Foreground(chromaToTcellColour(se.Colour))
	return
}

func chromaToTcellColour(c chroma.Colour) tcell.Color {
	return tcell.NewRGBColor(
		int32(c.Red()),
		int32(c.Green()),
		int32(c.Blue()),
	)
}
