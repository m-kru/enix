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
	"github.com/m-kru/enix/internal/tabbar"
	"github.com/m-kru/enix/internal/util"
	"github.com/m-kru/enix/internal/view"

	"github.com/gdamore/tcell/v2"
)

var Window window

var autoSaveTicker *time.Ticker

type window struct {
	Screen tcell.Screen
	Width  int
	Height int

	TabBarFrame frame.Frame
	TabFrame    frame.Frame

	Tabs       *tab.Tab // First tab
	CurrentTab *tab.Tab
}

func (w *window) RxMouseEvent(ev mouse.Event) {
	if w.TabBarFrame.Within(ev.X(), ev.Y()) {
		newCurrentTab := tabbar.RxMouseEvent(ev)
		if newCurrentTab != nil {
			Window.CurrentTab = newCurrentTab
		}
		return
	}

	if !w.TabFrame.Within(ev.X(), ev.Y()) {
		return
	}

	x, y := w.TabFrame.ToFramePosition(ev.X(), ev.Y())

	switch ev.(type) {
	case mouse.PrimaryClick:
		w.CurrentTab.PrimaryClick(x, y)
	case mouse.DoublePrimaryClick:
		// Implement word selection here.
	case mouse.PrimaryClickCtrl:
		w.CurrentTab.PrimaryClickCtrl(x, y)
	case mouse.WheelDown:
		for range 5 {
			w.CurrentTab.ViewDown()
		}
	case mouse.WheelUp:
		for range 5 {
			w.CurrentTab.ViewUp()
		}
	case mouse.WheelLeft:
		for range 5 {
			w.CurrentTab.ViewLeft()
		}
	case mouse.WheelRight:
		for range 5 {
			w.CurrentTab.ViewRight()
		}
	}
}

func (w *window) RxTcellEvent(ev tcell.Event) TcellEventReceiver {
	switch ev := ev.(type) {
	case *tcell.EventResize:
		w.Resize()
	case *tcell.EventKey:
		return w.RxTcellEventKey(ev)
	}

	return w
}

