package util

import (
	"fmt"
	"unicode"
)

type State int

const (
	inWord State = iota
	inSpace
	inSeq // In a sequence of non word runes
)

func IsWordRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
}

// IntWidth returns number of digits required to print n.
// TODO: Improve speed, not the fastest implementation.
func IntWidth(i int) int {
	return len(fmt.Sprintf("%d", i))
}

// PrevWordStart finds previous word start rune index.
func PrevWordStart(line []rune, idx int) (int, bool) {
	if idx == 0 {
		return 0, false
	}

	state := inSeq
	if idx == len(line) || unicode.IsSpace(line[idx]) {
		state = inSpace
	} else if IsWordRune(line[idx]) {
		state = inWord
	}

	for i := idx - 1; i >= 0; i-- {
		r := line[i]

		switch state {
		case inWord:
			if unicode.IsSpace(r) {
				if i < idx-1 {
					return i + 1, true
				}
				state = inSpace
			} else if !IsWordRune(r) {
				if i < idx-1 {
					return i + 1, true
				}
				state = inSeq
			}
		case inSpace:
			if IsWordRune(r) {
				state = inWord
			} else if !unicode.IsSpace(r) {
				state = inSeq
			}
		case inSeq:
			if IsWordRune(r) || unicode.IsSpace(r) {
				if i < idx-1 {
					return i + 1, true
				}
				state = inWord
			} else if unicode.IsSpace(r) {
				if i < idx-1 {
					return i + 1, true
				}
				state = inSpace
			}
		}
	}

	if state == inWord || state == inSeq {
		return 0, true
	}

	return 0, false
}

// WordEnd finds next word end rune index.
// The returned idx is the index of the first rune not belonging to the word.
func WordEnd(line []rune, idx int) (int, bool) {
	if idx >= len(line) {
		return 0, false
	}

	state := inSeq
	if IsWordRune(line[idx]) {
		state = inWord
	} else if unicode.IsSpace(line[idx]) {
		state = inSpace
	}

	for i := idx + 1; i < len(line); i++ {
		r := line[i]

		switch state {
		case inWord:
			if !IsWordRune(r) {
				return i, true
			}
		case inSpace:
			if IsWordRune(r) {
				state = inWord
			} else if !unicode.IsSpace(r) {
				state = inSeq
			}
		case inSeq:
			if IsWordRune(r) || unicode.IsSpace(r) {
				return i, true
			}
		}
	}

	if state == inWord || state == inSeq {
		return len(line), true
	}

	return 0, false
}

// WordStart finds next word start rune index.
func WordStart(line []rune, idx int) (int, bool) {
	if idx >= len(line)-1 {
		return 0, false
	}

	if idx < 0 {
		if !unicode.IsSpace(line[0]) {
			return 0, true
		}
		idx = 0
	}

	state := inSeq
	if IsWordRune(line[idx]) {
		state = inWord
	} else if unicode.IsSpace(line[idx]) {
		state = inSpace
	}

	for i := idx + 1; i < len(line); i++ {
		r := line[i]

		switch state {
		case inWord:
			if unicode.IsSpace(r) {
				state = inSpace
			} else if !IsWordRune(r) {
				return i, true
			}
		case inSpace:
			if !unicode.IsSpace(r) {
				return i, true
			}
		case inSeq:
			if IsWordRune(r) {
				return i, true
			} else if unicode.IsSpace(r) {
				state = inSpace
			}
		}
	}

	return len(line), false
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
		return ""
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
