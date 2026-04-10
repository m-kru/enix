package util

import (
	"fmt"
	"strconv"
	"strings"
)

func AddIndent(str string, indent string, indentFirstLine bool) string {
	b := strings.Builder{}

	addIndent := indentFirstLine

	for _, r := range str {
		if addIndent {
			b.WriteString(indent)
		}

		b.WriteRune(r)

		addIndent = r == '\n'
	}

	return b.String()
}

// In the case of invalid format, returned line and column equal 1, not 0.
func ParseLineAndColumnString(str string) (int, int, error) {
	if !strings.HasPrefix(str, "+") {
		return 1, 1, fmt.Errorf("line and column string must start with '+'")
	}

	str = str[1:]

	// Handle line only case
	if !strings.Contains(str, ":") {
		line, err := strconv.Atoi(str)
		return line, 1, err
	}

	// Handle line and column case
	lineStr, colStr, _ := strings.Cut(str, ":")
	var line, col int
	var err error
	line, err = strconv.Atoi(lineStr)
	if err != nil {
		return 1, 1, err
	}
	col, err = strconv.Atoi(colStr)
	if err != nil {
		return 1, 1, err
	}

	return line, col, nil
}
