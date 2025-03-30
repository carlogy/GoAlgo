package introduction

import "testing"

func TestAverage(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1},
			expected: 1,
		}, {
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			expected: 4,
		}, {
			nums:     []int{12, 12, 12},
			expected: 12,
		}, {
			nums:     nil,
			expected: 0,
		}, {
			nums:     []int{0},
			expected: 0,
		}, {
			nums:     []int{100, 200, 300, 400, 500},
			expected: 300,
		}, {
			nums:     []int{5, 10, 200, 3000, 5000},
			expected: 1643,
		}, {
			nums:     []int{12_345, 618_222, 58_832_221, 2_180_831_475, 8_663_863_102},
			expected: 2_180_831_473,
		},
	}

	for _, test := range tests {
		got := average(test.nums)

		if got != test.expected {
			t.Errorf("Got:\t%d\tExpected:\t%d\t", got, test.expected)
		}
	}
}
