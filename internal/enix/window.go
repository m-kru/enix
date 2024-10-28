package enix

import (
	"fmt"
	"log"
	"unicode"

	"github.com/m-kru/enix/internal/arg"
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/exec"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/mouse"
	"github.com/m-kru/enix/internal/tab"
	"github.com/m-kru/enix/internal/tabbar"

	"github.com/gdamore/tcell/v2"
)

type Window struct {
	Config     *cfg.Config
	Colors     *cfg.Colorscheme
	Keys       *cfg.Keybindings
	InsertKeys *cfg.Keybindings // Insert mode keybindings

	Mouse  mouse.Mouse
	TabBar tabbar.TabBar

	Screen tcell.Screen
	Width  int
	Height int

	TabBarFrame *frame.Frame
	TabFrame    frame.Frame

	Tabs       *tab.Tab // First tab
	CurrentTab *tab.Tab

	Prompt *Prompt
}

func (w *Window) RxMouseEvent(ev mouse.Event) {
	if !w.TabFrame.Within(ev.X(), ev.Y()) {
		// Currently only mouse events within tab frame are handled.
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
		w.CurrentTab.ViewDown()
	case mouse.WheelUp:
		w.CurrentTab.ViewUp()
	case mouse.WheelLeft:
		w.CurrentTab.ViewLeft()
	case mouse.WheelRight:
		w.CurrentTab.ViewRight()
	}
}

func (w *Window) RxTcellEvent(ev tcell.Event) TcellEventReceiver {
	switch ev := ev.(type) {
	case *tcell.EventResize:
		w.Resize()
	case *tcell.EventKey:
		return w.RxTcellEventKey(ev)
	}

	return w
}

