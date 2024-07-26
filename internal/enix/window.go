package enix

import (
	"log"

	"github.com/m-kru/enix/internal/arg"
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/tab"

	"github.com/gdamore/tcell/v2"
)

type Window struct {
	Colors *cfg.Colorscheme
	Keys   *cfg.Keybindings

	Screen tcell.Screen
	Width  int
	Height int

	Tabs       *tab.Tab // First tab pointer
	CurrentTab *tab.Tab

	Prompt *Prompt
}

func (w *Window) RxEvent(ev tcell.Event) EventReceiver {
	switch ev := ev.(type) {
	case *tcell.EventResize:
		w.Screen.Sync()
	case *tcell.EventKey:
		switch w.Keys.ToCmd(ev) {
		case "cmd":
			w.Prompt.Activate("text ", "shadow")
			return w.Prompt
		case "escape":
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

	return w
}

// Resize handles all the required logic when screen is resized.
func (w *Window) Resize() {
	w.Screen.Clear()
	w.Screen.Sync()

	width, height := w.Screen.Size()

	w.Width = width - 1
	w.Height = height - 2

	w.Prompt.Frame = frame.Frame{
		Screen: w.Screen,
		X:      0,
		Y:      height - 1,
		Width:  width,
		Height: 1,
	}
}

func (w *Window) Render() {

	tabFrame := frame.Frame{
		Screen: w.Screen,
		X:      0,
		Y:      0,
		Width:  w.Width,
		Height: w.Height,
	}
	w.CurrentTab.Render(tabFrame)
}

func Start(colors *cfg.Colorscheme, keys *cfg.Keybindings) {
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
		Colors: colors,
		Keys:   keys,
		Screen: screen,
		Width:  width - 1,
		Height: height - 2, // One line for prompt
	}

	p := Prompt{
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

	w.Prompt = &p
	p.Window = &w

	if len(arg.Files) == 0 {
		w.Tabs = tab.Empty(colors, screen, 0, width-1, 0, height-2)
		w.CurrentTab = w.Tabs
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
