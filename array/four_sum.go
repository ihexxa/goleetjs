package array

import (
	"sort"
	"fmt"
)

func noDuplicate(keys1 , keys2 []int) bool {
	set := map[int]bool{}
	for _, val := range append(keys1, keys2...) {
		set[val] = true
	}

	return len(set) == 4
}

// incorrect: [0,0,0,0] 0
func filterDups(results [][]int) [][]int {
	set := map[string][]int{}
	for _, result := range results {
		set[fmt.Sprintf("%d-%d-%d-%d", result[0], result[1], result[2], result[3])] = result
	}
	noDupResults := [][]int{}
	for _, result := range set {
		noDupResults = append(noDupResults, result)
	}
	return noDupResults
}

func fourSum(nums []int, target int) [][]int {
	sum2Combines := make(map[int]map[string][]int, 0)
	sort.Ints(nums)

	for i:=0; i<len(nums); i++ {
		for j:=i+1; j<len(nums); j++ {
			sum := nums[i]+nums[j]
			_, ok := sum2Combines[sum]
			if !ok {
				sum2Combines[sum] = make(map[string][]int, 0)
			}
			combine := []int{i, j}
			sum2Combines[sum][fmt.Sprintf("%d+%d",i, j)] = combine
		}
	}

	results := make([][]int, 0)
	for key1, combines := range sum2Combines {
		for pattern, combine := range combines {
			remain := target - (nums[combine[0]] + nums[combine[1]])
			if remain != key1 {
				remainCombines, ok := sum2Combines[remain]
				if ok {
					for _, rCombine := range remainCombines {
						if noDuplicate(combine, rCombine) {
							result := []int{}
							for _, id := range append(combine, rCombine...) {
								result = append(result, nums[id])
							}
							sort.Ints(result)
							results = append(results, result)
						}
					}
				}
			} else {
				for pattern2, rCombine := range combines {
					if pattern2 != pattern && noDuplicate(combine, rCombine) {
						result := []int{}
						for _, id := range append(combine, rCombine...) {
							result = append(result, nums[id])
						}
						sort.Ints(result)
						results = append(results, result)
					}
				}
			}
		}
	}

	return filterDups(results)
}