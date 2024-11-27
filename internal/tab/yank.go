package tab

import (
	"strings"

	"github.com/m-kru/enix/internal/clip"
)

func (tab *Tab) Yank() {
	if len(tab.Cursors) > 0 {
		tab.yankCursors()
	} else {
		tab.yankSelections()
	}
}

func (tab *Tab) yankCursors() {
	b := strings.Builder{}

	for _, c := range tab.Cursors {
		b.WriteString(c.Line.String())
		b.WriteString(tab.Newline)
	}

	clip.Write(b.String())
}

func (tab *Tab) yankSelections() {
	b := strings.Builder{}

	for i, s := range tab.Selections {
		b.WriteString(s.ToString())
		if i < len(tab.Selections)-1 {
			b.WriteString(tab.Newline)
		}
	}

	clip.Write(b.String())
}
