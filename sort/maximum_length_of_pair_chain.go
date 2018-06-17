package sort

import (
	"sort"
)

type bySec [][]int
func (bs bySec) Len() int { return len(bs) }
func (bs bySec) Less(i, j int) bool { return bs[i][1] < bs[j][1] }
func (bs bySec) Swap(i, j int) { bs[i], bs[j] = bs[j], bs[i] }

func findLongestChain(pairs [][]int) int {
	if len(pairs) == 0 {
		return 0
	}

	sorted := bySec(pairs)
	sort.Sort(sorted)

	chain := []int{}
	chain = append(chain, sorted[0][1])
	for _, pair := range sorted {
		if chain[len(chain) - 1] < pair[0] {
			chain = append(chain, pair[1])
		}
	}

	return len(chain)
}