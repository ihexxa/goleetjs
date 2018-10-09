package array

import (
	"fmt"
	"testing"
)

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

func TestMaxProduct(t *testing.T) {
	type testCase struct {
		List   []int
		Result int
	}

	cases := []testCase{
		{
			List:   []int{1, 2, 0, 3},
			Result: 3,
		},
		{
			List:   []int{1, -2, 0, 3},
			Result: 3,
		},
		{
			List:   []int{-1, -2},
			Result: 2,
		},
	}

	for _, test := range cases {
		if maxProduct(test.List) != test.Result {
			t.Fatal("maxProduct failed")
		}
	}
}

func TestMaxIncreaseKeepingSkyline(t *testing.T) {
	testCases := [][][]int{
		[][]int{
			[]int{1, 2},
			[]int{9, 4},
		},
		[][]int{
			[]int{1, 2, 3},
			[]int{9, 4, 7},
			[]int{0, 8, 6},
		},
	}

	rets := []int{
		1,
		2 + 1 + 0 + 0 + 4 + 0 + 8 + 0 + 1,
	}

	for i, testCase := range testCases {
		ret := maxIncreaseKeepingSkyline(testCase)
		if ret != rets[i] {
			t.Fatalf("TestMaxIncreaseKeepingSkyline %v %d", testCase, ret)
		}
	}
}

func TestPartitionLabels(t *testing.T) {
	inputs := []string{
		"ababcbacadefegdehijhklij",
		"ababc",
	}

	outputs := [][]int{
		[]int{9, 7, 8},
		[]int{4, 1},
	}

	sameSlice := func(s1, s2 []int) bool {
		for i, val := range s1 {
			if val != s2[i] {
				return false
			}
		}
		return true
	}

	for i, input := range inputs {
		ret := partitionLabels(input)
		if !sameSlice(ret, outputs[i]) {
			t.Fatalf("partition labels", ret, outputs[i])
		}
	}
}

func TestFourSum(t *testing.T) {
	fmt.Println(fourSum([]int{0, 0, 0, 0}, 0))
}

func TestNextGreaterElements(t *testing.T) {
	fmt.Println("nextGreater", []int{3,7,2,1}, nextGreaterElements([]int{3,7,4,1}))
}