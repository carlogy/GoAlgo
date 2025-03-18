package introduction

import "testing"

func TestMaxValue(t *testing.T) {
	tests := []struct {
		Nums     []int
		Expected int
	}{
		{
			Nums:     []int{4, 7, 2, 8, 10, 9},
			Expected: 10,
		}, {
			Nums:     []int{10, 5, 40, 41},
			Expected: 41,
		}, {
			Nums:     []int{-5, -2, -1, -11},
			Expected: -1,
		}, {
			Nums:     []int{42},
			Expected: 42,
		}, {
			Nums:     []int{1000, 8},
			Expected: 1000,
		}, {
			Nums:     []int{1000, 8, 9000},
			Expected: 9000,
		}, {
			Nums:     []int{2, 5, 1, 1, 4},
			Expected: 5,
		},
	}

	for _, test := range tests {
		got := maxValue(test.Nums)

		if got != test.Expected {
			t.Errorf("Got:\t%d\tExpected:\t%d", got, test.Expected)
		}
	}
}
