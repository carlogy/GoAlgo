package hashing

func intersectionWithDupes[T string | int](a []T, b []T) []T {

	countA := make(map[T]int)
	countB := make(map[T]int)

	result := make([]T, 0)

	for _, item := range a {

		_, ok := countA[item]
		if !ok {
			countA[item] = 0
		}

		countA[item]++
	}

	for _, item := range b {
		_, ok := countB[item]
		if !ok {
			countB[item] = 0
		}

		countB[item]++
	}

	for key, count := range countA {
		total, ok := countB[key]
		if ok {
			minC := min(count, total)

			if minC > 0 {
				for i := 0; i < minC; i++ {
					result = append(result, key)
				}
			}
		}
	}

	return result

}
