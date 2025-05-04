package enix

import (
	"fmt"
	"log"
	"time"
	"unicode"

	"github.com/m-kru/enix/internal/arg"
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/exec"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/mouse"
	"github.com/m-kru/enix/internal/tab"
	"github.com/m-kru/enix/internal/util"
	"github.com/m-kru/enix/internal/view"

	"github.com/gdamore/tcell/v2"
)

// Global quit flag
var Quit bool

var Screen tcell.Screen
var screenQuitChan chan struct{}
var tcellEventChan chan tcell.Event

// Frames
var TabBarFrame frame.Frame
var TabFrame frame.Frame
var StatusLineFrame frame.Frame
var PromptMenuFrame frame.Frame
var PromptFrame frame.Frame

var TabBar tabBar

var Tabs *tab.Tab // First tab
var CurrentTab *tab.Tab

var autoSaveTicker *time.Ticker

func RxMouseEvent(ev mouse.Event) {
	x := ev.X()
	y := ev.Y()

	if TabBarFrame.Within(x, y) {
		newCurrentTab := TabBar.RxMouseEvent(ev)
		if newCurrentTab != nil {
			CurrentTab = newCurrentTab
		}
		return
	}

	if PromptMenuFrame.Within(x, y) {
		Prompt.RxMouseEvent(ev)
		return
	}

	if !TabFrame.Within(x, y) {
		return
	}

	x, y = TabFrame.ToFramePosition(ev.X(), ev.Y())

	switch ev.(type) {
	case mouse.PrimaryClick:
		CurrentTab.PrimaryClick(x, y)
	case mouse.PrimaryClickAlt:
		CurrentTab.PrimaryClickAlt(x, y)
	case mouse.PrimaryClickCtrl:
		CurrentTab.PrimaryClickCtrl(x, y)
	case mouse.DoublePrimaryClick:
		CurrentTab.DoublePrimaryClick(x, y)
	case mouse.WheelDown:
		for range cfg.Cfg.MouseScrollMultiplier {
			CurrentTab.ViewDown()
		}
	case mouse.WheelUp:
		for range cfg.Cfg.MouseScrollMultiplier {
			CurrentTab.ViewUp()
		}
	case mouse.WheelLeft:
		for range cfg.Cfg.MouseScrollMultiplier {
			CurrentTab.ViewLeft()
		}
	case mouse.WheelRight:
		for range cfg.Cfg.MouseScrollMultiplier {
			CurrentTab.ViewRight()
		}
	}
}

func RxTcellEvent(ev tcell.Event) TcellEventReceiver {
	switch ev := ev.(type) {
	case *tcell.EventResize:
		Resize()
	case *tcell.EventKey:
		return RxTcellEventKey(ev)
	}

	return nil
}

