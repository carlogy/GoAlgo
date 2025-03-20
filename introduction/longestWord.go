package introduction

import "strings"

func longestWord(sentence string) string {
	wordSlice := strings.Split(sentence, " ")
	longest := ""

	for _, word := range wordSlice {

		if len(word) >= len(longest) {
			longest = word
		}
	}

	return longest

}
