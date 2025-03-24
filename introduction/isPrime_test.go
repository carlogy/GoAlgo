package introduction

import "testing"

func TestIsPrime(t *testing.T) {

	tests := []struct {
		num      int
		expected bool
	}{
		{
			num:      2,
			expected: true,
		}, {
			num:      3,
			expected: true,
		}, {
			num:      4,
			expected: false,
		}, {
			num:      5,
			expected: true,
		}, {
			num:      6,
			expected: false,
		}, {
			num:      7,
			expected: true,
		}, {
			num:      8,
			expected: false,
		}, {
			num:      25,
			expected: false,
		}, {
			num:      31,
			expected: true,
		}, {
			num:      2017,
			expected: true,
		}, {
			num:      2048,
			expected: false,
		}, {
			num:      1,
			expected: false,
		}, {
			num:      713,
			expected: false,
		}, {
			num:      0,
			expected: false,
		},
	}

	for _, test := range tests {
		got := isPrime(test.num)

		if got != test.expected {
			t.Errorf("Num:\t%d\tGot:\t%t\tExpected:\t%t", test.num, got, test.expected)
		}
	}
}
