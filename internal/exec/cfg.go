package exec

import (
	"fmt"
	"github.com/m-kru/enix/internal/cfg"
)

func ConfigDir(args []string) (string, error) {
	if len(args) != 0 {
		return "", fmt.Errorf(
			"config-dir: provided %d args, expected 0", len(args),
		)
	}

	return cfg.ConfigDir, nil
}
