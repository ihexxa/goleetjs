package goleetjs

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 || k <= 0 {
		return false
	}

	var m = make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		var cur = nums[i]
		if _, ok := m[cur]; ok {
			return true
		}
		m[cur] = cur
		if i-k >= 0 {
			delete(m, nums[i-k])
		}
	}
	return false
}
