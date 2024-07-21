package arg

import (
	"fmt"
	"log"
	"os"
)

// Prase parses command line arguments.
func Parse() {
	param := ""
	val := false // Expected parameter value

	handleFlag := func(f string) {
		switch f {
		case "-batch":
			Batch = true
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
			case "-script":
				Script = arg
			default:
				panic(fmt.Sprintf("unhandled param '%s', implement me", param))
			}
		} else if isValidFlag(arg) {
			handleFlag(arg)
		} else if isValidParam(arg) {
			handleParam(arg)
		} else if arg[0] == '-' {
			log.Fatalf("invalid option '%s'", arg)
		} else {
			Files = append(Files, arg)
		}
	}

	validate()
}

// Function validate validates arguments after parsing.
func validate() {
	if Batch && Script == "" {
		log.Fatalf("batch mode requires the '-script' parameter to be set")
	}
}