func RxTcellEventKey(ev *tcell.EventKey) TcellEventReceiver {
	var err error
	var info string
	updateView := true
	tab := CurrentTab

	if tab.State != "" {
		cmd := tab.RxEventKey(ev)
		if cmd != "" {
			Prompt.Activate(cmd, "")
			return &Prompt
		}
		return nil
	}

	if ev.Key() == tcell.KeyRune {
		r := ev.Rune()
		if unicode.IsDigit(r) {
			RxDigit(r)
			return nil
		}
	}

	c, err := cfg.Keys.ToCmd(ev)
	if err != nil {
		Prompt.ShowError(fmt.Sprintf("%v", err))
		return nil
	}

	if c.Name == "" {
		return nil
	}

	if tab.RepCount != 0 {
		c.RepCount = tab.RepCount
		tab.RepCount = 0
	}

	// Limit the repetition count for the esc command.
	// If user by accident provides high repetition count,
	// then ignoring it would introduce high latency.
	// It doesn't make sense to run in more than 4 or 5 times.
	// However, to be immune to potential future changes set
	// the limit to 10.
	if c.Name == "esc" {
		if c.RepCount > 10 {
			c.RepCount = 10
		}
	}

	for range c.RepCount {
		switch c.Name {
		case "add-cursor":
			err = exec.AddCursor(c.Args, tab)
		case "a", "align":
			err = exec.Align(c.Args, tab)
		case "backspace":
			err = exec.Backspace(c.Args, tab)
		case "change":
			err = exec.Change(c.Args, tab)
		case "cmd":
			Prompt.Activate("", "")
			return &Prompt
		case "cut":
			info, err = exec.Cut(c.Args, tab)
		case "del":
			err = exec.Del(c.Args, tab)
		case "down":
			err = exec.Down(c.Args, tab)
		case "e", "edit":
			Prompt.Activate("edit ", "")
			return &Prompt
		case "esc":
			updateView, err = exec.Esc(c.Args, tab)
			Prompt.Clear()
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
			Prompt.Activate("help ", "")
			return &Prompt
		case "indent":
			tab.Indent()
		case "insert":
			tab.Insert()
		case "insert-line-above":
			err = exec.InsertLineAbove(c.Args, tab)
		case "insert-line-below":
			err = exec.InsertLineBelow(c.Args, tab)
		case "insert-tab":
			err = exec.InsertTab(c.Args, tab)
		case "join":
			err = exec.Join(c.Args, tab)
		case "left":
			err = exec.Left(c.Args, tab)
		case "line-count":
			info, err = exec.LineCount(c.Args, tab)
		case "line-down":
			err = exec.LineDown(c.Args, tab)
		case "line-end":
			err = exec.LineEnd(c.Args, tab)
		case "line-up":
			err = exec.LineUp(c.Args, tab)
		case "line-start":
			err = exec.LineStart(c.Args, tab)
		case "m", "mark":
			info, err = exec.Mark(c.Args, tab)
		case "mb", "match-bracket":
			info, err = exec.MatchBracket(c.Args, tab)
		case "mc", "match-curly":
			info, err = exec.MatchCurly(c.Args, tab)
		case "mp", "match-paren":
			info, err = exec.MatchParen(c.Args, tab)
		case "newline":
			err = exec.Newline(c.Args, tab)
		case "path":
			Prompt.Activate("path ", tab.GetWord())
			return &Prompt
		case "paste":
			err = exec.Paste(c.Args, tab)
		case "paste-before":
			err = exec.PasteBefore(c.Args, tab)
		case "pwd":
			info, err = exec.Pwd(c.Args)
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
				Tabs = CurrentTab.First()
				TabBar.Update()
			}
			updateView = false
		case "redo":
			err = exec.Redo(c.Args, tab)
		case "replace":
			tab.State = "replace"
		case "right":
			err = exec.Right(c.Args, tab)
		case "s", "save":
			info, err = exec.Save(c.Args, tab, cfg.Cfg.TrimOnSave)
			updateView = false
		case "search":
			Prompt.Activate("search ", tab.GetWord())
			return &Prompt
		case "sel-all":
			err = exec.SelAll(c.Args, tab)
		case "sb", "sel-bracket":
			err = exec.SelBracket(c.Args, tab)
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
		case "sel-switch-cursor":
			err = exec.SelSwitchCursor(c.Args, tab)
		case "sel-tab-end":
			err = exec.SelTabEnd(c.Args, tab)
		case "sel-to-tab":
			CurrentTab, err = exec.SelToTab(c.Args, tab)
			TabBar.Update()
		case "sel-up":
			err = exec.SelUp(c.Args, tab)
		case "sel-word":
			err = exec.SelWord(c.Args, tab)
		case "sel-word-end":
			err = exec.SelWordEnd(c.Args, tab)
		case "sel-word-start":
			err = exec.SelWordStart(c.Args, tab)
		case "sh":
			Prompt.Activate(c.String(), "")
			return &Prompt
		case "space":
			err = exec.Space(c.Args, tab)
		case "spawn-down":
			err = exec.SpawnDown(c.Args, tab)
		case "spawn-up":
			err = exec.SpawnUp(c.Args, tab)
		case "suspend":
			err = exec.Suspend(c.Args, Screen)
			InitScreen()
		case "t", "tab":
			CurrentTab, err = exec.Tab(c.Args, Tabs)
			TabBar.Update()
		case "tn", "tab-next":
			CurrentTab, err = exec.TabNext(c.Args, tab)
			TabBar.Update()
		case "tp", "tab-prev":
			CurrentTab, err = exec.TabPrev(c.Args, tab)
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
		case "view-down-half":
			err = exec.ViewDownHalf(c.Args, tab)
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
		case "view-up-half":
			err = exec.ViewUpHalf(c.Args, tab)
			updateView = false
		case "word-end":
			err = exec.WordEnd(c.Args, tab)
		case "word-start":
			err = exec.WordStart(c.Args, tab)
		case "yank":
			info, err = exec.Yank(c.Args, tab)
		case "prev-word-start":
			err = exec.PrevWordStart(c.Args, tab)
		default:
			err = fmt.Errorf(
				"invalid or unimplemeneted command '%s', if unimplemented report on https://github.com/m-kru/enix",
				c.Name,
			)
			updateView = false
		}

		if err != nil {
			break
		}
	}

	if updateView {
		tab.UpdateView()
	}

	if err != nil {
		Prompt.ShowError(fmt.Sprintf("%v", err))
	} else if info != "" {
		Prompt.ShowInfo(info)
	} else {
		Prompt.Clear()
	}

	return nil
}

func RxDigit(digit rune) {
	d := int(digit - '0')

	t := CurrentTab
	t.RepCount = t.RepCount*10 + d
	if t.RepCount < 0 {
		t.RepCount = 0
	}
}

// Resize handles all the required logic when screen is resized.
func Resize() {
	Screen.Fill(' ', cfg.Style.Default)
	Screen.Sync()
}

