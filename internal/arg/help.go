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

  -dump-config       Dump config to stdout.
  -dump-keys         Dump normal mode keybindings to stdout.
  -dump-insert-keys  Dump insert mode keybindings to stdout.
  -dump-prompt-keys  Dump prompt keybindings to stdout.
  -help              Display help.
  -profile           Enable CPU profiling and dump results to ./enix.prof file.
  -version           Display version.

Parameters:

  -config  Read the configuration from the provided file.
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
