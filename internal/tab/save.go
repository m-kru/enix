package tab

import (
	"fmt"
	"io"
)

func (tab *Tab) Save(strWr io.StringWriter) error {
	l := tab.Lines
	i := 1
	for {
		if l == nil {
			break
		}

		nl := tab.Newline
		if l.Next == nil {
			nl = ""
		}
		_, err := strWr.WriteString(fmt.Sprintf("%s%s", l.String(), nl))
		if err != nil {
			return fmt.Errorf("%s:%d: %v", tab.Path, i, err)
		}

		l = l.Next
		i++
	}

	tab.HasChanges = false

	return nil
}
