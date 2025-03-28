package introduction

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{7, 4, 3, 100, 2343243, 343434, 1, 2, 32},
			expected: 2686826,
		}, {
			nums:     []int{12, 12, 12},
			expected: 36,
		}, {
			nums:     []int{10, 200, 3000, 5000, 4},
			expected: 8214,
		}, {
			nums:     nil,
			expected: 0,
		}, {
			nums:     []int{1},
			expected: 1,
		}, {
			nums:     []int{123456789},
			expected: 123456789,
		}, {
			nums:     []int{-1, -2, -3},
			expected: -6,
		}, {
			nums:     []int{0, 0, 0, 0, 0},
			expected: 0,
		},
	}

	for _, test := range tests {
		got := sum(test.nums)

		if got != test.expected {
			t.Errorf("Got:\t%d\tExpected:\t%d", got, test.expected)
		}
	}
}
