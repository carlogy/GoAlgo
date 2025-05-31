package hashing

import "fmt"

func pairProduct(nums []int, target int) []int {

	productCount := make(map[int]int)

	for i, num := range nums {
		fmt.Println(i, num)
		var complement int
		if num != 0 {
			complement = target / num
		}

		ind, ok := productCount[complement]
		if !ok {
			productCount[num] = i
			continue
		}

		return []int{ind, i}

	}
	return nil
}
