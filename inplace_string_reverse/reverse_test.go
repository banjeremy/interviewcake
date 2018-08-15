package reverse

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		desc  string
		given []rune
		want  []rune
	}{
		{
			desc:  "reverses even-lengthed string",
			given: []rune("abcdef12"),
			want:  []rune("21fedcba"),
		},
		{
			desc:  "reverses odd-lengthed string",
			given: []rune("abcdef123"),
			want:  []rune("321fedcba"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := Reverse(tc.given)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
