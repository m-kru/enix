package enix

import (
	"fmt"
	"strings"

	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cmd"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/exec"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/view"

	"github.com/gdamore/tcell/v2"
)

type PromptState int

const (
	InText PromptState = iota
	InShadow
	InError
)

// Prompt represents command line prompt.
type Prompt struct {
	Screen tcell.Screen
	Frame  frame.Frame

	// History of executed commands.
	History    []string
	HistoryIdx int

	Window *Window

	Line   *line.Line
	Cursor *cursor.Cursor
	View   view.View

	ShadowText string

	State PromptState
}

func (p *Prompt) Clear() {
	for x := range p.Frame.Width {
		p.Frame.SetContent(x, 0, ' ', cfg.Colors.Default)
	}
	p.Screen.HideCursor()
	p.Screen.Show()
}

func (p *Prompt) ShowError(msg string) {
	x := 0
	for _, r := range msg {
		if x == p.Frame.Width {
			break
		}
		p.Frame.SetContent(x, 0, r, cfg.Colors.Error)
		x++
	}
	for {
		if x == p.Frame.Width {
			break
		}
		p.Frame.SetContent(x, 0, ' ', cfg.Colors.Prompt)
		x++
	}
	p.Screen.Show()
	p.State = InError
}

func (p *Prompt) ShowInfo(msg string) {
	x := 0
	for _, r := range msg {
		p.Frame.SetContent(x, 0, r, cfg.Colors.Default)
		x++
	}
	for {
		if x == p.Frame.Width {
			break
		}
		p.Frame.SetContent(x, 0, ' ', cfg.Colors.Default)
		x++
	}
	p.Screen.Show()
	p.State = InError
}

// Currently assume text + shadow text always fits screen width.
func (p *Prompt) Activate(text, shadowText string) {
	p.HistoryIdx = len(p.History)

	if text == "" && shadowText == "" && len(p.History) > 0 {
		p.HistoryIdx--
		shadowText = p.History[p.HistoryIdx]
	}

	p.Line, _ = line.FromString(text)

	p.Cursor = cursor.New(p.Line, 1, len(text))

	p.ShadowText = shadowText

	p.State = InText
	if len(shadowText) > 0 {
		p.State = InShadow
	}

	p.View.Line = 1
	p.View.Column = 1
	p.View.Width = p.Frame.Width - 2 - 1 // - 1 because of cursor
	p.View.Height = 1

	p.Render()
}

func (p *Prompt) Render() {
	p.Frame.SetContent(0, 0, ':', cfg.Colors.Prompt)

	if !p.View.IsVisible(p.Cursor.View()) {
		p.View = p.View.MinAdjust(p.Cursor.View())
	}

	p.Line.Render(1, p.Frame.Line(1, 0), p.View, nil, nil)

	if len(p.ShadowText) > 0 {
		for i, r := range p.ShadowText {
			p.Frame.SetContent(i+1+p.Line.RuneCount(), 0, r, cfg.Colors.PromptShadow)
		}
	}

	p.Cursor.Render(p.Frame.Line(1, 0), p.View)

	p.Screen.Show()
}

func (p *Prompt) Backspace() {
	switch p.State {
	case InShadow:
		p.ShadowText = ""
		p.State = InText
	case InText:
		p.Cursor.Backspace()
	}
}

func (p *Prompt) Delete() {
	switch p.State {
	case InShadow:
		p.ShadowText = ""
		p.State = InText
	case InText:
		p.Cursor.Delete()
	}
}

func (p *Prompt) Down() {
	if p.HistoryIdx == len(p.History) {
		return
	}

	p.HistoryIdx++
	p.Line, _ = line.FromString("")
	p.ShadowText = p.History[p.HistoryIdx]
	p.State = InShadow
	p.Cursor.RuneIdx = 0
}

func (p *Prompt) Up() {
	if p.HistoryIdx == 0 {
		return
	}

	p.HistoryIdx--
	p.Line, _ = line.FromString("")
	p.ShadowText = p.History[p.HistoryIdx]
	p.State = InShadow
	p.Cursor.RuneIdx = 0
}

func (p *Prompt) Left() {
	switch p.State {
	case InShadow:
		p.Line.Append([]byte(p.ShadowText))
		p.ShadowText = ""
		p.State = InText
	case InText:
		p.Cursor.Left()
	}
}

func (p *Prompt) Right() {
	switch p.State {
	case InShadow:
		p.Line.Append([]byte(p.ShadowText))
		p.Cursor.Line = p.Line
		p.Cursor.RuneIdx += len(p.ShadowText)
		p.ShadowText = ""
		p.State = InText
	case InText:
		p.Cursor.Right()
	}
}

