package util

import (
	"testing"
)

func TestPrevWordStart(t *testing.T) {
	var tests = []struct {
		line     []rune
		startIdx int
		wantIdx  int
		wantOk   bool
	}{
		{[]rune("foo"), 0, 0, false},
		{[]rune("    "), 3, 0, false},
		{[]rune("Hello World!"), 6, 0, true},
		{[]rune("a-b"), 2, 1, true},
		{[]rune("foo + bar"), 6, 4, true},
		{[]rune("abc def  agh"), 9, 4, true},
		{[]rune("aa bb_cc dd"), 9, 3, true},
	}

	for _, test := range tests {
		idx, ok := PrevWordStart(test.line, test.startIdx)
		if idx != test.wantIdx || ok != test.wantOk {
			t.Fatalf(
				"str: \"%s\", startIdx: %d, ok: %t, want idx: %d, want ok: %t",
				string(test.line), test.startIdx, ok, test.wantIdx, test.wantOk,
			)
		}
	}
}

func TestWordEnd(t *testing.T) {
	var tests = []struct {
		line     []rune
		startIdx int
		wantIdx  int
		wantOk   bool
	}{
		{[]rune("    "), 1, 0, false},
		{[]rune("abc"), 3, 0, false},
		{[]rune("foo"), 0, 3, true},
		{[]rune("aa bb_cc dd"), 2, 8, true},
		{[]rune("aa bb"), 1, 2, true},
	}

	for _, test := range tests {
		idx, ok := WordEnd(test.line, test.startIdx)
		if idx != test.wantIdx || ok != test.wantOk {
			t.Fatalf(
				"str: \"%s\", startIdx: %d, got idx: %d, ok: %t, want idx: %d, want ok: %t",
				string(test.line), test.startIdx, idx, ok, test.wantIdx, test.wantOk,
			)
		}
	}
}

func TestWordStart(t *testing.T) {
	var tests = []struct {
		line     []rune
		startIdx int
		wantIdx  int
		wantOk   bool
	}{
		{[]rune("    "), 1, 0, false},
		{[]rune("  abc"), 0, 2, true},
		{[]rune("  abc def"), 2, 6, true},
		{[]rune("  def12"), 4, 0, false},
	}

	for _, test := range tests {
		idx, ok := WordStart(test.line, test.startIdx)
		if idx != test.wantIdx || ok != test.wantOk {
			t.Fatalf(
				"str: \"%s\", startIdx: %d, ok: %t, want idx: %d, want ok: %t",
				string(test.line), test.startIdx, ok, test.wantIdx, test.wantOk,
			)
		}
	}
}

func TestGetWord(t *testing.T) {
	var tests = []struct {
		line []rune
		idx  int
		want string
	}{
		{[]rune("lorem ipsum"), 0, "lorem"},
		{[]rune("lorem ipsum"), 2, "lorem"},
		{[]rune("lorem ipsum"), 4, "lorem"},
		{[]rune("lorem ipsum"), 6, "ipsum"},
		{[]rune("lorem ipsum"), 10, "ipsum"},
		{[]rune(" {} "), 1, ""},
		{[]rune("[]"), 1, ""},
		{[]rune(" "), 0, ""},
		{[]rune("a\tb"), 1, ""},
		{[]rune("1+2"), 1, ""},
	}

	for _, test := range tests {
		got := GetWord(test.line, test.idx)
		if got != test.want {
			t.Fatalf(
				"str: \"%s\", got: %s, want: %s",
				string(test.line), got, test.want,
			)
		}
	}
}
