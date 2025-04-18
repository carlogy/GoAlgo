package hashing

func mostFrequentChar(s string) string {

	if len(s) == 0 {
		return ""
	}

	charCount := make(map[string]int)

	for _, char := range s {
		_, ok := charCount[string(char)]
		if !ok {
			charCount[string(char)] = 0
		}

		charCount[string(char)]++
	}

	type largest struct {
		char  string
		count int
	}

	l := largest{"", 0}

	for _, char := range s {

		count := charCount[string(char)]

		if l.count < count {
			l.char = string(char)
			l.count = count
		}
	}

	return l.char
}
