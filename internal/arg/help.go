package arg

import (
	"fmt"
	"os"
)

var helpMsg string = `enix - terminal-based text editor trying to follow the *nix philosophy.

Version: %s

Usage:

  enix [flags] [parameters] [path/to/text/file ...] [+line[:column]]

Flags:

  -help     Display help.
  -version  Display version.

Parameters:

  -script  Instead of opening files in the interactive tui mode, execute
           commands from the script on each file. Each command in the script
           file must be placed in a separate line. Empty lines are ignored.
           Lines starting with the '#' character are treated as comment lines.
           This parameter is mostly useful for enix internal regression tests.
`

func printHelp() {
	fmt.Printf(helpMsg, Version)
	os.Exit(0)
}
