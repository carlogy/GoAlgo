package hashing

func intersection[T int | string](a, b []T) []T {

	count := make(map[T]int)

	interSlice := make([]T, 0)

	for i, v := range b {

		_, ok := count[v]
		if !ok {
			count[v] = i
		}
	}

	for _, v := range a {
		_, ok := count[v]
		if ok {
			interSlice = append(interSlice, v)
		}
	}

	return interSlice

}
