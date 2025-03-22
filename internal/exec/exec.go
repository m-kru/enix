package exec

import (
	"fmt"

	"github.com/m-kru/enix/internal/cmd"
	"github.com/m-kru/enix/internal/tab"
)

func Exec(c cmd.Command, tab *tab.Tab) error {
	var err error

	for range c.RepCount {
		switch c.Name {
		case "add-cursor":
			err = AddCursor(c.Args, tab)
		case "align":
			err = Align(c.Args, tab)
		case "backspace":
			err = Backspace(c.Args, tab)
		case "change":
			err = Change(c.Args, tab)
		case "del":
			err = Del(c.Args, tab)
		case "down":
			err = Down(c.Args, tab)
		case "esc":
			_, err = Esc(c.Args, tab)
		case "go":
			err = Go(c.Args, tab)
		case "join":
			err = Join(c.Args, tab)
		case "left":
			err = Left(c.Args, tab)
		case "line-down":
			err = LineDown(c.Args, tab)
		case "line-end":
			err = LineEnd(c.Args, tab)
		case "line-up":
			err = LineUp(c.Args, tab)
		case "mark":
			_, err = Mark(c.Args, tab)
		case "newline":
			err = Newline(c.Args, tab)
		case "paste-before":
			err = PasteBefore(c.Args, tab)
		case "prev-word-start":
			err = PrevWordStart(c.Args, tab)
		case "right":
			err = Right(c.Args, tab)
		case "rune":
			err = Rune(c.Args, tab)
		case "save":
			_, err = Save(c.Args, tab, false)
		case "sel-line":
			err = SelLine(c.Args, tab)
		case "sel-right":
			err = SelRight(c.Args, tab)
		case "sel-switch-cursor":
			err = SelSwitchCursor(c.Args, tab)
		case "sel-word-end":
			err = SelWordEnd(c.Args, tab)
		case "space":
			err = Space(c.Args, tab)
		case "spawn-down":
			err = SpawnDown(c.Args, tab)
		case "spawn-up":
			err = SpawnUp(c.Args, tab)
		case "tab":
			err = Tab(c.Args, tab)
		case "trim":
			err = Trim(c.Args, tab)
		case "undo":
			err = Undo(c.Args, tab)
		case "up":
			err = Up(c.Args, tab)
		case "word-end":
			err = WordEnd(c.Args, tab)
		case "word-start":
			err = WordStart(c.Args, tab)
		case "yank":
			_, err = Yank(c.Args, tab)
		default:
			err = fmt.Errorf("invalid or unimplemented command '%s'", c.Name)
		}

		if err != nil {
			return err
		}
	}

	return nil
}
