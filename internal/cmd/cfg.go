package cmd

import (
	"fmt"
	"github.com/m-kru/enix/internal/cfg"
	"strconv"
	"strings"
)

func CfgTabWidth(args string, config *cfg.Config) error {
	fields := strings.Fields(args)
	if len(fields) != 1 {
		return fmt.Errorf(
			"tab-width: provided %d args, expected 1", len(fields),
		)
	}

	n, err := strconv.Atoi(fields[0])
	if err != nil {
		return fmt.Errorf("tab-width: %v", err)
	}

	if n <= 0 {
		return fmt.Errorf("tab-width: width must be greater than 0")
	}

	config.TabWidth = n

	return nil
}
