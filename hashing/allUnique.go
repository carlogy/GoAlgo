package hashing

func allUnique[T int | string](s []T) bool {

	if len(s) == 0 || len(s) == 1 {
		return true
	}

	unique := make(map[T]bool)

	for _, item := range s {
		_, ok := unique[item]
		if !ok {
			unique[item] = true
			continue
		}

		return false
	}

	return true
}
