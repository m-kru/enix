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
		idx      int // Test index
		str      string
		col      int
		tabWidth int
		want     Want
	}{
		{
			idx:      0,
			str:      "\tdef",
			col:      1,
			tabWidth: 4,
			want:     Want{0, 0, true},
		},
		{
			idx:      1,
			str:      "\tfoo",
			col:      2,
			tabWidth: 4,
			want:     Want{0, 1, true},
		},
		{
			idx:      2,
			str:      "\ta\tb",
			col:      9,
			tabWidth: 4,
			want:     Want{3, 0, true},
		},
		{
			idx:      3,
			str:      "\tif",
			col:      5,
			tabWidth: 5,
			want:     Want{0, 4, true},
		},
		{
			idx:      4,
			str:      "\t\ta",
			col:      13,
			tabWidth: 6,
			want:     Want{2, 0, true},
		},
		{
			idx:      5,
			str:      "\tab\ta",
			col:      13,
			tabWidth: 6,
			want:     Want{4, 0, true},
		},
		{
			idx:      6,
			str:      "\t\ta",
			col:      3,
			tabWidth: 1,
			want:     Want{2, 0, true},
		},
		{
			idx:      7,
			str:      "\ta\ta",
			col:      4,
			tabWidth: 1,
			want:     Want{3, 0, true},
		},
		{
			idx:      8,
			str:      "世界",
			col:      4,
			tabWidth: 8,
			want:     Want{1, 1, true},
		},
		{
			idx:      9,
			str:      "世a界",
			col:      4,
			tabWidth: 8,
			want:     Want{2, 0, true},
		},
		{
			idx:      10,
			str:      "a界世",
			col:      5,
			tabWidth: 8,
			want:     Want{2, 1, true},
		},
	}

	for _, test := range tests {
		line, _ := FromString(test.str)

		runeIdx, runeSubcol, ok := line.RuneIdx(test.col, test.tabWidth)
		want := test.want
		if runeIdx != want.runeIdx || runeSubcol != want.runeSubcol || ok != want.ok {
			t.Fatalf(
				"%d:%s:\nruneIdx: %d, runeSubcol: %d, ok: %t, want.runeIdx: %d, want.runeSubcol: %d, want.ok: %t",
				test.idx, test.str, runeIdx, runeSubcol, ok, want.runeIdx, want.runeSubcol, want.ok,
			)
		}
	}
}
