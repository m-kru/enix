package tab

import (
	"fmt"
	"strings"

	"github.com/m-kru/enix/internal/clip"
)

func (tab *Tab) Yank() string {
	if len(tab.Cursors) > 0 {
		return tab.yankCursors()
	} else {
		return tab.yankSelections()
	}
}

func (tab *Tab) yankCursors() string {
	b := strings.Builder{}

	for _, c := range tab.Cursors {
		b.WriteString(c.Line.String())
		b.WriteString(tab.Newline)
	}

	clip.Write(b.String())

	if len(tab.Cursors) == 1 {
		return "yanked line"
	} else {
		return fmt.Sprintf("yanked %d lines", len(tab.Cursors))
	}
}

func (tab *Tab) yankSelections() string {
	b := strings.Builder{}

	for i, s := range tab.Selections {
		b.WriteString(s.ToString())
		if i < len(tab.Selections)-1 {
			b.WriteString(tab.Newline)
		}
	}

	clip.Write(b.String())

	if len(tab.Selections) == 1 {
		return "yanked selection"
	} else {
		return fmt.Sprintf("yanked %d selections", len(tab.Selections))
	}
}
