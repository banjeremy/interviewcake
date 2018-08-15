package hical

import (
	"reflect"
	"testing"
)

func TestMergeRanges(t *testing.T) {
	testCases := []struct {
		desc  string
		given []Range
		want  []Range
	}{
		{
			desc: "merges ranges",
			given: []Range{
				{0, 1},
				{3, 5},
				{4, 8},
				{10, 12},
				{9, 10},
			},
			want: []Range{
				{0, 1},
				{3, 8},
				{9, 12},
			},
		},
		{
			desc: "merges touching ranges",
			given: []Range{
				Range{1, 2},
				Range{2, 3},
			},
			want: []Range{
				Range{1, 3},
			},
		},
		{
			desc: "merges subsumed ranges",
			given: []Range{
				Range{1, 5},
				Range{2, 3},
			},
			want: []Range{
				Range{1, 5},
			},
		},
		{
			desc: "merges multiple subsumed ranges",
			given: []Range{
				Range{1, 10},
				Range{2, 6},
				Range{3, 5},
				Range{7, 9},
			},
			want: []Range{
				Range{1, 10},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := MergeRanges(tC.given)
			if !reflect.DeepEqual(tC.want, got) {
				t.Errorf("want: %v, got: %v", tC.want, got)
			}
		})
	}
}
