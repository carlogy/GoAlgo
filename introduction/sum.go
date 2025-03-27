package introduction

func sum(nums []int) int {

	if nums == nil {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	total := 0
	for _, v := range nums {
		total += v
	}

	return total
}