func (p *Prompt) PrevWordStart() {
	switch p.State {
	case InShadow:
		p.ShadowText = ""
		p.State = InText
		p.Cursor.PrevWordStart()
	case InText:
		p.Cursor.PrevWordStart()
	}
}

func (p *Prompt) WordEnd() {
	switch p.State {
	case InShadow:
		p.Line.Append([]byte(p.ShadowText))
		p.ShadowText = ""
		p.State = InText
		p.Cursor.WordEnd()
	case InText:
		p.Cursor.WordEnd()
	}
}

func (p *Prompt) LineStart() {
	switch p.State {
	case InShadow:
		p.ShadowText = ""
		p.State = InText
		p.Cursor.LineStart()
	case InText:
		p.Cursor.LineStart()
	}
}

func (p *Prompt) Enter() TcellEventReceiver {
	if p.State == InShadow {
		p.Line.Append([]byte(p.ShadowText))
		p.ShadowText = ""
		p.State = InText
	}

	if len(p.History) == cap(p.History) {
		p.History = p.History[1:]
	}
	p.History = append(p.History, p.Line.String())

	return p.Exec()
}

func (p *Prompt) HandleRune(r rune) {
	switch p.State {
	case InShadow:
		p.ShadowText = ""
		p.State = InText
		p.Cursor.InsertRune(r)
	case InText:
		p.Cursor.InsertRune(r)
	}
}

func (p *Prompt) RxTcellEvent(ev tcell.Event) TcellEventReceiver {
	switch ev := ev.(type) {
	case *tcell.EventResize:
		p.Window.Resize()
		p.Window.Render()
	case *tcell.EventKey:
		cmd, err := cfg.PromptKeys.ToCmd(ev)
		if err != nil {
			p.ShowError(fmt.Sprintf("%v", err))
			return p.Window
		}

		switch cmd.Name {
		case "backspace":
			p.Backspace()
		case "del":
			p.Delete()
		case "down":
			p.Down()
		case "enter":
			return p.Enter()
		case "esc", "quit":
			p.Clear()
			return p.Window
		case "left":
			p.Left()
		case "line-start":
			p.LineStart()
		case "right":
			p.Right()
		case "prev-word-start":
			p.PrevWordStart()
		case "up":
			p.Up()
		case "word-end":
			p.WordEnd()
		default:
			switch ev.Key() {
			case tcell.KeyRune:
				p.HandleRune(ev.Rune())
			}
		}
	}

	p.Render()

	return p
}

