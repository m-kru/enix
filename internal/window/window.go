package window

import (
	"log"

	"github.com/m-kru/enix/internal/cfg"

	"github.com/gdamore/tcell/v2"
)

func Start(colors cfg.Colorscheme) {
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

	for {
		ev := screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			
		}
	}
}
