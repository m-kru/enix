package cmd

import (
	"fmt"
	"github.com/m-kru/enix/internal/cfg"
	"strconv"
)

func CfgTabWidth(args []string, config *cfg.Config) error {
	if len(args) != 1 {
		return fmt.Errorf(
			"tab-width: provided %d args, expected 1", len(args),
		)
	}

	n, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("tab-width: %v", err)
	}

	if n <= 0 {
		return fmt.Errorf("tab-width: width must be greater than 0")
	}

	config.TabWidth = n

	return nil
}
