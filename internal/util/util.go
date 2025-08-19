package util

import (
	"strings"
	"unicode/utf8"
)

var fileNameToType = map[string]string{
	".bashrc": "sh", ".profile": "sh", "bspwmrc": "sh",
	".tclshrc":       "tcl",
	"COMMIT_EDITMSG": "git-commit",
	"Dockerfile":     "docker",
	"Makefile":       "make", "makefile": "make",
}

var fileExtToType = map[string]string{
	"c": "c", "h": "c",
	"conf": "conf",
	"dts":  "dts", "dtsi": "dts", "dtso": "dts",
	"fbd":   "fbdl",
	"go":    "go",
	"json":  "json",
	"tex":   "tex",
	"md":    "markdown",
	"mk":    "make",
	"patch": "patch",
	"py":    "python",
	"tcl":   "tcl",
	"toml":  "toml",
	"typ":   "typst",
	"vhd":   "vhdl", "vhdl": "vhdl",
	"sh": "sh", "bash": "sh", "csh": "sh", "ksh": "sh", "mksh": "sh", "zsh": "sh",
	"xml": "xml",
}

func IsValidFiletype(ft string) bool {
	for _, typ := range fileNameToType {
		if ft == typ {
			return true
		}
	}

	for _, typ := range fileExtToType {
		if ft == typ {
			return true
		}
	}

	return false
}

// FileNameToType returns file type based on the file name.
func FileNameToType(name string) string {
	if typ, ok := fileNameToType[name]; ok {
		return typ
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

func ShebangToFiletype(sb string) string {
	if !strings.HasPrefix(sb, "#!") {
		return ""
	}

	if strings.Contains(sb, "bash") {
		return "sh"
	} else if strings.Contains(sb, "csh") {
		return "sh"
	} else if strings.Contains(sb, "ksh") {
		return "h"
	} else if strings.Contains(sb, "mksh") {
		return "sh"
	} else if strings.Contains(sb, "python") {
		return "python"
	} else if strings.Contains(sb, "tclsh") {
		return "tcl"
	} else if strings.Contains(sb, "sh") {
		return "sh"
	} else if strings.Contains(sb, "zsh") {
		return "sh"
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
