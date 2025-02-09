package util

import (
	"strings"
	"unicode/utf8"
)

var fileExtToType = map[string]string{
	"c": "c", "h": "c",
	"fbd":  "fbdl",
	"go":   "go",
	"json": "json",
	"tex":  "latex",
	"md":   "markdown",
	"py":   "python",
	"tcl":  "tcl",
	"toml": "toml",
	"typ":  "typst",
	"vhd":  "vhdl", "vhdl": "vhdl",
	"sh": "sh", "bash": "sh", "csh": "sh", "ksh": "sh", "mksh": "sh", "zsh": "sh",
}

// FileNameToType returns file type based on the file name.
func FileNameToType(name string) string {
	switch name {
	case "COMMIT_EDITMSG":
		return "git-commit"
	case "Makefile", "makefile":
		return "make"
	case "Dockerfile":
		return "docker"
	case ".bashrc", ".profile", "bspwmrc":
		return "sh"
	}

	ss := strings.Split(name, ".")
	if len(ss) == 1 {
		return ""
	}

	ext := strings.ToLower(ss[len(ss)-1])

	if typ, ok := fileExtToType[ext]; ok {
		return typ
	}

	return ""
}

func IsBracket(r rune) bool {
	return r == '(' || r == ')' ||
		r == '[' || r == ']' ||
		r == '{' || r == '}' ||
		r == '<' || r == '>'
}

// ByteIdxToRuneIdx returns rune index in string (represented as byte slice) based on byte index.
// The function assumes provided string has only valid runes.
// In case of an invalid rune, value 0 is returned.
func ByteIdxToRuneIdx(buf []byte, byteIdx int) int {
	rIdx := 0
	bIdx := 0
	for {
		r, rLen := utf8.DecodeRune(buf[bIdx:])
		if r == utf8.RuneError {
			return rIdx
		}

		if bIdx >= byteIdx {
			break
		}

		bIdx += rLen
		rIdx++
	}

	return rIdx
}
