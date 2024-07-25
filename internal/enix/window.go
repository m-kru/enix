package enix

import (
	"log"

	"github.com/m-kru/enix/internal/cfg"
	_ "github.com/m-kru/enix/internal/cmd"

	"github.com/gdamore/tcell/v2"
)

type Window struct {
	Screen tcell.Screen
	Width  int
	Height int

	Colors *cfg.Colorscheme
	Keys   *cfg.Keybindings

	Tab        *Tab // First tab
	CurrentTab *Tab

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

	w.Prompt.Width = width
	w.Prompt.Y = height - 1
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
		Screen: screen,
		Width:  width - 1,
		Height: height - 2, // One line for prompt
		Colors: colors,
		Keys:   keys,
	}

	p := Prompt{
		Screen: screen,
		Width:  width,
		Y:      height - 1,
		Colors: colors,
		Keys:   keys,
	}

	w.Prompt = &p
	p.Window = &w

	var evRcvr EventReceiver = &w

	for {
		ev := screen.PollEvent()
		evRcvr = evRcvr.RxEvent(ev)
		if evRcvr == nil {
			return
		}
	}
}
