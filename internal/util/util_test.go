package util

import (
	"testing"
)

func TestFileNameToType(t *testing.T) {
	var tests = []struct {
		name string
		want string
	}{
		{"COMMIT_EDITMSG", "git-commit"},
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
