package recursion

import "testing"

func TestSumNumbersRecursive(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{5, 2, 9, 10},
			expected: 26,
		}, {

			nums:     []int{1, -1, 1, -1, 1, -1, 1},
			expected: 1,
		}, {
			nums:     []int{},
			expected: 0,
		}, {
			nums:     []int{1000, 0, 0, 0, 0, 0, 1},
			expected: 1001,
		}, {
			nums:     []int{700, 70, 7},
			expected: 777,
		}, {
			nums:     []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1},
			expected: -55,
		}, {
			nums:     []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expected: 0,
		},
	}

	for _, test := range tests {
		got := sumNumbersRecursive(test.nums)

		if got != test.expected {
			t.Errorf("For:\t%v\nGot:\t%d\tExpected:\t%d\t", test.nums, got, test.expected)
		}
	}
}
