package util

import (
	"fmt"
	"unicode"
)

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

// WordStart finds next word start index.
func WordStart(line []rune, idx int) (int, bool) {
	if idx >= len(line)-1 {
		return 0, false
	}

	for {
		idx++
		if IsWordRune(line[idx]) && !IsWordRune(line[idx-1]) {
			return idx, true
		}

		if idx == len(line)-1 {
			break
		}
	}

	return 0, false
}

// GetWord returns word containing index idx.
// In case of whitespaces an empty string is returned.
func GetWord(line []rune, idx int) string {
	if len(line) == 0 || idx >= len(line) {
		return ""
	}

	if unicode.IsSpace(line[idx]) {
		return ""
	}

	if !IsWordRune(line[idx]) {
		return string(line[idx : idx+1])
	}

	left := idx
	for {
		if left == 0 {
			break
		}
		left--

		if !IsWordRune(line[left]) {
			left++
			break
		}
	}

	right := idx
	for {
		if right == len(line)-1 {
			break
		}
		right++

		if !IsWordRune(line[right]) {
			right--
			break
		}
	}

	return string(line[left : right+1])
}
