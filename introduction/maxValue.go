package introduction

import "math"

func maxValue(nums []int) int {
	max := math.MinInt

	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}
