package sort

import (
	"testing"
)

func TestFindLongestChain(t *testing.T) {
	testCases := [][][]int{
		[][]int{
			[]int{1, 2},
			[]int{2, 3},
		},
		[][]int{
			[]int{1, 2},
			[]int{3, 4},
		},
		[][]int{
			[]int{1, 2},
			[]int{1, 3},
			[]int{3, 4},
			[]int{3, 5},
			[]int{7, 8},
		},
	}

	results := []int{
		1,
		2,
		3,
	}

	for i, testCase := range testCases {
		result := findLongestChain(testCase)
		if result != results[i] {
			t.Fatalf("TestFindLongestChain: %v %v", testCase, result)
		}
	}
}