// Exec executes command.
func (p *Prompt) Exec() TcellEventReceiver {
	c, err := cmd.Parse(strings.TrimSpace(p.Line.String()))
	if err != nil {
		p.ShowError(fmt.Sprintf("%v", err))
		return p.Window
	}

	var info string
	tab := p.Window.CurrentTab
	updateView := true

	for range c.RepCount {
		switch c.Name {
		case "":
		case "add-cursor":
			err = exec.AddCursor(c.Args, tab)
			// Do nothing
		case "align":
			err = exec.Align(c.Args, tab)
		case "backspace":
			err = exec.Backspace(c.Args, tab)
		case "exec":
			p.Activate("", "")
			return p
		case "exec-error":
			p.ShowError(strings.Join(c.Args, " "))
			return p.Window
		case "exec-info":
			p.ShowInfo(strings.Join(c.Args, " "))
			return p.Window
		case "change":
			err = exec.Change(c.Args, tab)
		case "config-dir":
			info, err = exec.ConfigDir(c.Args)
		case "cursor-count":
			p.ShowInfo(fmt.Sprintf("%d", len(tab.Cursors)))
			return p.Window
		case "cut":
			err = exec.Cut(c.Args, tab)
		case "del":
			err = exec.Del(c.Args, tab)
		case "down":
			err = exec.Down(c.Args, tab)
		case "dump-cursor":
			info, err = exec.DumpCursor(c.Args, tab)
		case "esc":
			err = exec.Esc(c.Args, tab)
		case "find-next":
			err = exec.FindNext(c.Args, tab)
		case "find-prev":
			err = exec.FindPrev(c.Args, tab)
		case "find-sel-next":
			err = exec.FindSelNext(c.Args, tab)
		case "find-sel-prev":
			err = exec.FindSelPrev(c.Args, tab)
		case "g", "go":
			err = exec.Go(c.Args, tab)
		case "h", "help":
			tab, err = exec.Help(c.Args, tab)
			if err == nil {
				p.Window.CurrentTab = tab
			}
			updateView = false
		case "insert-line-above":
			err = exec.InsertLineAbove(c.Args, tab)
		case "insert-line-below":
			err = exec.InsertLineBelow(c.Args, tab)
		case "join":
			err = exec.Join(c.Args, tab)
		case "key-name":
			knTab, err := exec.KeyName(c.Args, tab)
			if err == nil {
				p.Window.CurrentTab = knTab
			}
		case "left":
			err = exec.Left(c.Args, tab)
		case "line-down":
			err = exec.LineDown(c.Args, tab)
		case "line-end":
			err = exec.LineEnd(c.Args, tab)
		case "line-start":
			err = exec.LineStart(c.Args, tab)
		case "line-up":
			err = exec.LineUp(c.Args, tab)
		case "newline":
			err = exec.Newline(c.Args, tab)
		case "m", "mark":
			info, err = exec.Mark(c.Args, tab)
		case "o", "open":
			tab, err = exec.Open(c.Args, tab)
			if err == nil {
				p.Window.CurrentTab = tab
			}
			updateView = false
		case "pwd":
			info, err = exec.Pwd(c.Args)
		case "paste":
			err = exec.Paste(c.Args, tab)
		case "right":
			err = exec.Right(c.Args, tab)
		case "rune":
			err = exec.Rune(c.Args, tab)
		case "quit", "q":
			tab, err = exec.Quit(c.Args, tab, false)
			if err == nil && tab == nil {
				return nil
			} else if tab != nil {
				p.Window.Tabs = tab.First()
				p.Window.CurrentTab = tab
			}
			updateView = false
		case "quit!", "q!":
			tab, _ = exec.Quit(c.Args, tab, true)
			if tab == nil {
				return nil
			} else {
				p.Window.Tabs = tab.First()
				p.Window.CurrentTab = tab
			}
			updateView = false
		case "redo":
			err = exec.Redo(c.Args, tab)
		case "save":
			info, err = exec.Save(c.Args, tab, cfg.Cfg.TrimOnSave)
		case "search":
			err = exec.Search(c.Args, tab)
		case "sel-count":
			p.ShowInfo(fmt.Sprintf("%d", len(tab.Selections)))
			return p.Window
		case "sel-down":
			err = exec.SelDown(c.Args, tab)
		case "sel-left":
			err = exec.SelLeft(c.Args, tab)
		case "sel-line":
			err = exec.SelLine(c.Args, tab)
		case "sel-prev-word-start":
			err = exec.SelPrevWordStart(c.Args, tab)
		case "sel-right":
			err = exec.SelRight(c.Args, tab)
		case "sel-to-tab":
			p.Window.CurrentTab, err = exec.SelToTab(c.Args, tab)
		case "sel-up":
			err = exec.SelUp(c.Args, tab)
		case "sel-word-end":
			err = exec.SelWordEnd(c.Args, tab)
		case "space":
			err = exec.Space(c.Args, tab)
		case "spawn-down":
			err = exec.SpawnDown(c.Args, tab)
		case "spawn-up":
			err = exec.SpawnUp(c.Args, tab)
		case "suspend":
			err = exec.Suspend(c.Args, p.Window.Screen)
		case "tab":
			err = exec.Tab(c.Args, tab)
		case "tab-count":
			p.ShowInfo(fmt.Sprintf("%d", p.Window.Tabs.Count()))
			return p.Window
		case "tn", "tab-next":
			p.Window.CurrentTab, err = exec.TabNext(c.Args, tab)
		case "tp", "tab-prev":
			p.Window.CurrentTab, err = exec.TabPrev(c.Args, tab)
		case "trim":
			err = exec.Trim(c.Args, tab)
		case "undo":
			err = exec.Undo(c.Args, tab)
		case "up":
			err = exec.Up(c.Args, tab)
		case "view-center":
			err = exec.ViewCenter(c.Args, tab)
		case "view-down":
			err = exec.ViewDown(c.Args, tab)
			updateView = false
		case "view-left":
			err = exec.ViewLeft(c.Args, tab)
			updateView = false
		case "view-right":
			err = exec.ViewRight(c.Args, tab)
			updateView = false
		case "view-up":
			err = exec.ViewUp(c.Args, tab)
			updateView = false
		case "prev-word-start":
			err = exec.PrevWordStart(c.Args, tab)
		case "word-end":
			err = exec.WordEnd(c.Args, tab)
		case "word-start":
			err = exec.WordStart(c.Args, tab)
		case "yank":
			err = exec.Yank(c.Args, tab)
		default:
			p.ShowError(
				fmt.Sprintf(
					"invalid or unimplemented command '%s', if unimplemented report on https://github.com/m-kru/enix",
					c.Name,
				),
			)
			return p.Window
		}

		if err != nil {
			break
		}
	}

	if err != nil {
		p.ShowError(fmt.Sprintf("%v", err))
		return p.Window
	}

	if updateView {
		tab.UpdateView()
	}

	if info != "" {
		p.ShowInfo(info)
	} else {
		p.Clear()
	}

	return p.Window
}
