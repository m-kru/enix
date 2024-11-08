package line

import (
	"testing"
)

func TestRuneIdx(t *testing.T) {
	type Want struct {
		runeIdx    int
		runeSubcol int
		ok         bool
	}

	var tests = []struct {
		idx  int // Test index
		str  string
		col  int
		want Want
	}{
		{
			idx:  0,
			str:  "\tdef",
			col:  1,
			want: Want{0, 0, true},
		},
		{
			idx:  1,
			str:  "\tfoo",
			col:  2,
			want: Want{0, 1, true},
		},
		{
			idx:  2,
			str:  "\ta\tb",
			col:  17,
			want: Want{3, 0, true},
		},
		{
			idx:  3,
			str:  "\tif",
			col:  5,
			want: Want{0, 4, true},
		},
		{
			idx:  4,
			str:  "\t\ta",
			col:  17,
			want: Want{2, 0, true},
		},
		{
			idx:  5,
			str:  "\tab\ta",
			col:  17,
			want: Want{4, 0, true},
		},
		{
			idx:  6,
			str:  "\t\ta",
			col:  17,
			want: Want{2, 0, true},
		},
		{
			idx:  7,
			str:  "世界",
			col:  4,
			want: Want{1, 1, true},
		},
		{
			idx:  8,
			str:  "世a界",
			col:  4,
			want: Want{2, 0, true},
		},
		{
			idx:  9,
			str:  "a界世",
			col:  5,
			want: Want{2, 1, true},
		},
	}

	for _, test := range tests {
		line, _ := FromString(test.str)

		runeIdx, runeSubcol, ok := line.RuneIdx(test.col)
		want := test.want
		if runeIdx != want.runeIdx || runeSubcol != want.runeSubcol || ok != want.ok {
			t.Fatalf(
				"%d:%s:\nruneIdx: %d, runeSubcol: %d, ok: %t, want.runeIdx: %d, want.runeSubcol: %d, want.ok: %t",
				test.idx, test.str, runeIdx, runeSubcol, ok, want.runeIdx, want.runeSubcol, want.ok,
			)
		}
	}
}
