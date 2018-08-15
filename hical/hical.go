package hical

import "sort"

type Range struct {
	Start int
	End   int
}

type byStartTime []Range

func (s byStartTime) Len() int {
	return len(s)
}

func (s byStartTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byStartTime) Less(i, j int) bool {
	return s[i].Start < s[j].Start
}

func MergeRanges(ranges []Range) []Range {
	// sort the slice by start time
	sort.Sort(byStartTime(ranges))

	var merged []Range
	curr := ranges[0]

	for i := 1; i < len(ranges); i++ {
		next := ranges[i]

		// if next.end < curr.end, next gets absorbed (skip it)
		if next.End < curr.End {
			continue
		} else if next.Start <= curr.End {
			// else if next.start <= curr.end, curr.end = next.end
			curr.End = next.End
		} else {
			merged = append(merged, curr)
			curr = next
		}
	}

	merged = append(merged, curr)

	return merged
}
