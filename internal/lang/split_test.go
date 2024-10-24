package lang

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/m-kru/enix/internal/line"
)

var regions = []*Region{
	&Region{
		Name: "Default",
	},
	&Region{
		Name:        "Line comment",
		StartRegexp: regexp.MustCompile(`//`),
		EndRegexp:   regexp.MustCompile(`$`),
	},
	&Region{
		Name:        "Block comment",
		StartRegexp: regexp.MustCompile(`/\*`),
		EndRegexp:   regexp.MustCompile(`\*/`),
	},
	&Region{
		Name:        "String",
		StartRegexp: regexp.MustCompile(`"`),
		EndRegexp:   regexp.MustCompile(`"`),
	},
}

func TestSplit(t *testing.T) {
	var tests = []struct {
		idx       int
		text      string
		startLine int
		endLine   int
		want      []Section
	}{
		{
			idx:       0,
			text:      `int main(int argc, char *argv[])`,
			startLine: 1,
			endLine:   1,
			want: []Section{
				Section{
					StartLine: 1, StartIdx: 0, EndLine: 1, EndIdx: 32, Region: regions[0],
				},
			},
		},
		{
			idx:       1,
			text:      `int N = 5; // List size`,
			startLine: 1,
			endLine:   1,
			want: []Section{
				Section{
					StartLine: 1, StartIdx: 0, EndLine: 1, EndIdx: 11, Region: regions[0],
				},
				Section{
					StartLine: 1, StartIdx: 11, EndLine: 1, EndIdx: 23, Region: regions[1],
				},
			},
		},
		{
			idx:       2,
			text:      `// Line comment`,
			startLine: 1,
			endLine:   1,
			want: []Section{
				Section{
					StartLine: 1, StartIdx: 0, EndLine: 1, EndIdx: 15, Region: regions[1],
				},
			},
		},
		{
			idx:       3,
			text:      `int a; /* */ //`,
			startLine: 1,
			endLine:   1,
			want: []Section{
				Section{
					StartLine: 1, StartIdx: 0, EndLine: 1, EndIdx: 7, Region: regions[0],
				},
				Section{
					StartLine: 1, StartIdx: 7, EndLine: 1, EndIdx: 12, Region: regions[2],
				},
				Section{
					StartLine: 1, StartIdx: 12, EndLine: 1, EndIdx: 13, Region: regions[0],
				},
				Section{
					StartLine: 1, StartIdx: 13, EndLine: 1, EndIdx: 15, Region: regions[1],
				},
			},
		},
		{
			idx:       4,
			text:      `"a""b"`,
			startLine: 1,
			endLine:   1,
			want: []Section{
				Section{
					StartLine: 1, StartIdx: 0, EndLine: 1, EndIdx: 3, Region: regions[3],
				},
				Section{
					StartLine: 1, StartIdx: 3, EndLine: 1, EndIdx: 6, Region: regions[3],
				},
			},
		},
	}

	for i, test := range tests {
		if test.idx != i {
			t.Fatalf("invalid test index, got %d, want %d", test.idx, i)
		}

		lines := line.FromString(test.text)
		secs, line := splitIntoSections(regions, lines, test.startLine, test.endLine)

		if !reflect.DeepEqual(secs, test.want) {
			t.Fatalf(
				"test %d:\ntext:\n%s\ngot:\n%+v\nwant:\n%+v\n",
				i, test.text, secs, test.want,
			)
		}

		if line != lines.Get(test.startLine) {
			t.Fatalf("invalid start line returned, got %d, want %d", line.Num(), test.startLine)
		}
	}
}
