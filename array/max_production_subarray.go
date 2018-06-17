package array

func maxProduct(nums []int) int {
	minProds := make([]int, len(nums))
	maxProds := make([]int, len(nums))

	max := nums[0]
	minProds[0] = nums[0]
	maxProds[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num == 0 {
			minProds[i] = 0
			maxProds[i] = 0
		} else if num > 0 {
			if num > num*minProds[i-1] {
				minProds[i] = num * minProds[i-1]
			} else {
				minProds[i] = num
			}
			if num < num*maxProds[i-1] {
				maxProds[i] = num * maxProds[i-1]
			} else {
				maxProds[i] = num
			}
		} else {
			if num > num*maxProds[i-1] {
				minProds[i] = num * maxProds[i-1]
			} else {
				minProds[i] = num
			}
			if num < num*minProds[i-1] {
				maxProds[i] = num * minProds[i-1]
			} else {
				maxProds[i] = num
			}
		}

		if maxProds[i] > max {
			max = maxProds[i]
		}
	}

	return max
}
