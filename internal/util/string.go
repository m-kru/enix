package util

import (
	"strings"
)

func AddIndent(str string, indent string) string {
	b := strings.Builder{}

	addIndent := true

	for _, r := range str {
		if addIndent {
			b.WriteString(indent)
		}

		b.WriteRune(r)

		addIndent = r == '\n'
	}

	return b.String()
}
