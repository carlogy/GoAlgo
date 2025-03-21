package introduction

import "testing"

func TestLongestWord(t *testing.T) {
	tests := []struct {
		sentence string
		expected string
	}{
		{
			sentence: "what a wonderful world",
			expected: "wonderful",
		}, {
			sentence: "have a nice day",
			expected: "nice",
		}, {
			sentence: "the quick brown fox jumped over the lazy dog",
			expected: "jumped",
		}, {
			sentence: "who did eat the ham",
			expected: "ham",
		}, {
			sentence: "potato",
			expected: "potato",
		},
	}

	for _, test := range tests {
		got := longestWord(test.sentence)

		if got != test.expected {
			t.Errorf("Got:\t%s\tExpected:\t%s", got, test.expected)
		}
	}

}
