package tab

import (
	"fmt"
	"io"
)

func (t *Tab) Save(strWr io.StringWriter) error {
	l := t.Lines
	i := 1
	for {
		if l == nil {
			break
		}

		_, err := strWr.WriteString(
			fmt.Sprintf("%s%s", l.String(), t.Newline),
		)
		if err != nil {
			return fmt.Errorf("%s:%d: %v", t.Path, i, err)
		}

		l = l.Next
		i++
	}

	t.HasChanges = false

	return nil
}
