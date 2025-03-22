package enix

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cmd"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/exec"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/help"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/mouse"
	enixTcell "github.com/m-kru/enix/internal/tcell"
	"github.com/m-kru/enix/internal/util"
	"github.com/m-kru/enix/internal/view"

	"github.com/gdamore/tcell/v2"
)

type PromptState int

const (
	Inactive PromptState = iota
	InText
	InShadow
	InCmdMenu
	InPathMenu
	TabReloadQuestion
)

var Prompt prompt
var PromptMenu *menu

// Prompt represents command line prompt.
type prompt struct {
	// History of executed commands.
	History    []string
	HistoryIdx int

	Line   *line.Line
	Cursor *cursor.Cursor
	View   view.View

	ShadowText string

	State PromptState

	PathDir string // Used for menu path
}

func (p *prompt) Clear() {
	frame := PromptFrame

	for x := range frame.Width {
		frame.SetContent(x, 0, ' ', cfg.Style.Default)
	}

	if PromptMenu != nil {
		PromptMenu = nil
	}

	p.State = Inactive
}

func (p *prompt) ShowError(msg string) {
	frame := PromptFrame

	x := 0
	for _, r := range msg {
		if x == frame.Width {
			break
		}
		frame.SetContent(x, 0, r, cfg.Style.Error)
		x++
	}
	for {
		if x == frame.Width {
			break
		}
		frame.SetContent(x, 0, ' ', cfg.Style.Prompt)
		x++
	}

	p.State = Inactive

	Screen.Show()
}

func (p *prompt) ShowInfo(msg string) {
	frame := PromptFrame

	x := 0
	for _, r := range msg {
		frame.SetContent(x, 0, r, cfg.Style.Default)
		x++
	}
	for {
		if x == frame.Width {
			break
		}
		frame.SetContent(x, 0, ' ', cfg.Style.Default)
		x++
	}

	p.State = Inactive

	Screen.Show()
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

	width, _ := Screen.Size()

	p.View.Line = 1
	p.View.Column = 1
	p.View.Width = width - 2 - 1 // - 1 because of cursor
	p.View.Height = 1
}

func (p *prompt) AskTabReload(frame frame.Frame) {
	p.State = TabReloadQuestion
	str := "file was modified externally, reload y/n?"
	x := 0
	for _, r := range str {
		frame.SetContent(x, 0, r, cfg.Style.Warning)
		x++
	}

	frame.SetContent(x, 0, ' ', cfg.Style.Warning)
	x++
	frame.SetContent(x, 0, ' ', cfg.Style.Warning.Reverse(true))
	x++

	for {
		if x == frame.Width {
			break
		}
		frame.SetContent(x, 0, ' ', cfg.Style.Warning)
		x++
	}

	Screen.Show()
}

