package frame

import (
	"testing"
)

func TestWithin(t *testing.T) {
	var tests = []struct {
		frame Frame
		x     int
		y     int
		want  bool
	}{
		{Frame{nil, 0, 0, 1, 1}, 0, 0, true},
		{Frame{nil, 1, 1, 5, 5}, 0, 0, false},
		{Frame{nil, 1, 1, 5, 5}, 1, 0, false},
		{Frame{nil, 1, 1, 5, 5}, 0, 1, false},
		{Frame{nil, 1, 1, 5, 5}, 1, 1, true},
		{Frame{nil, 1, 1, 5, 5}, 6, 6, false},
		{Frame{nil, 1, 1, 5, 5}, 6, 7, false},
		{Frame{nil, 1, 1, 5, 5}, 7, 6, false},
		{Frame{nil, 1, 1, 5, 5}, 3, 3, true},
	}

	for i, test := range tests {
		got := test.frame.Within(test.x, test.y)
		if got != test.want {
			t.Fatalf("%d: got: %t, want %t", i, got, test.want)
		}
	}
}
