package introduction

import (
	"testing"
)

func TestMedian(t *testing.T) {
	tests := []struct {
		nums     []int
		expected float64
	}{
		{
			nums:     []int{12, 12, 12},
			expected: 12,
		}, {
			nums:     []int{10, 200, 3000, 5000, 4},
			expected: 200,
		}, {
			nums:     []int{7, 4, 3, 100, 2343243, 343434, 1, 2, 32},
			expected: 7,
		}, {
			nums:     nil,
			expected: 0,
		}, {
			nums:     []int{0},
			expected: 0,
		}, {
			nums:     []int{1000000},
			expected: 1000000,
		},
		{
			nums:     []int{5, 2, 3, 7, 6, 4},
			expected: 4.5,
		}, {
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: 5.5,
		},
	}

	for _, test := range tests {
		got := median(test.nums)
		if got != test.expected {
			t.Errorf("Got:\t%f\tExpected:\t%f\t", got, test.expected)
		}
	}

}
