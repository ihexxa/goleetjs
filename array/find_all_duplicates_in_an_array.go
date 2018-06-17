package array

func findDuplicates(nums []int) []int {
	dups := []int{}

	for _, num := range nums {
		realNum := num
		if num < 0 {
			realNum = -num;
		}

		if nums[realNum-1] > 0 {
			nums[realNum-1] *= -1
		} else {
			dups = append(dups, realNum)
		}
	}

	return dups
}