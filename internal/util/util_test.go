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

func TestWordStart(t *testing.T) {
	var tests = []struct {
		str      string
		startIdx int
		wantIdx  int
		wantOk   bool
	}{
		{"foo", 0, 0, false},
		{"    ", 3, 0, false},
		{"Hello World!", 6, 0, true},
		{"a-b", 2, 0, true},
		{"foo + bar", 6, 0, true},
		{"abc def  agh", 9, 4, true},
		{"aa bb_cc dd", 9, 3, true},
	}

	for _, test := range tests {
		idx, ok := WordStart(test.str, test.startIdx)
		if idx != test.wantIdx || ok != test.wantOk {
			t.Fatalf(
				"str: \"%s\", startIdx: %d, ok: %t, want idx: %d, want ok: %t",
				test.str, test.startIdx, ok, test.wantIdx, test.wantOk,
			)
		}
	}
}

func TestWordEnd(t *testing.T) {
	var tests = []struct {
		str      string
		startIdx int
		wantIdx  int
		wantOk   bool
	}{
		{"    ", 1, 0, false},
		{"abc", 3, 0, false},
		{"foo", 0, 2, true},
		{"aa bb_cc dd", 2, 7, true},
	}

	for _, test := range tests {
		idx, ok := WordEnd(test.str, test.startIdx)
		if idx != test.wantIdx || ok != test.wantOk {
			t.Fatalf(
				"str: \"%s\", startIdx: %d, ok: %t, want idx: %d, want ok: %t",
				test.str, test.startIdx, ok, test.wantIdx, test.wantOk,
			)
		}
	}
}
