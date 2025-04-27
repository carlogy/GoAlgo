package hashing

func pairSum(nums []int, target int) []int {

	sumCount := make(map[int]int)

	for i, num := range nums {

		complement := target - num

		ind, ok := sumCount[complement]
		if !ok {
			sumCount[num] = i
			continue
		}

		return []int{ind, i}
	}

	return nil
}