func (p *prompt) Render() {
	frame := PromptFrame

	if p.State == Inactive {
		return
	} else if p.State == TabReloadQuestion {
		p.AskTabReload(frame)
		return
	}

	frame.SetContent(0, 0, ':', cfg.Style.Prompt)

	if !p.View.IsVisible(p.Cursor.View()) {
		p.View = p.View.MinAdjust(p.Cursor.View())
	}

	p.Line.Render(1, frame.Line(1, 0), p.View, nil, nil)

	if len(p.ShadowText) > 0 {
		for i, r := range p.ShadowText {
			frame.SetContent(i+1+p.Line.RuneCount(), 0, r, cfg.Style.PromptShadow)
		}
	}

	p.Cursor.Render(frame.Line(1, 0), p.View)

	Screen.Show()
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

func (p *prompt) openPathMenu(path string) {
	var dir string
	var base string

	if strings.HasSuffix(path, string(os.PathSeparator)) {
		dir = path
	} else if path != "" {
		dir = filepath.Dir(path)
		if len(dir) == 1 && dir[0] == '.' {
			dir = ""
		}
		if len(dir) > 1 {
			dir += string(os.PathSeparator)
		}

		base = filepath.Base(path)
	}

	p.PathDir = dir

	names := util.DirEntries(dir, base)
	if len(names) == 0 {
		return
	}

	PromptMenu = newMenu(names, 0, cfg.Style.Menu, cfg.Style.MenuItem)

	path = dir + names[0]
	fields := strings.Fields(p.Line.String())
	if len(fields) > 1 {
		fields = fields[0 : len(fields)-1]
	}
	fields = append(fields, path)
	text := strings.Join(fields, " ")

	p.Line, _ = line.FromString(text)
	p.Cursor = cursor.New(p.Line, 1, len(text))

	p.State = InPathMenu
}

func (p *prompt) closeMenu() {
	PromptMenu = nil
	p.State = InText

	// The focus is still on prompt.
	// However, the status line changes its position one line down,
	// so the tab must be rerendered.
	Render(true)
}

func (p *prompt) HandleBacktab() {
	if p.State == InCmdMenu {
		p.HandleBacktabCmdMenu()
		return
	} else if p.State == InPathMenu {
		p.HandleBacktabPathMenu()
		return
	}
}

func (p *prompt) HandleBacktabCmdMenu() {
	_, text := PromptMenu.Prev()
	p.Line, _ = line.FromString(text)
	p.Cursor = cursor.New(p.Line, 1, len(text))
}

func (p *prompt) HandleBacktabPathMenu() {
	_, name := PromptMenu.Prev()

	path := p.PathDir + name
	fields := strings.Fields(p.Line.String())
	fields = fields[0 : len(fields)-1]
	fields = append(fields, path)
	text := strings.Join(fields, " ")

	p.Line, _ = line.FromString(text)
	p.Cursor = cursor.New(p.Line, 1, len(text))
}

func (p *prompt) HandleTab() {
	if p.State == InCmdMenu {
		p.HandleTabCmdMenu()
		return
	} else if p.State == InPathMenu {
		p.HandleTabPathMenu()
		return
	}

	if p.State == InShadow {
		p.ShadowText = ""
		p.State = InText
	}

	// Check if menu should be opened for command name
	fields := strings.Fields(p.Line.String())
	if len(fields) == 0 || (len(fields) == 1 && p.Cursor.RuneIdx == utf8.RuneCountInString(fields[0])) {
		prefix := ""
		if len(fields) > 0 {
			prefix = fields[0]
		}

		itemNames := help.GetCommandNames(prefix)
		if len(itemNames) == 0 {
			return
		}

		PromptMenu = newMenu(itemNames, 0, cfg.Style.Menu, cfg.Style.MenuItem)
		p.State = InCmdMenu

		text := itemNames[0]
		p.Line, _ = line.FromString(text)
		p.Cursor = cursor.New(p.Line, 1, len(text))

		return
	}

	// Check if menu should be opened for file path
	cmd := fields[0]
	if help.IsPathCmd(cmd) {
		if p.Cursor.RuneIdx == len(cmd) || p.Cursor.RuneIdx != p.Line.RuneCount() {
			return
		}

		path := ""
		if len(fields) > 1 {
			path = fields[len(fields)-1]
		}

		p.openPathMenu(path)
	}

}

func (p *prompt) HandleTabCmdMenu() {
	_, text := PromptMenu.Next()
	p.Line, _ = line.FromString(text)
	p.Cursor = cursor.New(p.Line, 1, len(text))
}

func (p *prompt) HandleTabPathMenu() {
	_, name := PromptMenu.Next()

	path := p.PathDir + name
	fields := strings.Fields(p.Line.String())
	fields = fields[0 : len(fields)-1]
	fields = append(fields, path)
	text := strings.Join(fields, " ")

	p.Line, _ = line.FromString(text)
	p.Cursor = cursor.New(p.Line, 1, len(text))
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

func (p *prompt) RxMouseEvent(ev mouse.Event) {
	var text string

	if PromptMenuFrame.Within(ev.X(), ev.Y()) {
		_, text = PromptMenu.RxMouseEvent(ev)
	}

	if p.State == InCmdMenu {
		p.Line, _ = line.FromString(text)
		p.Cursor = cursor.New(p.Line, 1, len(text))
	}
}

func (p *prompt) RxTcellEvent(ev tcell.Event) TcellEventReceiver {
	if p.State == TabReloadQuestion {
		return p.rxTcellEventTabReloadQuestion(ev)
	}

	switch ev := ev.(type) {
	case *tcell.EventResize:
		Resize()
	case *tcell.EventKey:
		// Code responsible for catching events related to menu handling
		keyName := enixTcell.EventKeyName(ev)
		if keyName == "Tab" {
			p.HandleTab()
			break
		} else if keyName == "Backtab" {
			p.HandleBacktab()
			break
		} else if p.State == InCmdMenu || p.State == InPathMenu {
			p.closeMenu()
		}

		cmd, err := cfg.KeysPrompt.ToCmd(ev)
		if err != nil {
			p.ShowError(fmt.Sprintf("%v", err))
			return nil
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
			return nil
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

	return p
}

func (p *prompt) rxTcellEventTabReloadQuestion(ev tcell.Event) TcellEventReceiver {
	var r rune

	switch ev := ev.(type) {
	case *tcell.EventResize:
		Resize()
		p.AskTabReload(PromptFrame)
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
	tab := CurrentTab
	if r == 'y' {
		err = tab.Reload()
	}

	tab.ModTime = util.FileModTime(tab.Path)
	p.State = InText
	p.Clear()

	if err != nil {
		p.ShowError(err.Error())
	}

	return nil
}

// Exec executes command.
func (p *prompt) Exec() TcellEventReceiver {
	c, err := cmd.Parse(strings.TrimSpace(p.Line.String()))
	if err != nil {
		p.ShowError(fmt.Sprintf("%v", err))
		return nil
	}

	var info string
	tab := CurrentTab
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
			err = fmt.Errorf("%s", strings.Join(c.Args, " "))
		case "exec-info":
			info = strings.Join(c.Args, " ")
		case "change":
			err = exec.Change(c.Args, tab)
		case "cursor-count":
			info = fmt.Sprintf("%d", len(tab.Cursors))
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
				CurrentTab = tab
				TabBar.Update()
			}
			updateView = false
		case "esc":
			updateView, err = exec.Esc(c.Args, tab)
		case "ft", "filetype":
			err = exec.Filetype(c.Args, tab)
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
				CurrentTab = tab
				TabBar.Update()
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
				CurrentTab = knTab
				TabBar.Update()
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
		case "path":
			err = exec.Path(c.Args, tab)
		case "paste":
			err = exec.Paste(c.Args, tab)
		case "paste-before":
			err = exec.PasteBefore(c.Args, tab)
		case "right":
			err = exec.Right(c.Args, tab)
		case "rune":
			err = exec.Rune(c.Args, tab)
		case "quit", "q":
			tab, err = exec.Quit(c.Args, tab, false)
			if err == nil && tab == nil {
				Quit = true
				return nil
			} else if tab != nil {
				Tabs = tab.First()
				CurrentTab = tab
				TabBar.Update()
			}
			updateView = false
		case "quit!", "q!":
			tab, _ = exec.Quit(c.Args, tab, true)
			if tab == nil {
				Quit = true
				return nil
			} else {
				Tabs = tab.First()
				CurrentTab = tab
				TabBar.Update()
			}
			updateView = false
		case "redo":
			err = exec.Redo(c.Args, tab)
		case "s", "save":
			info, err = exec.Save(c.Args, tab, cfg.Cfg.TrimOnSave)
			updateView = false
		case "search":
			err = exec.Search(c.Args, tab)
			updateView = false
		case "sel-all":
			err = exec.SelAll(c.Args, tab)
		case "sel-count":
			info = fmt.Sprintf("%d", len(tab.Selections))
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
			CurrentTab, err = exec.SelToTab(c.Args, tab)
			TabBar.Update()
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
			err = exec.Suspend(c.Args, Screen)
			InitScreen()
		case "tab":
			err = exec.Tab(c.Args, tab)
		case "tab-count":
			info = fmt.Sprintf("%d", Tabs.Count())
		case "tn", "tab-next":
			CurrentTab, err = exec.TabNext(c.Args, CurrentTab)
			TabBar.Update()
		case "tp", "tab-prev":
			CurrentTab, err = exec.TabPrev(c.Args, CurrentTab)
			TabBar.Update()
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
			err = fmt.Errorf(
				"invalid or unimplemented command '%s', if unimplemented report on https://github.com/m-kru/enix",
				c.Name,
			)
		}

		if err != nil {
			break
		}
	}

	if err != nil {
		p.ShowError(fmt.Sprintf("%v", err))
		return nil
	}

	if updateView {
		tab.UpdateView()
	}

	if info != "" {
		p.ShowInfo(info)
	} else {
		p.Clear()
	}

	return nil
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
