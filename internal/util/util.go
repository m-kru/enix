package util

import (
	"fmt"
	"strings"
	"unicode"
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

func IsWordRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
}

// IntWidth returns number of digits required to print n.
// TODO: Improbe speed, not the fastest implementation.
func IntWidth(i int) int {
	return len(fmt.Sprintf("%d", i))
}

// PrevWordStart finds previous word start index.
func PrevWordStart(line []rune, idx int) (int, bool) {
	if idx == 0 {
		return 0, false
	}

	for {
		idx--
		if idx == 0 {
			if IsWordRune(line[idx]) {
				return idx, true
			} else {
				break
			}
		}

		if IsWordRune(line[idx]) && !IsWordRune(line[idx-1]) {
			return idx, true
		}
	}

	return 0, false
}

// WordEnd finds next word end index.
func WordEnd(line []rune, idx int) (int, bool) {
	if idx >= len(line)-1 {
		return 0, false
	}

	for {
		idx++
		if idx == len(line)-1 {
			if IsWordRune(line[idx]) {
				return idx, true
			} else {
				break
			}
		}

		if IsWordRune(line[idx]) && !IsWordRune(line[idx+1]) {
			return idx, true
		}
	}

	return 0, false
}
