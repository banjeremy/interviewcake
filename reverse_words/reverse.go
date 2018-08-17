package reverse

func ReverseWords(s []rune) []rune {
	// first reverse all characters
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}

	leftIndex := 0
	// then reverse each word
	for rightIndex := 0; rightIndex < len(s); rightIndex++ {
		if rightIndex == len(s)-1 || s[rightIndex+1] == ' ' { // end of word
			word := s[leftIndex : rightIndex+1]

			// reverse each letter of the word
			for i := 0; i < len(word)/2; i++ {
				word[i], word[len(word)-1-i] = word[len(word)-1-i], word[i]
			}
			leftIndex = rightIndex + 2 // skip over spaces
		}
	}

	return s
}
