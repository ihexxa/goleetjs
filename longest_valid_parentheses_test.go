package goleetjs

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	tests := map[string]int{
		"())":           2,
		"()":            2,
		")(":            0,
		"":              0,
		"(()()":         4,
		")()()()((()))": 12,
	}

	for test, expect := range tests {
		result := longestValidParentheses(test)
		if expect != result {
			t.Errorf("test failed: %s expect %d, result, %d", test, expect, result)
		}
	}
}
