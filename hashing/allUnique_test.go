package hashing

import "testing"

func TestAllUnique(t *testing.T) {

	tests := []struct {
		s        []string
		expected bool
	}{
		{
			s:        []string{"q", "r", "s", "a"},
			expected: true,
		},
		{
			s:        []string{"q", "r", "s", "a", "r", "z"},
			expected: false,
		}, {
			s:        []string{"red", "blue", "yellow", "green", "orange"},
			expected: true,
		}, {
			s:        []string{"cat", "cat", "dog"},
			expected: false,
		}, {
			s:        []string{"a", "u", "t", "u", "m", "n"},
			expected: false,
		},
	}

	for _, test := range tests {
		got := allUnique(test.s)

		if got != test.expected {
			t.Errorf("For:\t%v\nGot:\t%t\tExpected:\t%t", test.s, got, test.expected)
		}

	}

}
