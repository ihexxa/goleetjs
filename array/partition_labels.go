package array

func partitionLabels(S string) []int {
	if len(S) == 0 {
		return []int{0}
	} else if len(S) == 1 {
		return []int{1}
	}

	labels := []rune(S)
	rightest := map[rune]int{}
	for i, c := range labels {
		rightest[c] = i
	}

	left := 0
	right := 1
	parts := []int{}
	for i, c := range labels {
		if i == right {
			parts = append(parts, right-left)
			left = right
			right++
		}

		rightestPos := rightest[c]
		if rightestPos + 1 > right {
			right = rightestPos + 1
		}
	}

	if left <= len(labels) -1 {
		parts = append(parts, len(labels)-left)
	}

	return parts
}
