package enix

import (
	"fmt"
	"log"
	"strings"

	"github.com/m-kru/enix/internal/arg"
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cmd"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/mouse"
	"github.com/m-kru/enix/internal/tab"

	"github.com/gdamore/tcell/v2"
)

type Window struct {
	Config     *cfg.Config
	Colors     *cfg.Colorscheme
	Keys       *cfg.Keybindings
	InsertKeys *cfg.Keybindings // Insert mode keybindings

	Mouse mouse.Mouse

	Screen tcell.Screen
	Width  int
	Height int

	TabFrame   frame.Frame
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
	}
}

func (w *Window) RxTcellEvent(ev tcell.Event) TcellEventReceiver {
	var err error
	updateView := true
	tab := w.CurrentTab

	switch ev := ev.(type) {
	case *tcell.EventResize:
		w.Resize()
	case *tcell.EventKey:
		if tab.InInsertMode {
			tab.RxEventKey(ev)
			return w
		}

		name, argStr := w.Keys.ToCmd(ev)
		args := strings.Fields(argStr)

		if name == "" {
			break
		}

		switch name {
		case "add-cursor":
			err = cmd.AddCursor(args, tab)
		case "cmd":
			tab.HasFocus = false
			w.Prompt.Activate("", "")
			return w.Prompt
		case "down":
			err = cmd.Down(args, tab)
		case "esc":
			err = cmd.Esc(args, tab)
			w.Prompt.Clear()
		case "find":
			tab.HasFocus = false
			w.Prompt.Activate("find ", "todo")
			return w.Prompt
		case "help":
			tab.HasFocus = false
			w.Prompt.Activate("help ", "")
			return w.Prompt
		case "insert":
			tab.InInsertMode = true
		case "left":
			err = cmd.Left(args, tab)
		case "line-end":
			err = cmd.LineEnd(args, tab)
		case "line-start":
			err = cmd.LineStart(args, tab)
		case "newline":
			err = cmd.Newline(args, tab)
		case "quit", "q":
			err = cmd.Quit(args, tab, false)
			if err == nil {
				return nil
			}
		case "quit!", "q!":
			_ = cmd.Quit(args, tab, true)
			return nil
		case "right":
			err = cmd.Right(args, tab)
		case "save":
			err = cmd.Save(args, tab, w.Config.TrimOnSave)
		case "space":
			err = cmd.Space(args, tab)
		case "spawn-down":
			err = cmd.SpawnDown(args, tab)
		case "spawn-up":
			err = cmd.SpawnUp(args, tab)
		case "tab":
			err = cmd.Tab(args, tab)
		case "trim":
			err = cmd.Trim(args, tab)
		case "up":
			err = cmd.Up(args, tab)
		case "view-down":
			err = cmd.ViewDown(args, tab)
			updateView = false
		case "view-left":
			err = cmd.ViewLeft(args, tab)
			updateView = false
		case "view-right":
			err = cmd.ViewRight(args, tab)
			updateView = false
		case "view-up":
			err = cmd.ViewUp(args, tab)
			updateView = false
		case "word-end":
			err = cmd.WordEnd(args, tab)
		case "word-start":
			err = cmd.WordStart(args, tab)
		case "prev-word-start":
			err = cmd.PrevWordStart(args, tab)
		default:
			err = fmt.Errorf(
				"invalid or unimplemeneted command '%s', if unimplemented report on https://github.com/m-kru/enix",
				name,
			)
			updateView = false
		}
	}

	if updateView {
		tab.UpdateView()
	}

	if err != nil {
		w.Prompt.ShowError(fmt.Sprintf("%v", err))
	}

	return w
}

// Resize handles all the required logic when screen is resized.
func (w *Window) Resize() {
	w.Screen.Clear()
	w.Screen.Sync()

	width, height := w.Screen.Size()

	w.Width = width
	w.Height = height - 1

	w.TabFrame = frame.Frame{
		Screen: w.Screen,
		X:      0,
		Y:      0,
		Width:  w.Width,
		Height: w.Height,
	}

	w.Prompt.Frame = frame.Frame{
		Screen: w.Screen,
		X:      0,
		Y:      height - 1,
		Width:  width,
		Height: 1,
	}
}

func (w *Window) Render() {
	w.CurrentTab.Render(w.TabFrame)

	w.Screen.Show()
}

func (w *Window) OpenArgFiles() {
	w.Tabs = tab.Open(w.Config, w.Colors, w.InsertKeys, arg.Files[0])

	prevT := w.Tabs

	for i := 1; i < len(arg.Files); i++ {
		t := tab.Open(w.Config, w.Colors, w.InsertKeys, arg.Files[0])
		prevT.Next = t
		t.Prev = prevT
		prevT = t
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

	// Create buffered channel to prevent deadlocks.
	w.Mouse.EventChan = make(chan mouse.Event, 8)

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

	w.TabFrame = frame.Frame{
		Screen: w.Screen,
		X:      0,
		Y:      0,
		Width:  w.Width,
		Height: w.Height,
	}

	w.Prompt = &p
	p.Window = &w

	if len(arg.Files) == 0 {
		w.Tabs = tab.FromString(config, colors, insertKeys, "", "No Name")
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
				w.Mouse.RxTcellEventMouse(ev)
			default:
				tcellEvRcvr = tcellEvRcvr.RxTcellEvent(ev)
				if tcellEvRcvr == &w {
					w.CurrentTab.HasFocus = true
				} else if tcellEvRcvr == nil {
					return
				}
			}
		case ev := <-w.Mouse.EventChan:
			w.RxMouseEvent(ev)
		}

		w.Render()
	}
}
