package cmd

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	var tests = []struct {
		line string
		want Command
	}{
		{
			line: "10",
			want: Command{RepCount: 1, Name: "go", Args: []string{"10"}},
		},
		{
			line: "10:12",
			want: Command{RepCount: 1, Name: "go", Args: []string{"10:12"}},
		},
		{
			line: "go 10:12",
			want: Command{RepCount: 1, Name: "go", Args: []string{"10:12"}},
		},
		{
			line: "10 12",
			want: Command{RepCount: 1, Name: "go", Args: []string{"10", "12"}},
		},
		{
			line: "down",
			want: Command{RepCount: 1, Name: "down", Args: nil},
		},
		{
			line: "10 up",
			want: Command{RepCount: 10, Name: "up", Args: nil},
		},
		{
			line: "1 right",
			want: Command{RepCount: 1, Name: "right", Args: nil},
		},
		{
			line: "replace abc def",
			want: Command{RepCount: 1, Name: "replace", Args: []string{"abc", "def"}},
		},
	}

	for _, test := range tests {
		got, err := Parse(test.line)

		if err != nil {
			t.Fatalf("unexpected error %v for line %q", err, test.line)
		}

		want := test.want
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("\nline: %s\ngot:  %+v\nwant: %+v", test.line, got, want)
		}

	}
}
