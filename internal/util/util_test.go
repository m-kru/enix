package util

import (
	"testing"
)

func TestFileNameToType(t *testing.T) {
	var tests = []struct {
		name string
		want string
	}{
		{"makefile", "make"},
		{"Makefile", "make"},
		{"Dockerfile", "docker"},
		{"main.c", "c"},
		{"main.C", "c"},
		{"c", ""},
		{"main.go", "go"},
		{"go", ""},
		{"main.rs", "rust"},
		{"rs", ""},
		{"test.py", "python"},
		{"cfg.json", "json"},
	}

	for _, test := range tests {
		got := FileNameToType(test.name)
		if got != test.want {
			t.Fatalf("name: %s, got: %s, want %s", test.name, got, test.want)
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
		{[]rune("foo"), 0, 2, true},
		{[]rune("aa bb_cc dd"), 2, 7, true},
	}

	for _, test := range tests {
		idx, ok := WordEnd(test.line, test.startIdx)
		if idx != test.wantIdx || ok != test.wantOk {
			t.Fatalf(
				"str: \"%s\", startIdx: %d, ok: %t, want idx: %d, want ok: %t",
				string(test.line), test.startIdx, ok, test.wantIdx, test.wantOk,
			)
		}
	}
}

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
		{[]rune("a-b"), 2, 0, true},
		{[]rune("foo + bar"), 6, 0, true},
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