func (w *Window) RxTcellEventKey(ev *tcell.EventKey) TcellEventReceiver {
	var err error
	var info string
	updateView := true
	tab := w.CurrentTab

	if tab.State != "" {
		tab.RxEventKey(ev)
		return w
	}

	if ev.Key() == tcell.KeyRune {
		r := ev.Rune()
		if unicode.IsDigit(r) {
			return w.RxDigit(r)
		}
	}

	c, err := w.Keys.ToCmd(ev)
	if err != nil {
		w.Prompt.ShowError(fmt.Sprintf("%v", err))
		return w
	}

	if c.Name == "" {
		return w
	}

	if tab.RepCount != 0 {
		c.RepCount = tab.RepCount
		tab.RepCount = 0
	}

	for i := 0; i < c.RepCount; i++ {
		switch c.Name {
		case "add-cursor":
			err = exec.AddCursor(c.Args, tab)
		case "backspace":
			err = exec.Backspace(c.Args, tab)
		case "cmd":
			tab.HasFocus = false
			w.Prompt.Activate("", "")
			return w.Prompt
		case "config-dir":
			info, err = exec.ConfigDir(c.Args)
		case "del":
			err = exec.Del(c.Args, tab)
		case "down":
			err = exec.Down(c.Args, tab)
		case "esc":
			err = exec.Esc(c.Args, tab)
			w.Prompt.Clear()
		case "find":
			tab.HasFocus = false
			w.Prompt.Activate("find ", "todo")
			return w.Prompt
		case "g", "go":
			err = exec.Go(c.Args, tab)
		case "help":
			tab.HasFocus = false
			w.Prompt.Activate("help ", "")
			return w.Prompt
		case "insert":
			tab.State = "insert"
		case "join":
			err = exec.Join(c.Args, tab)
		case "left":
			err = exec.Left(c.Args, tab)
		case "line-end":
			err = exec.LineEnd(c.Args, tab)
		case "line-start":
			err = exec.LineStart(c.Args, tab)
		case "m", "mark":
			info, err = exec.Mark(c.Args, tab)
		case "newline":
			err = exec.Newline(c.Args, tab)
		case "o", "open":
			tab.HasFocus = false
			w.Prompt.Activate("open ", "")
			return w.Prompt
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
		case "replace":
			tab.State = "replace"
		case "right":
			err = exec.Right(c.Args, tab)
		case "save":
			info, err = exec.Save(c.Args, tab, w.Config.TrimOnSave)
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
		case "up":
			err = exec.Up(c.Args, tab)
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
		case "word-end":
			err = exec.WordEnd(c.Args, tab)
		case "word-start":
			err = exec.WordStart(c.Args, tab)
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
		w.Prompt.ShowError(fmt.Sprintf("%v", err))
	} else if info != "" {
		w.Prompt.ShowInfo(info)
	}

	return w
}

func (w *Window) RxDigit(digit rune) TcellEventReceiver {
	d := int(digit - '0')

	t := w.CurrentTab
	t.RepCount = t.RepCount*10 + d
	if t.RepCount < 0 {
		t.RepCount = 0
	}

	return w
}

// Resize handles all the required logic when screen is resized.
func (w *Window) Resize() {
	w.Screen.Fill(' ', w.Colors.Default)
	w.Screen.Sync()

	width, height := w.Screen.Size()

	w.Width = width
	w.Height = height - 1

	w.Prompt.Frame = frame.Frame{
		Screen: w.Screen,
		X:      0,
		Y:      height - 1,
		Width:  width,
		Height: 1,
	}
}

func (w *Window) Render() {
	w.TabBarFrame = nil
	w.TabFrame = frame.Frame{
		Screen: w.Screen,
		X:      0,
		Y:      0,
		Width:  w.Width,
		Height: w.Height,
	}

	if w.Tabs.Count() > 1 {
		w.TabFrame.Y++
		w.TabFrame.Height--

		w.TabBarFrame = &frame.Frame{
			Screen: w.Screen,
			X:      0,
			Y:      0,
			Width:  w.Width,
			Height: 1,
		}
	}

	if w.TabBarFrame != nil {
		w.TabBar.Render(w.Tabs, w.CurrentTab, w.Colors, *w.TabBarFrame)
	}

	w.CurrentTab.Render(w.TabFrame)

	w.Screen.Show()
}

func (w *Window) OpenArgFiles() {
	w.Tabs = tab.Open(w.Config, w.Colors, w.InsertKeys, arg.Files[0])

	for i := 1; i < len(arg.Files); i++ {
		t := tab.Open(w.Config, w.Colors, w.InsertKeys, arg.Files[i])
		w.Tabs.Append(t)
	}

	w.CurrentTab = w.Tabs
}

func Start(
	config *cfg.Config,
	colors *cfg.Colorscheme,
	keys *cfg.Keybindings,
	promptKeys *cfg.Keybindings,
	insertKeys *cfg.Keybindings,
) {
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
	// Otherwise the application can  die without leaving any diagnostic trace.
	quit := func() {
		maybePanic := recover()
		screen.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	width, height := screen.Size()

	w := Window{
		Config:     config,
		Colors:     colors,
		Keys:       keys,
		InsertKeys: insertKeys,
		Screen:     screen,
		Width:      width,
		Height:     height - 1, // One line for prompt
	}

	p := Prompt{
		Config: config,
		Colors: colors,
		Keys:   promptKeys,
		Screen: screen,
		Frame: frame.Frame{
			Screen: screen,
			X:      0,
			Y:      height - 1,
			Width:  width,
			Height: 1,
		},
	}

	w.Prompt = &p
	p.Window = &w

	if len(arg.Files) == 0 {
		w.Tabs = tab.FromString(config, colors, insertKeys, "", "no-name")
		w.CurrentTab = w.Tabs
	} else {
		w.OpenArgFiles()
	}

	w.Render()

	var tcellEvRcvr TcellEventReceiver = &w
	tcellEventChan := make(chan tcell.Event)
	go screen.ChannelEvents(tcellEventChan, nil)

	for {
		select {
		case ev := <-tcellEventChan:
			switch ev := ev.(type) {
			case *tcell.EventMouse:
				mEv := w.Mouse.RxTcellEventMouse(ev)
				if mEv != nil {
					w.RxMouseEvent(mEv)
				} else {
					// Don't render the whole window, as nothing happened.
					// Don't waste CPU.
					continue
				}
			default:
				tcellEvRcvr = tcellEvRcvr.RxTcellEvent(ev)
				if tcellEvRcvr == &w {
					w.CurrentTab.HasFocus = true
				} else if tcellEvRcvr == nil {
					return
				}
			}
		}

		w.Render()
	}
}
