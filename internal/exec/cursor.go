package exec

import (
	"fmt"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/mark"
	"github.com/m-kru/enix/internal/tab"
	"strconv"
	"strings"
	"unicode"
)

func Down(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"down: provided %d args, expected 0", len(args),
		)
	}

	tab.Down()

	return nil
}

func Left(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"left: provided %d args, expected 0", len(args),
		)
	}

	tab.Left()

	return nil
}

func Right(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"right: provided %d args, expected 0", len(args),
		)
	}

	tab.Right()

	return nil
}

func Up(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"up: provided %d args, expected 0", len(args),
		)
	}

	tab.Up()

	return nil
}

func End(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"end: expected 0 args, provided %d", len(args),
		)
	}

	tab.Cursors = []*cursor.Cursor{
		&cursor.Cursor{
			Line:    tab.Lines.Last(),
			LineNum: tab.LineCount,
		},
	}

	return nil
}

func Go(args []string, tab *tab.Tab) error {
	var err error
	line := 0
	col := 0

	if len(args) == 0 {
		return fmt.Errorf("go: missing at least line number or mark name")
	}

	if len(args) == 1 && !unicode.IsDigit([]rune(args[0])[0]) {
		return goMark(args[0], tab)
	}

	if len(args) == 1 {
		if strings.Contains(args[0], ":") {
			lineStr, colStr, _ := strings.Cut(args[0], ":")
			line, err = strconv.Atoi(lineStr)
			if err != nil {
				return fmt.Errorf("go: %v", err)
			}
			col, err = strconv.Atoi(colStr)
			if err != nil {
				return fmt.Errorf("go: %v", err)
			}
		} else {
			line, err = strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("go: %v", err)
			}
		}
	} else if len(args) == 2 {
		line, err = strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("go: %v", err)
		}
		col, err = strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("go: %v", err)
		}
	} else {
		return fmt.Errorf("go: expected at most 2 args, provided %d", len(args))
	}

	goCmd(line, col, tab)

	return nil
}

func goCmd(lineNum, col int, tab *tab.Tab) {
	if lineNum < 1 {
		lineNum = 1
	}
	if col < 1 {
		col = 1
	}
	if lineNum > tab.LineCount {
		lineNum = tab.LineCount
	}

	line := tab.Lines.Get(lineNum)
	if col > line.RuneCount()+1 {
		col = line.RuneCount() + 1
	}

	tab.Cursors = []*cursor.Cursor{
		&cursor.Cursor{
			Line:    line,
			LineNum: lineNum,
			RuneIdx: col - 1,
			Idx:     col - 1,
		},
	}
}

func goMark(name string, tab *tab.Tab) error {
	m, ok := tab.Marks[name]
	if !ok {
		return fmt.Errorf("go: no '%s' mark", name)
	}

	switch m := m.(type) {
	case *mark.CursorMark:
		tab.Cursors = cursor.Clone(m.Cursors)
	default:
		panic("selection mark unimplemented")
	}

	return nil
}

func PrevWordStart(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"prev-word-start: expected 0 args, provided %d", len(args),
		)
	}

	return prevWordStart(tab)
}

func prevWordStart(tab *tab.Tab) error {
	for _, c := range tab.Cursors {
		c.PrevWordStart()
	}

	tab.Cursors = cursor.Prune(tab.Cursors)

	return nil
}

func WordEnd(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"word-end: expected 0 args, provided %d", len(args),
		)
	}

	return wordEnd(tab)
}

func wordEnd(tab *tab.Tab) error {
	for _, c := range tab.Cursors {
		c.WordEnd()
	}

	tab.Cursors = cursor.Prune(tab.Cursors)

	return nil
}

func WordStart(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"word-start: expected 0 args, provided %d", len(args),
		)
	}

	return wordStart(tab)
}

func wordStart(tab *tab.Tab) error {
	for _, c := range tab.Cursors {
		c.WordStart()
	}

	tab.Cursors = cursor.Prune(tab.Cursors)

	return nil
}

func SpawnDown(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"spawn-down: expected 0 args, provided %d", len(args),
		)
	}

	return spawnDown(tab)
}

func spawnDown(tab *tab.Tab) error {
	newCurs := make([]*cursor.Cursor, 0, len(tab.Cursors))

	for _, c := range tab.Cursors {
		nc := c.SpawnDown()

		if nc == nil {
			continue
		}

		newCurs = append(newCurs, nc)
	}

	if len(newCurs) > 0 {
		tab.Cursors = append(tab.Cursors, newCurs...)
	}

	tab.Cursors = cursor.Prune(tab.Cursors)

	return nil
}

func SpawnUp(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"spawn-up: expected 0 args, provided %d", len(args),
		)
	}

	return spawnUp(tab)
}

func spawnUp(tab *tab.Tab) error {
	newCurs := make([]*cursor.Cursor, 0, len(tab.Cursors))

	for _, c := range tab.Cursors {
		nc := c.SpawnUp()

		if nc == nil {
			continue
		}

		newCurs = append(newCurs, nc)
	}

	if len(newCurs) > 0 {
		tab.Cursors = append(tab.Cursors, newCurs...)
	}

	tab.Cursors = cursor.Prune(tab.Cursors)

	return nil
}

func LineStart(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"line-start: expected 0 args, provided %d", len(args),
		)
	}

	return lineStart(tab)
}

func lineStart(tab *tab.Tab) error {
	for _, c := range tab.Cursors {
		c.LineStart()
	}

	tab.Cursors = cursor.Prune(tab.Cursors)

	return nil
}

func LineEnd(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"line-end: expected 0 args, provided %d", len(args),
		)
	}

	return lineEnd(tab)
}

func lineEnd(tab *tab.Tab) error {
	for _, c := range tab.Cursors {
		c.LineEnd()
	}

	tab.Cursors = cursor.Prune(tab.Cursors)

	return nil
}

func AddCursor(args []string, tab *tab.Tab) error {
	if len(args) < 1 || 2 < len(args) {
		return fmt.Errorf(
			"add-cursor: provided %d args, expected 1 or 2", len(args),
		)
	}

	line, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("add-cursor: %v", err)
	}
	if line < 1 {
		return fmt.Errorf(
			"add-cursor: line number must be positive, current value %d", line,
		)
	}

	col := 1
	if len(args) == 2 {
		col, err = strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("add-cursor: %v", err)
		}
	}
	if col < 1 {
		return fmt.Errorf(
			"add-cursor: column number must be positive, current value %d", col,
		)
	}

	tab.AddCursor(line, col)

	return nil
}

func DumpCursor(args []string, tab *tab.Tab) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf(
			"dump-cursor: provided %d args, expected 1", len(args),
		)
	}

	n, err := strconv.Atoi(args[0])
	if err != nil {
		return "", fmt.Errorf("dump-cursor: %v", err)
	}

	if n < 1 {
		return "", fmt.Errorf(
			"dump-cursor: cursor index must be positive, current value %d", n,
		)
	}

	if n > len(tab.Cursors) {
		return "", fmt.Errorf(
			"dump-cursor: can't get %d cursor, there are %d cursors",
			n, len(tab.Cursors),
		)
	}

	c := tab.Cursors[n-1]

	return fmt.Sprintf("%v", *c), nil
}
