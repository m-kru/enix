package enix

import (
	"fmt"
	"log"

	"github.com/m-kru/enix/internal/arg"
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cmd"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/tab"

	"github.com/gdamore/tcell/v2"
)

type Window struct {
	Config *cfg.Config
	Colors *cfg.Colorscheme
	Keys   *cfg.Keybindings

	Screen tcell.Screen
	Width  int
	Height int

	TabFrame   frame.Frame
	Tabs       *tab.Tab // First tab
	CurrentTab *tab.Tab

	Prompt *Prompt
}

func (w *Window) RxEvent(ev tcell.Event) EventReceiver {
	var err error

	switch ev := ev.(type) {
	case *tcell.EventResize:
		w.Resize()
	case *tcell.EventKey:
		name, args := w.Keys.ToCmd(ev)

		switch name {
		case "cmd":
			w.Prompt.Activate("", "")
			return w.Prompt
		case "down":
			err = cmd.Down(args, w.CurrentTab)
		case "left":
			err = cmd.Left(args, w.CurrentTab)
		case "right":
			err = cmd.Right(args, w.CurrentTab)
		case "up":
			err = cmd.Up(args, w.CurrentTab)
		case "spawn-down":
			w.CurrentTab.CursorSpawnDown()
		case "esc":
			w.Prompt.Clear()
		case "find":
			w.Prompt.Activate("find ", "todo")
			return w.Prompt
		case "help":
			w.Prompt.Activate("help ", "")
			return w.Prompt
		case "quit":
			return nil
		}
	}

	if err != nil {
		w.Prompt.ShowError(fmt.Sprintf("%v", err))
	}

	w.Render()

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
	w.Tabs = tab.Open(w.Config, w.Colors, arg.Files[0])

	prevT := w.Tabs

	for i := 1; i < len(arg.Files); i++ {
		t := tab.Open(w.Config, w.Colors, arg.Files[0])
		prevT.Next = t
		t.Prev = prevT
		prevT = t
	}

	w.CurrentTab = w.Tabs
}

func Start(config *cfg.Config, colors *cfg.Colorscheme, keys *cfg.Keybindings) {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = screen.Init()
	if err != nil {
		log.Fatalf("%v", err)
	}

	screen.Clear()

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
		Config: config,
		Colors: colors,
		Keys:   keys,
		Screen: screen,
		Width:  width,
		Height: height - 1, // One line for prompt
	}

	p := Prompt{
		Config: config,
		Colors: colors,
		Keys:   keys,
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
		w.Tabs = tab.FromString(config, colors, "", "No Name")
		w.CurrentTab = w.Tabs
	} else {
		w.OpenArgFiles()
	}

	w.Render()

	var evRcvr EventReceiver = &w

	for {
		ev := screen.PollEvent()
		evRcvr = evRcvr.RxEvent(ev)
		if evRcvr == nil {
			return
		}
	}
}