// Do not rerender tab if focus is on the prompt.
// This reduces responsiveness in the case of large files.
func Render(renderTab bool) {
	width, height := Screen.Size()

	// Set frames

	TabFrame = frame.Frame{
		Screen: Screen,
		X:      0,
		Y:      0,
		Width:  width,
		Height: height - 2, // Minus status line and prompt line
	}

	StatusLineFrame = frame.Frame{
		Screen: Screen,
		X:      0,
		Y:      height - 2,
		Width:  width,
		Height: 1,
	}

	if PromptMenu != nil {
		PromptMenuFrame = StatusLineFrame
		StatusLineFrame.Y--
		TabFrame.Height--
	}

	// Tab bar
	if Tabs.Count() > 1 {
		TabFrame.Y++
		TabFrame.Height--
		f := frame.Frame{
			Screen: Screen,
			X:      0,
			Y:      0,
			Width:  width,
			Height: 1,
		}
		TabBarFrame = f
	} else {
		TabBarFrame = frame.NilFrame()
	}

	PromptFrame = frame.Frame{
		Screen: Screen,
		X:      0,
		Y:      height - 1,
		Width:  width,
		Height: 1,
	}

	// Render objects

	if Tabs.Count() > 1 {
		TabBar.Render(TabBarFrame)
	}

	if renderTab {
		CurrentTab.Render()
	}

	renderStatusLine()

	if PromptMenu != nil {
		PromptMenu.Render(PromptMenuFrame)
	}

	Prompt.Render()

	Screen.Show()
}

func OpenArgFiles() {
	errMsg := ""

	for _, file := range arg.Files {
		t, err := tab.Open(&TabFrame, file)
		if t != nil {
			if Tabs == nil {
				Tabs = t
			} else {
				Tabs.Append(t)
			}
			if CurrentTab == nil {
				CurrentTab = t
			}

			t.Go(arg.Line, arg.Column)
			t.ViewCenter()
		}
		if err != nil {
			errMsg += err.Error() + "\n\n"
		}
	}

	if len(errMsg) > 0 {
		errTab := tab.FromString(&TabFrame, errMsg, "error.enix")
		if Tabs == nil {
			Tabs = errTab
		} else {
			Tabs.Append(errTab)
		}
		CurrentTab = errTab
	}
}

func InitScreen() {
	if Screen != nil {
		screenQuitChan <- struct{}{}
		Screen.Fini()
		close(screenQuitChan)
	}

	var err error

	Screen, err = tcell.NewScreen()
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = Screen.Init()
	if err != nil {
		log.Fatalf("%v", err)
	}

	Screen.Clear()
	Screen.EnableMouse()

	tcellEventChan = make(chan tcell.Event)
	screenQuitChan = make(chan struct{})
	go Screen.ChannelEvents(tcellEventChan, screenQuitChan)
}

func Start() {
	InitScreen()

	// Catch panics in a defer, clean up, and re-raise them.
	// Otherwise the application can die without leaving any diagnostic trace.
	quit := func() {
		maybePanic := recover()
		Screen.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	width, height := Screen.Size()

	TabBarFrame = frame.NilFrame()
	TabFrame = frame.Frame{Screen: Screen, X: 0, Y: 0, Width: width, Height: height - 2}
	StatusLineFrame = frame.NilFrame()
	PromptMenuFrame = frame.NilFrame()
	PromptFrame = frame.NilFrame()

	Prompt = prompt{
		History:    make([]string, 0, 64),
		HistoryIdx: 0,
		Line:       nil,
		Cursor:     nil,
		View:       view.Zero(),
		ShadowText: "",
		State:      Inactive,
		PathDir:    "",
	}

	if len(arg.Files) == 0 {
		Tabs = tab.FromString(&TabFrame, "", "no-name")
		CurrentTab = Tabs
	} else {
		OpenArgFiles()
	}

	TabBar.Init()
	Render(true)

	changeWatcher := time.NewTicker(500 * time.Millisecond)

	// Init autosave ticker and stop it if autosave is disabled.
	// autoSaveTicker can't be nil, because of polling on channel in the select.
	autoSaveTicker = time.NewTicker(time.Second)
	if cfg.Cfg.AutoSave == 0 {
		autoSaveTicker.Stop()
	} else {
		autoSaveTicker.Reset(time.Duration(cfg.Cfg.AutoSave) * time.Second)
	}

	var tcellEvRcvr TcellEventReceiver

	for {
		if Quit {
			return
		}

		renderTab := false

		select {
		case ev := <-tcellEventChan:
			switch ev := ev.(type) {
			case *tcell.EventMouse:
				mEv := mouse.RxTcellEventMouse(ev)
				if mEv != nil {
					RxMouseEvent(mEv)
					renderTab = true
				} else {
					// Don't render the whole window, as nothing happened.
					// Don't waste CPU.
					continue
				}
			default:
				if tcellEvRcvr == nil {
					tcellEvRcvr = RxTcellEvent(ev)
				} else {
					tcellEvRcvr = tcellEvRcvr.RxTcellEvent(ev)
				}

				if tcellEvRcvr == nil {
					renderTab = true
				}
			}
		case <-changeWatcher.C:
			tab := CurrentTab
			if tab.ModTime.Compare(util.FileModTime(tab.Path)) < 0 {
				Prompt.AskTabReload(PromptFrame)
				tcellEvRcvr = &Prompt
			}
		case <-autoSaveTicker.C:
			if Prompt.State == TabReloadQuestion {
				continue
			}

			tab := CurrentTab.First()
			for tab != nil {
				tab.AutoSave()
				tab = tab.Next
			}
		}

		Render(renderTab)
	}
}
