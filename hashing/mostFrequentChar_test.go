package hashing

import "testing"

func TestMostFrequentChar(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{
			s:        "bookkeeper",
			expected: "e",
		}, {
			s:        "david",
			expected: "d",
		}, {
			s:        "abby",
			expected: "b",
		}, {
			s:        "mississippi",
			expected: "i",
		}, {
			s:        "potato",
			expected: "o",
		}, {
			s:        "eleventennine",
			expected: "e",
		}, {
			s:        "riverbed",
			expected: "r",
		},
	}

	for _, test := range tests {

		got := mostFrequentChar(test.s)

		if test.expected != got {
			t.Errorf("For:\t%s\nGot:\t%s\tExpected:\t%s\t", test.s, got, test.expected)
		}
	}
}
