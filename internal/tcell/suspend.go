package tcell

import (
	"fmt"
	"syscall"

	"github.com/gdamore/tcell/v2"
)

func Suspend(screen tcell.Screen) error {
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
