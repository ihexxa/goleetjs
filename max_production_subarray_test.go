
package goleetjs

import (
    "testing"
)

type testCase struct {
    List   []int
    Result int
}
func TestMaxProduct(t *testing.T) {
    cases := []testCase{
        {
            List: []int{1, 2, 0, 3},
            Result: 3,
        },
        {
            List: []int{1, -2, 0, 3},
            Result: 3,
        },
        {
            List: []int{-1, -2},
            Result: 2,
        },
    }

    for _, test := range cases {
        if maxProduct(test.List) != test.Result {
            t.Fatal("maxProduct failed")
        }
    }
}
