package hashing

import (
	"reflect"
	"testing"
)

func TestPairSum(t *testing.T) {

	largeNums := make([]int, 30001)

	for i := 0; i < 30001; i++ {
		largeNums[i] = i
	}

	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{3, 2, 5, 4, 1},
			target:   8,
			expected: []int{0, 2},
		}, {
			nums:     []int{4, 7, 9, 2, 5, 1},
			target:   5,
			expected: []int{0, 5},
		}, {
			nums:     []int{4, 7, 9, 2, 5, 1},
			target:   3,
			expected: []int{3, 5},
		}, {
			nums:     []int{1, 6, 7, 2},
			target:   13,
			expected: []int{1, 2},
		}, {
			nums:     []int{9, 9},
			target:   18,
			expected: []int{0, 1},
		}, {
			nums:     []int{6, 4, 2, 8},
			target:   12,
			expected: []int{1, 3},
		},
		{
			nums:     largeNums,
			target:   59999,
			expected: []int{29999, 30000},
		},
	}

	for _, test := range tests {
		got := pairSum(test.nums, test.target)

		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("For:\t%v\nGot:\t%v\tExpected:\t%v\t", test.nums, got, test.expected)
		}
	}
}
