package introduction

import "math"

func minimumValue(nums []int) int {
	if nums == nil {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	min := math.MaxInt

	for i := 0; i < len(nums); i++ {

		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}
