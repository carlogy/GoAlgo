package introduction

import "testing"

func TestMinimumValue(t *testing.T) {

	tests := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{7, 4, 3, 100, 2343243, 343434, 1, 2, 32},
			expected: 1,
		}, {
			nums:     []int{12, 12, 12},
			expected: 12,
		}, {
			nums:     []int{1},
			expected: 1,
		}, {
			nums:     []int{1, 2, 3, 4, 5},
			expected: 1,
		}, {
			nums:     []int{5, 4, 3, 2, 1},
			expected: 1,
		}, {
			nums:     []int{100, 200, 300, 400, 500},
			expected: 100,
		}, {
			nums:     []int{500, 400, 300, 200, 100},
			expected: 100,
		}, {
			nums:     nil,
			expected: 0,
		},
	}

	for _, test := range tests {

		got := minimumValue(test.nums)

		if got != test.expected {
			t.Errorf("Got:\t%d\tExpected:\t%d", got, test.expected)
		}

	}

}
