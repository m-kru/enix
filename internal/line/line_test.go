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
		line     *Line
		col      int
		tabWidth int
		want     Want
	}{
		{
			idx:      0,
			line:     FromString("\tdef"),
			col:      1,
			tabWidth: 4,
			want:     Want{0, 0, true},
		},
		{
			idx:      1,
			line:     FromString("\tfoo"),
			col:      2,
			tabWidth: 4,
			want:     Want{0, 1, true},
		},
		{
			idx:      2,
			line:     FromString("\ta\tb"),
			col:      9,
			tabWidth: 4,
			want:     Want{3, 0, true},
		},
		{
			idx:      3,
			line:     FromString("\tif"),
			col:      5,
			tabWidth: 5,
			want:     Want{0, 4, true},
		},
		{
			idx:      4,
			line:     FromString("\t\ta"),
			col:      13,
			tabWidth: 6,
			want:     Want{2, 0, true},
		},
		{
			idx:      5,
			line:     FromString("\tab\ta"),
			col:      13,
			tabWidth: 6,
			want:     Want{4, 0, true},
		},
		{
			idx:      6,
			line:     FromString("\t\ta"),
			col:      3,
			tabWidth: 1,
			want:     Want{2, 0, true},
		},
		{
			idx:      7,
			line:     FromString("\ta\ta"),
			col:      4,
			tabWidth: 1,
			want:     Want{3, 0, true},
		},
		{
			idx:      8,
			line:     FromString("世界"),
			col:      4,
			tabWidth: 8,
			want:     Want{1, 1, true},
		},
		{
			idx:      9,
			line:     FromString("世a界"),
			col:      4,
			tabWidth: 8,
			want:     Want{2, 0, true},
		},
		{
			idx:      10,
			line:     FromString("a界世"),
			col:      5,
			tabWidth: 8,
			want:     Want{2, 1, true},
		},
	}

	for _, test := range tests {
		runeIdx, runeSubcol, ok := test.line.RuneIdx(test.col, test.tabWidth)
		want := test.want
		if runeIdx != want.runeIdx || runeSubcol != want.runeSubcol || ok != want.ok {
			t.Fatalf(
				"%d:%s:\nruneIdx: %d, runeSubcol: %d, ok: %t, want.runeIdx: %d, want.runeSubcol: %d, want.ok: %t",
				test.idx, test.line.String(), runeIdx, runeSubcol, ok, want.runeIdx, want.runeSubcol, want.ok,
			)
		}
	}
}

func TestWordEnd(t *testing.T) {
	var tests = []struct {
		line     *Line
		startIdx int
		wantIdx  int
		wantOk   bool
	}{
		{FromString("    "), 1, 0, false},
		{FromString("abc"), 3, 0, false},
		{FromString("foo"), 0, 2, true},
		{FromString("aa bb_cc dd"), 2, 7, true},
	}

	for _, test := range tests {
		idx, ok := test.line.WordEnd(test.startIdx)
		if idx != test.wantIdx || ok != test.wantOk {
			t.Fatalf(
				"str: \"%s\", startIdx: %d, ok: %t, want idx: %d, want ok: %t",
				test.line.String(), test.startIdx, ok, test.wantIdx, test.wantOk,
			)
		}
	}
}

func TestPrevWordStart(t *testing.T) {
	var tests = []struct {
		line     *Line
		startIdx int
		wantIdx  int
		wantOk   bool
	}{
		{FromString("foo"), 0, 0, false},
		{FromString("    "), 3, 0, false},
		{FromString("Hello World!"), 6, 0, true},
		{FromString("a-b"), 2, 0, true},
		{FromString("foo + bar"), 6, 0, true},
		{FromString("abc def  agh"), 9, 4, true},
		{FromString("aa bb_cc dd"), 9, 3, true},
	}

	for _, test := range tests {
		idx, ok := test.line.PrevWordStart(test.startIdx)
		if idx != test.wantIdx || ok != test.wantOk {
			t.Fatalf(
				"str: \"%s\", startIdx: %d, ok: %t, want idx: %d, want ok: %t",
				test.line.String(), test.startIdx, ok, test.wantIdx, test.wantOk,
			)
		}
	}
}
