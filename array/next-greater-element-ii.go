package array

func nextGreaterElements(nums []int) []int {
	nums2 := make([]int, 0)
	nums2 = append(nums2, nums...)
	nums2 = append(nums2, nums...)
	result := make([]int, len(nums2))
	stack := make([]int, 0)
	
	for i:=0; i<len(nums2); i++ { // out of index...
		if len(stack) == 0 || nums2[stack[len(stack)-1]] >= nums2[i] {
			stack = append(stack, i)
		} else {
			for len(stack) != 0 && nums2[stack[len(stack)-1]] < nums2[i] {
				result[stack[len(stack)-1]] = nums2[i] // should be number
				stack = stack[0:len(stack)-1]
			}
			stack = append(stack, i)
		}
	}

	for len(stack) > 0 { // empty stack
		result[stack[len(stack)-1]] = -1
		stack = stack[0:len(stack)-1]
	}
	return result[0:len(nums)]
}