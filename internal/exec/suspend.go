package exec

import (
	"fmt"

	"github.com/gdamore/tcell/v2"

	enixTcell "github.com/m-kru/enix/internal/tcell"
)

func Suspend(args []string, screen tcell.Screen) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"suspend: expected 0 args, provided %d", len(args),
		)
	}

	return enixTcell.Suspend(screen)
}
