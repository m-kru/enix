package script

import (
	"os"

	"github.com/m-kru/enix/internal/arg"
	"github.com/m-kru/enix/internal/cmd"
	"github.com/m-kru/enix/internal/exec"
	"github.com/m-kru/enix/internal/tab"
)

func parseScriptFile() ([]cmd.Command, error) {
	script, err := os.ReadFile(arg.Script)
	if err != nil {
		return nil, err
	}

	return cmd.ParseScript(string(script))
}

func ExecOnFiles() error {
	cmds, err := parseScriptFile()
	if err != nil {
		return err
	}

	for _, file := range arg.Files {
		tab, err := tab.Open(nil, file)
		if err != nil {
			return err
		}
		err = Exec(tab, cmds)
		if err != nil {
			return err
		}
	}

	return nil
}

func Exec(tab *tab.Tab, cmds []cmd.Command) error {
	for _, cmd := range cmds {
		err := exec.Exec(cmd, tab)
		if err != nil {
			return err
		}
	}

	return nil
}
