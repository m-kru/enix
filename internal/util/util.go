package util

import (
	"strings"
	"unicode/utf8"
)

var fileExtToType = map[string]string{
	"c": "c", "h": "c",
	"css":  "css",
	"elm":  "elm",
	"go":   "go",
	"html": "html",
	"java": "java",
	"jl":   "julia",
	"json": "json",
	"kt":   "kotlin", "kts": "kotlin",
	"lisp":  "lisp",
	"tex":   "latex",
	"lua":   "lua",
	"php":   "php",
	"proto": "protobuf",
	"py":    "python",
	"r":     "r",
	"rb":    "ruby",
	"rs":    "rust",
	"sql":   "sql",
	"tcl":   "tcl",
	"toml":  "toml",
	"typ":   "typst",
	"vhd":   "vhdl", "vhdl": "vhdl",
	"yaml": "yaml", "yml": "yaml",
	"zig": "zig", "zon": "zig",
}

// FileNameToType returns file type based on the file name.
func FileNameToType(name string) string {
	switch name {
	case "Makefile", "makefile":
		return "make"
	case "Dockerfile":
		return "docker"
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

// ByteIdxToRuneIdx returns rune index in string based on byte index.
// The function assumes provided string has only valid runes.
// In case of invalid runes value 0 is returned.
func ByteIdxToRuneIdx(str string, bidx int) int {
	idx := 0
	for rIdx, r := range str {
		rl := utf8.RuneLen(r)
		if rl < 0 {
			return 0
		}

		if idx >= bidx {
			return rIdx
		}

		idx += rl
	}

	return 0
}
