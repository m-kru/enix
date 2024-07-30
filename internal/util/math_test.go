package util

import (
	"testing"
)

func TestNextPowerOfTwo(t *testing.T) {
	var tests = []struct {
		n    int
		want int
	}{
		{7, 8},
		{1, 1},
		{1, 1},
		{33, 64},
		{113, 128},
		{200, 256},
		{512, 512},
		{1023, 1024},
	}

	for _, test := range tests {
		got := NextPowerOfTwo(test.n)
		if got != test.want {
			t.Fatalf("n: %d, got: %d, want: %d", test.n, got, test.want)
		}
	}
}
