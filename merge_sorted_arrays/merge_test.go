package merge

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		s1   []int
		s2   []int
		want []int
	}{
		{
			desc: "merges slices of the same length",
			s1:   []int{3, 4, 6, 10, 11, 15},
			s2:   []int{1, 5, 8, 12, 14, 19},
			want: []int{1, 3, 4, 5, 6, 8, 10, 11, 12, 14, 15, 19},
		},
		{
			desc: "merges slices of the different lengths",
			s1:   []int{6, 10, 11, 15},
			s2:   []int{1, 5, 8, 12, 14, 19, 20},
			want: []int{1, 5, 6, 8, 10, 11, 12, 14, 15, 19, 20},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := mergeSlices(tc.s1, tc.s2)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("wanted: %v, got: %v", tc.want, got)
			}
		})
	}
}
