package reverse

func Reverse(s []rune) []rune {

	for left := 0; left < len(s)/2; left++ {
		right := len(s) - 1 - left
		s[left], s[right] = s[right], s[left]
	}

	return s
}
