package hashing

func exclusiveItems[T int | string](a, b []T) []T {

	bMap := make(map[T]bool)
	aMap := make(map[T]bool)
	result := make([]T, 0)

	for _, item := range b {

		_, ok := bMap[item]
		if !ok {
			bMap[item] = true
		}
	}

	for _, item := range a {
		_, ok := aMap[item]
		if !ok {
			aMap[item] = true
		}
	}

	for key, _ := range bMap {

		_, ok := aMap[key]
		if !ok {
			result = append(result, key)
		}
	}

	for key, _ := range aMap {
		_, ok := bMap[key]
		if !ok {
			result = append(result, key)
		}
	}

	return result
}
