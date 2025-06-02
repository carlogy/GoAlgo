package recursion

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{
			n:        3,
			expected: 6,
		}, {
			n:        6,
			expected: 720,
		}, {
			n:        18,
			expected: 6402373705728000,
		}, {
			n:        1,
			expected: 1,
		}, {
			n:        13,
			expected: 6227020800,
		}, {
			n:        0,
			expected: 1,
		},
	}

	for _, test := range tests {
		got := factorial(test.n)

		if got != test.expected {
			t.Errorf("For:\t%d\nGot:\t%d\tExpected\t%d\t", test.n, got, test.expected)
		}
	}
}
