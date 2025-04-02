package introduction

import (
	"slices"
)

func median(nums []int) float64 {
	total := len(nums)

	if total == 1 {
		return float64(nums[0])
	}

	if total == 0 {
		return 0
	}

	slices.Sort(nums)

	if total%2 == 0 {
		mid1 := float64(nums[total/2])
		mid2 := float64(nums[total/2-1])
		middleAverage := (mid1 + mid2) / 2
		return middleAverage

	} else {

		med := float64(nums[(total)/2])
		return med
	}

}
