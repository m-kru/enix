package cmd

import (
	"fmt"
	"github.com/m-kru/enix/internal/cfg"
	"strconv"
	"strings"
)

func CfgTabWidth(args string, config *cfg.Config) error {
	sstr := strings.Fields(args)
	if len(sstr) != 1 {
		return fmt.Errorf(
			"cfg-tab-width: provided %d args, expected 1", len(sstr),
		)
	}

	n, err := strconv.Atoi(sstr[0])
	if err != nil {
		return fmt.Errorf("cfg-tab-width: %v", err)
	}

	if n <= 0 {
		return fmt.Errorf("cfg-tab-width: width must be greater than 0")
	}

	config.TabWidth = n

	return nil
}
