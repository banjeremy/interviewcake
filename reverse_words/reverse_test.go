package reverse

import (
	"reflect"
	"testing"
)

func TestReverseWords(t *testing.T) {
	given := []rune{
		'c', 'a', 'k', 'e', ' ',
		'p', 'o', 'u', 'n', 'd', ' ',
		's', 't', 'e', 'a', 'l',
	}

	want := []rune{
		's', 't', 'e', 'a', 'l', ' ',
		'p', 'o', 'u', 'n', 'd', ' ',
		'c', 'a', 'k', 'e',
	}

	got := ReverseWords(given)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %s, got: %s", string(want), string(got))
	}
}