func (w *window) RxTcellEventKey(ev *tcell.EventKey) TcellEventReceiver {
	var err error
	var info string
	updateView := true
	tab := w.CurrentTab

	if tab.State != "" {
		cmd := tab.RxEventKey(ev)
		if cmd != "" {
			Prompt.Activate(cmd, "")
			return &Prompt
		}
		return w
	}

	if ev.Key() == tcell.KeyRune {
		r := ev.Rune()
		if unicode.IsDigit(r) {
			return w.RxDigit(r)
		}
	}

	c, err := cfg.Keys.ToCmd(ev)
	if err != nil {
		Prompt.ShowError(fmt.Sprintf("%v", err))
		return w
	}

	if c.Name == "" {
		return w
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
			err = exec.Esc(c.Args, tab)
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
		case "newline":
			err = exec.Newline(c.Args, tab)
		case "paste":
			err = exec.Paste(c.Args, tab)
		case "pwd":
			info, err = exec.Pwd(c.Args)
		case "quit", "q":
			tab, err = exec.Quit(c.Args, tab, false)
			if err == nil && tab == nil {
				return nil
			} else if tab != nil {
				w.Tabs = tab.First()
				w.CurrentTab = tab
			}
			updateView = false
		case "quit!", "q!":
			tab, _ = exec.Quit(c.Args, tab, true)
			if tab == nil {
				return nil
			} else {
				w.Tabs = tab.First()
				w.Tabs = w.CurrentTab.First()
			}
			updateView = false
		case "redo":
			err = exec.Redo(c.Args, tab)
		case "replace":
			tab.State = "replace"
		case "right":
			err = exec.Right(c.Args, tab)
		case "save":
			info, err = exec.Save(c.Args, tab, cfg.Cfg.TrimOnSave)
			updateView = false
		case "search":
			Prompt.Activate("search ", tab.GetWord())
			return &Prompt
		case "sel-all":
			err = exec.SelAll(c.Args, tab)
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
			w.CurrentTab, err = exec.SelToTab(c.Args, tab)
		case "sel-up":
			err = exec.SelUp(c.Args, tab)
		case "sel-word":
			err = exec.SelWord(c.Args, tab)
		case "sel-word-end":
			err = exec.SelWordEnd(c.Args, tab)
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
			err = exec.Suspend(c.Args, w.Screen)
		case "tab":
			err = exec.Tab(c.Args, tab)
		case "tn", "tab-next":
			w.CurrentTab, err = exec.TabNext(c.Args, tab)
		case "tp", "tab-prev":
			w.CurrentTab, err = exec.TabPrev(c.Args, tab)
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

	return w
}

func (w *window) RxDigit(digit rune) TcellEventReceiver {
	d := int(digit - '0')

	t := w.CurrentTab
	t.RepCount = t.RepCount*10 + d
	if t.RepCount < 0 {
		t.RepCount = 0
	}

	return w
}

// Resize handles all the required logic when screen is resized.
func (w *window) Resize() {
	w.Screen.Fill(' ', cfg.Style.Default)
	w.Screen.Sync()

	width, height := w.Screen.Size()

	w.Width = width
	w.Height = height - 1

	Prompt.Frame = frame.Frame{
		Screen: w.Screen,
		X:      0,
		Y:      height - 1,
		Width:  width,
		Height: 1,
	}
}

func (w *window) Render() {
	w.TabFrame = frame.Frame{
		Screen: w.Screen,
		X:      0,
		Y:      0,
		Width:  w.Width,
		Height: w.Height,
	}

	// Tab bar
	if w.Tabs.Count() > 1 {
		w.TabFrame.Y++
		w.TabFrame.Height--
		f := frame.Frame{
			Screen: w.Screen,
			X:      0,
			Y:      0,
			Width:  w.Width,
			Height: 1,
		}
		tabbar.SetFrame(f)
		tabbar.Update(w.Tabs, w.CurrentTab)
		tabbar.Render(w.CurrentTab)
		w.TabBarFrame = f
	}

	w.CurrentTab.Render()

	w.Screen.Show()
}

func (w *window) OpenArgFiles() {
	errMsg := ""

	for _, file := range arg.Files {
		t, err := tab.Open(&w.TabFrame, file)
		if t != nil {
			if w.Tabs == nil {
				w.Tabs = t
			} else {
				w.Tabs.Append(t)
			}
			if w.CurrentTab == nil {
				w.CurrentTab = t
			}

			t.Go(arg.Line, arg.Column)
			t.ViewCenter()
		}
		if err != nil {
			errMsg += err.Error() + "\n\n"
		}
	}

	if len(errMsg) > 0 {
		errTab := tab.FromString(&w.TabFrame, errMsg, "error.enix")
		if w.Tabs == nil {
			w.Tabs = errTab
		} else {
			w.Tabs.Append(errTab)
		}
		w.CurrentTab = errTab
	}
}

func Start() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = screen.Init()
	if err != nil {
		log.Fatalf("%v", err)
	}

	screen.Clear()
	screen.EnableMouse()

	// Catch panics in a defer, clean up, and re-raise them.
	// Otherwise the application can die without leaving any diagnostic trace.
	quit := func() {
		maybePanic := recover()
		screen.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	width, height := screen.Size()

	Window = window{
		Screen:      screen,
		Width:       width,
		Height:      height - 1, // One line for prompt
		TabBarFrame: frame.Frame{Screen: screen, X: 0, Y: 0, Width: 0, Height: 0},
		TabFrame:    frame.Frame{Screen: screen, X: 0, Y: 0, Width: width, Height: height},
		Tabs:        nil,
		CurrentTab:  nil,
	}

	Prompt = prompt{
		Screen: screen,
		Frame: frame.Frame{
			Screen: screen,
			X:      0,
			Y:      height - 1,
			Width:  width,
			Height: 1,
		},
		History:    make([]string, 0, 64),
		HistoryIdx: 0,
		Line:       nil,
		Cursor:     nil,
		View:       view.Zero(),
		ShadowText: "",
		State:      InText,
	}

	if len(arg.Files) == 0 {
		Window.Tabs = tab.FromString(&Window.TabFrame, "", "no-name")
		Window.CurrentTab = Window.Tabs
	} else {
		Window.OpenArgFiles()
	}

	Window.Render()

	changeWatcher := time.NewTicker(500 * time.Millisecond)

	// Init autosave ticker and stop it if autosave is disabled.
	// autoSaveTicker can't be nil, because of polling on channel in the select.
	autoSaveTicker = time.NewTicker(time.Second)
	if cfg.Cfg.AutoSave == 0 {
		autoSaveTicker.Stop()
	} else {
		autoSaveTicker.Reset(time.Duration(cfg.Cfg.AutoSave) * time.Second)
	}

	var tcellEvRcvr TcellEventReceiver = &Window
	tcellEventChan := make(chan tcell.Event)
	go screen.ChannelEvents(tcellEventChan, nil)

	for {
		render := false

		select {
		case ev := <-tcellEventChan:
			switch ev := ev.(type) {
			case *tcell.EventMouse:
				mEv := mouse.RxTcellEventMouse(ev)
				if mEv != nil {
					Window.RxMouseEvent(mEv)
					render = true
				} else {
					// Don't render the whole window, as nothing happened.
					// Don't waste CPU.
					continue
				}
			default:
				tcellEvRcvr = tcellEvRcvr.RxTcellEvent(ev)
				if tcellEvRcvr == &Window {
					render = true
				} else if tcellEvRcvr == nil {
					return
				}
			}
		case <-changeWatcher.C:
			tab := Window.CurrentTab
			if tab.ModTime.Compare(util.FileModTime(tab.Path)) < 0 {
				Prompt.AskTabReload()
				tcellEvRcvr = &Prompt
			}
		case <-autoSaveTicker.C:
			if Prompt.State == TabReloadQuestion {
				continue
			}

			tab := Window.CurrentTab.First()
			for {
				if tab == nil {
					break
				}
				tab.AutoSave()
				tab = tab.Next
			}

			render = true
		}

		// Do not rerender if focus is on the prompt.
		// This reduces responsiveness in the case of large files.
		if render {
			Window.Render()
		}
	}
}
