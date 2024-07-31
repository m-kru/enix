package arg

var Batch bool
var Script string // Path to the script to be run.
var Line int = 1
var Column int = 1
var Files []string // Paths to text files.

func isValidFlag(f string) bool {
	flags := map[string]bool{
		"-batch": true, "-help": true, "-version": true,
	}
	if _, ok := flags[f]; ok {
		return true
	}
	return false
}

func isValidParam(p string) bool {
	params := map[string]bool{
		"-script": true,
	}
	if _, ok := params[p]; ok {
		return true
	}
	return false
}
