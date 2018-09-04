package merge

// https://www.interviewcake.com/question/python/merge-sorted-arrays?section=array-and-string-manipulation&course=fc1

func mergeSlices(s1, s2 []int) []int {
	var merged, longer, shorter []int

	if len(s1) > len(s2) {
		longer = s1
		shorter = s2
	} else {
		longer = s2
		shorter = s1
	}

	var shorterIndex, longerIndex int
	for shorterIndex < len(shorter) {
		if shorter[shorterIndex] < longer[longerIndex] {
			merged = append(merged, shorter[shorterIndex])
			shorterIndex++
		} else {
			merged = append(merged, longer[longerIndex])
			longerIndex++
		}
	}
	return append(merged, longer[longerIndex:]...)
}
