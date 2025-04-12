package hashing

import (
	"reflect"
	"slices"
	"testing"
)

func TestExclusiveItems(t *testing.T) {

	largeSlice := make([]int, 60001)
	largeSliceA := make([]int, 60001)
	largeSliceB := make([]int, 60001)

	for i := 0; i < 60001; i++ {
		largeSlice[i] = i
		largeSliceA[i] = i
		largeSliceB[i] = i
	}

	tests := []struct {
		a        []int
		b        []int
		expected []int
	}{
		{
			a:        []int{4, 2, 1, 6},
			b:        []int{3, 6, 9, 2, 10},
			expected: []int{4, 1, 3, 9, 10},
		}, {
			a:        []int{2, 4, 6},
			b:        []int{4, 2},
			expected: []int{6},
		}, {
			a:        []int{4, 2, 1},
			b:        []int{1, 2, 4, 6},
			expected: []int{6},
		}, {
			a:        []int{0, 1, 2},
			b:        []int{10, 11},
			expected: []int{0, 1, 2, 10, 11},
		}, {
			a:        largeSliceA,
			b:        largeSliceB,
			expected: largeSlice,
		},
	}

	for _, test := range tests {
		got := exclusiveItems(test.a, test.b)

		slices.Sort(got)
		slices.Sort(test.expected)

		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("For:\t%v\t%v\nGot:\t%v\tExpected:\t%v\t", test.a, test.b, got, test.expected)
		}
	}

}
