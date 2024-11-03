package view

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	var tests = []struct {
		idx  int // Test index
		v    View
		v2   View
		want View
	}{
		{idx: 0, v: View{1, 1, 5, 5}, v2: View{1, 1, 5, 5}, want: View{1, 1, 5, 5}},
		{idx: 1, v: View{1, 1, 5, 5}, v2: View{1, 1, 3, 4}, want: View{1, 1, 3, 4}},
		{idx: 2, v: View{1, 1, 5, 5}, v2: View{2, 2, 1, 1}, want: View{2, 2, 1, 1}},
		{idx: 3, v: View{3, 3, 5, 5}, v2: View{1, 1, 3, 3}, want: View{3, 3, 1, 1}},
		{idx: 4, v: View{1, 1, 5, 5}, v2: View{3, 3, 4, 4}, want: View{3, 3, 3, 3}},
		{idx: 5, v: View{8, 8, 6, 6}, v2: View{6, 9, 3, 2}, want: View{8, 9, 1, 2}},
		{idx: 6, v: View{2, 3, 3, 3}, v2: View{3, 4, 4, 4}, want: View{3, 4, 2, 2}},
		{idx: 7, v: View{9, 2, 1, 5}, v2: View{7, 3, 6, 2}, want: View{9, 3, 1, 2}},
		{idx: 8, v: View{15, 2, 3, 2}, v2: View{17, 3, 2, 2}, want: View{17, 3, 1, 1}},
		{idx: 9, v: View{8, 8, 6, 6}, v2: View{11, 9, 4, 4}, want: View{11, 9, 3, 4}},
		{idx: 10, v: View{2, 15, 5, 2}, v2: View{3, 16, 3, 4}, want: View{3, 16, 3, 1}},
	}

	for i, test := range tests {
		if test.idx != i {
			t.Fatalf("invalid test index, got %d, want %d", test.idx, i)
		}

		got := test.v.Intersection(test.v2)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("\n%d:\ngot:\n%+v\nwant:\n%+v\n", i, got, test.want)
		}
	}
}
