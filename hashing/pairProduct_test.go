package hashing

import (
	"reflect"
	"testing"
)

func TestPairProduct(t *testing.T) {

	largeSlice := make([]int, 30001)

	for i := 0; i < 30001; i++ {
		largeSlice[i] = i
	}

	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{3, 2, 5, 4, 1},
			target:   8,
			expected: []int{1, 3},
		}, {
			nums:     []int{3, 2, 5, 4, 1},
			target:   10,
			expected: []int{1, 2},
		}, {
			nums:     []int{4, 7, 9, 2, 5, 1},
			target:   5,
			expected: []int{4, 5},
		}, {
			nums:     []int{4, 7, 9, 2, 5, 1},
			target:   35,
			expected: []int{1, 4},
		}, {
			nums:     []int{3, 2, 5, 4, 1},
			target:   10,
			expected: []int{1, 2},
		}, {
			nums:     []int{4, 6, 8, 2},
			target:   16,
			expected: []int{2, 3},
		},
		{
			nums:     largeSlice,
			target:   899970000,
			expected: []int{29999, 30000},
		},
	}

	for _, test := range tests {
		got := pairProduct(test.nums, test.target)

		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("For:\t%v\t%d\nGot:\t%v\tExpected\t%v", test.nums, test.target, got, test.expected)
		}
	}
}
