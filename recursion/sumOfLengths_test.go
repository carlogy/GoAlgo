package recursion

import "testing"

func TestSumOfLengths(t *testing.T) {
	tests := []struct {
		s        []string
		expected int
	}{
		{
			s:        []string{"goat", "cat", "purple"},
			expected: 13,
		}, {
			s:        []string{"bike", "at", "pencils", "phone"},
			expected: 18,
		}, {
			s:        []string{},
			expected: 0,
		}, {
			s:        []string{"", " ", "  ", "   ", "    ", "     "},
			expected: 15,
		}, {
			s:        []string{"0", "313", "1234567890"},
			expected: 14,
		},
	}

	for _, test := range tests {
		got := sumOfLengths(test.s)

		if got != test.expected {
			t.Errorf("For:\t%v\nGot:\t%d\tExpected:\t%d", test.s, got, test.expected)
		}
	}
}
