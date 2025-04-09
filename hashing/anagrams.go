package hashing

func anagrams(s1, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	s1map := make(map[string]int)
	s2map := make(map[string]int)

	for _, char := range s1 {
		_, ok := s1map[string(char)]
		if !ok {
			s1map[string(char)] = 0
		}

		s1map[string(char)]++
	}

	for _, char := range s2 {
		_, ok := s2map[string(char)]
		if !ok {
			s2map[string(char)] = 0
		}

		s2map[string(char)]++
	}

	for char, num := range s1map {
		num2, ok := s2map[char]
		if !ok {
			return false
		}

		if num2 != num {
			return false
		}
	}

	return true
}
