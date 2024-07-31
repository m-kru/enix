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

  -batch    Run in batch mode instead of tui mode.
  -help     Display help.
  -version  Display version.

Parameters:

  -script  Path to the script with commands to be executed on the files in the
           batch mode. For the batch mode, this parameter is obligatory.
           In the case of the tui mode, the commands are executed on the files
           after program start before the user gains control. Each command must
           be placed in a separate line. Empty lines are ignored. Lines
           starting with the '#' character are treated as comment lines.
`

func printHelp() {
	fmt.Printf(helpMsg, Version)
	os.Exit(0)
}
