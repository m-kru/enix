package enix

import (
	"fmt"
	"log"
	"strings"

	"github.com/m-kru/enix/internal/arg"
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/exec"
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
	var info string
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
			err = exec.AddCursor(args, tab)
		case "backspace":
			err = exec.Backspace(args, tab)
		case "exec":
			tab.HasFocus = false
			w.Prompt.Activate("", "")
			return w.Prompt
		case "config-dir":
			info, err = exec.ConfigDir(args)
		case "del":
			err = exec.Del(args, tab)
		case "down":
			err = exec.Down(args, tab)
		case "esc":
			err = exec.Esc(args, tab)
			w.Prompt.Clear()
		case "find":
			tab.HasFocus = false
			w.Prompt.Activate("find ", "todo")
			return w.Prompt
		case "g", "go":
			err = exec.Go(args, tab)
		case "help":
			tab.HasFocus = false
			w.Prompt.Activate("help ", "")
			return w.Prompt
		case "insert":
			tab.InInsertMode = true
		case "join":
			err = exec.Join(args, tab)
		case "left":
			err = exec.Left(args, tab)
		case "line-end":
			err = exec.LineEnd(args, tab)
		case "line-start":
			err = exec.LineStart(args, tab)
		case "m", "mark":
			info, err = exec.Mark(args, tab)
		case "newline":
			err = exec.Newline(args, tab)
		case "quit", "q":
			err = exec.Quit(args, tab, false)
			if err == nil {
				return nil
			}
		case "quit!", "q!":
			_ = exec.Quit(args, tab, true)
			return nil
		case "right":
			err = exec.Right(args, tab)
		case "save":
			info, err = exec.Save(args, tab, w.Config.TrimOnSave)
		case "space":
			err = exec.Space(args, tab)
		case "spawn-down":
			err = exec.SpawnDown(args, tab)
		case "spawn-up":
			err = exec.SpawnUp(args, tab)
		case "suspend":
			err = exec.Suspend(args, w.Screen)
		case "tab":
			err = exec.Tab(args, tab)
		case "trim":
			err = exec.Trim(args, tab)
		case "up":
			err = exec.Up(args, tab)
		case "view-down":
			err = exec.ViewDown(args, tab)
			updateView = false
		case "view-left":
			err = exec.ViewLeft(args, tab)
			updateView = false
		case "view-right":
			err = exec.ViewRight(args, tab)
			updateView = false
		case "view-up":
			err = exec.ViewUp(args, tab)
			updateView = false
		case "word-end":
			err = exec.WordEnd(args, tab)
		case "word-start":
			err = exec.WordStart(args, tab)
		case "prev-word-start":
			err = exec.PrevWordStart(args, tab)
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
	} else if info != "" {
		w.Prompt.ShowInfo(info)
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
