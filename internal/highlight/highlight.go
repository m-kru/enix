package highlight

import "github.com/gdamore/tcell/v2"

type Highlight struct {
	Line     int
	StartIdx int
	EndIdx   int
	Style    tcell.Style
}
