package help

// IsValidCmd returns true if given command is a valid command.
func IsValidCmd(cmd string) bool {
	if _, ok := cmds[cmd]; ok {
		return true
	}
	return false
}
