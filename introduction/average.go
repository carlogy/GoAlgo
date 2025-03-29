package introduction

func average(nums []int) int {
	total := len(nums)

	if total == 1 {
		return nums[0]
	}

	if total == 0 {
		return 0
	}

	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum / total
}
