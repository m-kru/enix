package util

import "strings"

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
