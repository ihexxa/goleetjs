package tree

import (
	"fmt"
	"testing"
)

func TestReplaceWords(t *testing.T) {
	type Test struct {
		dict []string
		sentence string
		result string
	}

	testCases := []Test{
		Test{
			dict: []string{
				"just",
				"do",
				"it",
			},
			sentence: "justice doing iterating",
			result: "just do it",
		},
		Test{
			dict: []string{
				"abc",
				"abcd",
			},
			sentence: "abca abcde",
			result: "abc abc",
		},
	}

	for _, testCase := range testCases {
		result := replaceWords(testCase.dict, testCase.sentence)
		if result != testCase.result {
			t.Fatalf("\n%v %v -> %v want: %v", testCase.dict, testCase.sentence, result, testCase.result)
		} else {
			fmt.Printf("\n%v %v -> %v %v", testCase.dict, testCase.sentence, result, testCase.result)
		}
	}
}