package array

func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}

	var stack = make([]byte, 0)
	var gaps = make([]int, 0)
	gaps = shrink(s, stack, gaps)
	return countGap(gaps, len(s))
}

func shrink(s string, stack []byte, gaps []int) []int {
	for i := 0; i < len(s); i++ {
		bracket := s[i]
		if topByte(stack) == '(' && bracket == ')' {
			stack = popByte(stack)
			gaps = pop(gaps)
		} else {
			stack = append(stack, bracket)
			gaps = append(gaps, i)
		}
	}
	return gaps
}

func topByte(stack []byte) byte {
	if len(stack) == 0 {
		return '0'
	} else {
		return stack[len(stack)-1]
	}
}

func popByte(stack []byte) []byte {
	if len(stack) == 0 {
		return stack
	}
	return stack[:len(stack)-1]
}

func pop(stack []int) []int {
	if len(stack) == 0 {
		return stack
	}
	return stack[:len(stack)-1]
}

func countGap(gaps []int, length int) int {
	if len(gaps) == 0 {
		return length
	}

	max := 0
	gaps = append([]int{-1}, gaps...)
	gaps = append(gaps, length)
	for i := 1; i < len(gaps); i++ {
		gap := gaps[i] - gaps[i-1] - 1

		if max < gap {
			max = gap
		}
	}
	return max
}
