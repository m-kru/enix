package exec

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"syscall"
)

func Suspend(args []string, screen tcell.Screen) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"suspend: expected 0 args, provided %d", len(args),
		)
	}

	err := screen.Suspend()
	if err != nil {
		return fmt.Errorf("suspend: %v", err)
	}

	err = syscall.Kill(syscall.Getpid(), syscall.SIGSTOP)
	if err != nil {
		return fmt.Errorf("suspend: %v", err)
	}

	err = screen.Resume()
	if err != nil {
		return fmt.Errorf("suspend: %v", err)
	}

	return nil
}
