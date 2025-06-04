package recursion

func sumOfLengths(s []string) int {
	if len(s) == 0 {
		return 0
	}

	return len(s[0]) + sumOfLengths(s[1:])
}
