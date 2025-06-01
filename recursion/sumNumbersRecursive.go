package recursion

func sumNumbersRecursive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	return nums[0] + sumNumbersRecursive(nums[1:])
}
