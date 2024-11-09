package tab

import "regexp"

func (tab *Tab) Search(expr string) error {
	re, err := regexp.Compile(expr)
	if err != nil {
		return err
	}

	tab.SearchCtx.Regexp = re

	return nil
}
