package enix

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cmd"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/exec"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/util"
	"github.com/m-kru/enix/internal/view"

	"github.com/gdamore/tcell/v2"
)

type PromptState int

const (
	InText PromptState = iota
	InShadow
	InError
	TabReloadQuestion
)

var Prompt prompt

// Prompt represents command line prompt.
type prompt struct {
	Screen tcell.Screen
	Frame  frame.Frame

	// History of executed commands.
	History    []string
	HistoryIdx int

	Line   *line.Line
	Cursor *cursor.Cursor
	View   view.View

	ShadowText string

	State PromptState
}

func (p *prompt) Clear() {
	for x := range p.Frame.Width {
		p.Frame.SetContent(x, 0, ' ', cfg.Colors.Default)
	}
	p.Screen.HideCursor()
	p.Screen.Show()
}

func (p *prompt) ShowError(msg string) {
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

func (p *prompt) ShowInfo(msg string) {
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
func (p *prompt) Activate(text, shadowText string) {
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

func (p *prompt) AskTabReload() {
	p.State = TabReloadQuestion
	str := "file was modified externally, reload y/n?"
	x := 0
	for _, r := range str {
		p.Frame.SetContent(x, 0, r, cfg.Colors.Warning)
		x++
	}

	p.Frame.SetContent(x, 0, ' ', cfg.Colors.Warning)
	x++
	p.Frame.SetContent(x, 0, ' ', cfg.Colors.Warning.Reverse(true))
	x++

	for {
		if x == p.Frame.Width {
			break
		}
		p.Frame.SetContent(x, 0, ' ', cfg.Colors.Warning)
		x++
	}

	p.Screen.Show()
}

func (p *prompt) Render() {
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

func (p *prompt) Backspace() {
	switch p.State {
	case InShadow:
		p.ShadowText = ""
		p.State = InText
	case InText:
		p.Cursor.Backspace()
	}
}

func (p *prompt) Delete() {
	switch p.State {
	case InShadow:
		p.ShadowText = ""
		p.State = InText
	case InText:
		p.Cursor.Delete()
	}
}

func (p *prompt) Down() {
	if p.HistoryIdx == len(p.History) {
		return
	}

	p.HistoryIdx++
	p.Line, _ = line.FromString("")
	p.ShadowText = p.History[p.HistoryIdx]
	p.State = InShadow
	p.Cursor.RuneIdx = 0
}

func (p *prompt) Up() {
	if p.HistoryIdx == 0 {
		return
	}

	p.HistoryIdx--
	p.Line, _ = line.FromString("")
	p.ShadowText = p.History[p.HistoryIdx]
	p.State = InShadow
	p.Cursor.RuneIdx = 0
}

func (p *prompt) Left() {
	switch p.State {
	case InShadow:
		p.Line.Append([]byte(p.ShadowText))
		p.ShadowText = ""
		p.State = InText
	case InText:
		p.Cursor.Left()
	}
}

func (p *prompt) Right() {
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

func (p *prompt) PrevWordStart() {
	switch p.State {
	case InShadow:
		p.ShadowText = ""
		p.State = InText
		p.Cursor.PrevWordStart()
	case InText:
		p.Cursor.PrevWordStart()
	}
}

func (p *prompt) WordEnd() {
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

func (p *prompt) LineStart() {
	switch p.State {
	case InShadow:
		p.ShadowText = ""
		p.State = InText
		p.Cursor.LineStart()
	case InText:
		p.Cursor.LineStart()
	}
}

func (p *prompt) Enter() TcellEventReceiver {
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

func (p *prompt) HandleRune(r rune) {
	switch p.State {
	case InShadow:
		p.ShadowText = ""
		p.State = InText
		p.Cursor.InsertRune(r)
	case InText:
		p.Cursor.InsertRune(r)
	}
}

func (p *prompt) RxTcellEvent(ev tcell.Event) TcellEventReceiver {
	if p.State == TabReloadQuestion {
		return p.rxTcellEventTabReloadQuestion(ev)
	}

	switch ev := ev.(type) {
	case *tcell.EventResize:
		Window.Resize()
		Window.Render()
	case *tcell.EventKey:
		cmd, err := cfg.PromptKeys.ToCmd(ev)
		if err != nil {
			p.ShowError(fmt.Sprintf("%v", err))
			return &Window
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
			return &Window
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

func (p *prompt) rxTcellEventTabReloadQuestion(ev tcell.Event) TcellEventReceiver {
	var r rune

	switch ev := ev.(type) {
	case *tcell.EventResize:
		Window.Resize()
		Window.Render()
		p.AskTabReload()
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyRune:
			r = ev.Rune()
		}
	}

	if r != 'n' && r != 'y' {
		return p
	}

	var err error
	tab := Window.CurrentTab
	if r == 'y' {
		err = tab.Reload()
	}

	tab.ModTime = util.FileModTime(tab.Path)
	p.State = InText
	p.Clear()

	if err != nil {
		p.ShowError(err.Error())
	}

	return &Window
}

// Exec executes command.
func (p *prompt) Exec() TcellEventReceiver {
	c, err := cmd.Parse(strings.TrimSpace(p.Line.String()))
	if err != nil {
		p.ShowError(fmt.Sprintf("%v", err))
		return &Window
	}

	var info string
	tab := Window.CurrentTab
	updateView := true

	for range c.RepCount {
		switch c.Name {
		case "":
			// Do nothing
		case "autosave":
			err = execAutosave(c.Args)
		case "add-cursor":
			err = exec.AddCursor(c.Args, tab)
		case "a", "align":
			err = exec.Align(c.Args, tab)
		case "backspace":
			err = exec.Backspace(c.Args, tab)
		case "exec":
			p.Activate("", "")
			return p
		case "exec-error":
			p.ShowError(strings.Join(c.Args, " "))
			return &Window
		case "exec-info":
			p.ShowInfo(strings.Join(c.Args, " "))
			return &Window
		case "change":
			err = exec.Change(c.Args, tab)
		case "config-dir":
			info, err = exec.ConfigDir(c.Args)
		case "cursor-count":
			p.ShowInfo(fmt.Sprintf("%d", len(tab.Cursors)))
			return &Window
		case "cut":
			info, err = exec.Cut(c.Args, tab)
		case "del":
			err = exec.Del(c.Args, tab)
		case "down":
			err = exec.Down(c.Args, tab)
		case "dump-cursor":
			info, err = exec.DumpCursor(c.Args, tab)
		case "e", "edit":
			tab, err = exec.Edit(c.Args, tab)
			if err == nil {
				Window.CurrentTab = tab
			}
			updateView = false
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
				Window.CurrentTab = tab
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
				Window.CurrentTab = knTab
			}
		case "left":
			err = exec.Left(c.Args, tab)
		case "line-count":
			info, err = exec.LineCount(c.Args, tab)
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
				Window.Tabs = tab.First()
				Window.CurrentTab = tab
			}
			updateView = false
		case "quit!", "q!":
			tab, _ = exec.Quit(c.Args, tab, true)
			if tab == nil {
				return nil
			} else {
				Window.Tabs = tab.First()
				Window.CurrentTab = tab
			}
			updateView = false
		case "redo":
			err = exec.Redo(c.Args, tab)
		case "save":
			info, err = exec.Save(c.Args, tab, cfg.Cfg.TrimOnSave)
			updateView = false
		case "search":
			err = exec.Search(c.Args, tab)
			updateView = false
		case "sel-all":
			err = exec.SelAll(c.Args, tab)
		case "sel-count":
			p.ShowInfo(fmt.Sprintf("%d", len(tab.Selections)))
			return &Window
		case "sel-down":
			err = exec.SelDown(c.Args, tab)
		case "sel-left":
			err = exec.SelLeft(c.Args, tab)
		case "sel-line":
			err = exec.SelLine(c.Args, tab)
		case "sel-line-end":
			err = exec.SelLineEnd(c.Args, tab)
		case "sel-line-start":
			err = exec.SelLineStart(c.Args, tab)
		case "sel-prev-word-start":
			err = exec.SelPrevWordStart(c.Args, tab)
		case "sel-right":
			err = exec.SelRight(c.Args, tab)
		case "sel-tab-end":
			err = exec.SelTabEnd(c.Args, tab)
		case "sel-to-tab":
			Window.CurrentTab, err = exec.SelToTab(c.Args, tab)
		case "sel-switch-cursor":
			err = exec.SelSwitchCursor(c.Args, tab)
		case "sel-up":
			err = exec.SelUp(c.Args, tab)
		case "sel-word-end":
			err = exec.SelWordEnd(c.Args, tab)
		case "space":
			err = exec.Space(c.Args, tab)
		case "sh":
			err = exec.Sh(c.Args, tab)
		case "spawn-down":
			err = exec.SpawnDown(c.Args, tab)
		case "spawn-up":
			err = exec.SpawnUp(c.Args, tab)
		case "suspend":
			err = exec.Suspend(c.Args, Window.Screen)
		case "tab":
			err = exec.Tab(c.Args, tab)
		case "tab-count":
			p.ShowInfo(fmt.Sprintf("%d", Window.Tabs.Count()))
			return &Window
		case "tn", "tab-next":
			Window.CurrentTab, err = exec.TabNext(c.Args, tab)
		case "tp", "tab-prev":
			Window.CurrentTab, err = exec.TabPrev(c.Args, tab)
		case "trim":
			err = exec.Trim(c.Args, tab)
		case "trim-on-save":
			info, err = exec.TrimOnSave(c.Args)
		case "undo":
			err = exec.Undo(c.Args, tab)
		case "up":
			err = exec.Up(c.Args, tab)
		case "vc", "view-center":
			err = exec.ViewCenter(c.Args, tab)
		case "view-down":
			err = exec.ViewDown(c.Args, tab)
			updateView = false
		case "ve", "view-end":
			err = exec.ViewEnd(c.Args, tab)
			updateView = false
		case "view-left":
			err = exec.ViewLeft(c.Args, tab)
			updateView = false
		case "view-right":
			err = exec.ViewRight(c.Args, tab)
			updateView = false
		case "vs", "view-start":
			err = exec.ViewStart(c.Args, tab)
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
			info, err = exec.Yank(c.Args, tab)
		default:
			p.ShowError(
				fmt.Sprintf(
					"invalid or unimplemented command '%s', if unimplemented report on https://github.com/m-kru/enix",
					c.Name,
				),
			)
			return &Window
		}

		if err != nil {
			break
		}
	}

	if err != nil {
		p.ShowError(fmt.Sprintf("%v", err))
		return &Window
	}

	if updateView {
		tab.UpdateView()
	}

	if info != "" {
		p.ShowInfo(info)
	} else {
		p.Clear()
	}

	return &Window
}

// Auto save command modifies autoSaveTicker, and can't be executed by the exec package.
func execAutosave(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("autosave: expected 1 arg, provided %d", len(args))
	}

	period, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("autosave: %v", err)
	}

	if period < 0 {
		return fmt.Errorf("autosave: period value must be natural, current value %d", period)
	}

	if period == 0 {
		autoSaveTicker.Stop()
	} else {
		autoSaveTicker.Reset(time.Duration(period) * time.Second)
	}

	return nil
}
