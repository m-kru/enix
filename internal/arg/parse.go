package arg

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func handleLineAndColumn(arg string) {
	// Handle line only case
	if !strings.Contains(arg, ":") {
		n, err := strconv.Atoi(arg[1 : len(arg)-1])
		if err != nil {
			handleFile(arg)
		}
		Line = n
		return
	}

	// Handle line and column case
	lineStr, colStr, _ := strings.Cut(arg[1:len(arg)-1], ":")
	var line, col int
	var err error
	line, err = strconv.Atoi(lineStr)
	if err != nil {
		handleFile(arg)
		return
	}
	col, err = strconv.Atoi(colStr)
	if err != nil {
		handleFile(arg)
		return
	}

	Line = line
	Column = col
}

// addFile adds file path to the Files list if the path is not yet in the list.
func addFile(path string) {
	fileInfo, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		// Do nothing, file will be created later.
	} else if err == nil && fileInfo.IsDir() {
		log.Fatalf("cannot open '%s', it is a directory", path)
	} else if err != nil {
		log.Fatalf("file '%s': %v", path, err)
	}

	for _, f := range Files {
		if path == f {
			return
		}
	}
	Files = append(Files, path)
}

func handleFile(path string) {
	matches, err := filepath.Glob(path)
	if err != nil {
		log.Fatalf("cannot glob '%s' path: %v", path, err)
	}

	if len(matches) == 0 {
		addFile(path)
	}

	for _, m := range matches {
		addFile(m)
	}
}

// Prase parses command line arguments.
func Parse() {
	param := ""
	val := false // Expected parameter value

	handleFlag := func(f string) {
		switch f {
		case "-dump-config":
			DumpConfig = true
		case "-dump-keys":
			DumpKeys = true
		case "-dump-prompt-keys":
			DumpPromptKeys = true
		case "-help":
			printHelp()
		case "-version":
			printVersion()
		default:
			panic(fmt.Sprintf("unhandled flag '%s', implement me", f))
		}
	}

	handleParam := func(p string) {
		val = true
		if isValidParam(p) {
			param = p
		} else {
			log.Fatalf("invalid parameter '%s'", p)
		}
	}

	for _, arg := range os.Args[1:] {
		if val {
			val = false

			switch param {
			case "-config":
				Config = arg
			case "-script":
				Script = arg
			default:
				panic(fmt.Sprintf("unhandled param '%s', implement me", param))
			}
		} else if isValidFlag(arg) {
			handleFlag(arg)
		} else if isValidParam(arg) {
			handleParam(arg)
		} else if arg[0] == '+' {
			handleLineAndColumn(arg)
		} else if arg[0] == '-' {
			log.Fatalf("invalid option '%s', check 'enix -help'", arg)
		} else {
			handleFile(arg)
		}
	}

	if val {
		log.Fatalf("missing value for %s parameter", param)
	}

	validate()
}

// Function validate validates arguments after parsing.
func validate() {
	if Script != "" && len(Files) == 0 {
		log.Fatalf("script file provided, but no files to execute script on")
	}
}
