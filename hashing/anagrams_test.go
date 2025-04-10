package hashing

import "testing"

func TestAnagrams(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{
			s1:       "restful",
			s2:       "fluster",
			expected: true,
		}, {
			s1:       "cats",
			s2:       "tocs",
			expected: false,
		}, {
			s1:       "monkeyswrite",
			s2:       "newyorktimes",
			expected: true,
		}, {
			s1:       "paper",
			s2:       "reapa",
			expected: false,
		}, {
			s1:       "elbow",
			s2:       "below",
			expected: true,
		}, {
			s1:       "tax",
			s2:       "taxi",
			expected: false,
		}, {
			s1:       "taxi",
			s2:       "tax",
			expected: false,
		}, {
			s1:       "night",
			s2:       "thing",
			expected: true,
		}, {
			s1:       "abbc",
			s2:       "aabc",
			expected: false,
		}, {
			s1:       "po",
			s2:       "popp",
			expected: false,
		}, {
			s1:       "pp",
			s2:       "oo",
			expected: false,
		},
	}

	for _, test := range tests {
		got := anagrams(test.s1, test.s2)

		if got != test.expected {
			t.Errorf("For:\t%s->\t%s\nGot:\t%t\tExpected:\t%t\t", test.s1, test.s2, got, test.expected)
		}
	}
}
