package introduction

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		num      int
		expected []interface{}
	}{
		{
			num:      11,
			expected: []interface{}{1, 2, "fizz", 4, "buzz", "fizz", 7, 8, "fizz", "buzz", 11},
		}, {
			num:      2,
			expected: []interface{}{1, 2},
		}, {
			num:      16,
			expected: []interface{}{1, 2, "fizz", 4, "buzz", "fizz", 7, 8, "fizz", "buzz", 11, "fizz", 13, 14, "fizzbuzz", 16},
		}, {
			num:      32,
			expected: []interface{}{1, 2, "fizz", 4, "buzz", "fizz", 7, 8, "fizz", "buzz", 11, "fizz", 13, 14, "fizzbuzz", 16, 17, "fizz", 19, "buzz", "fizz", 22, 23, "fizz", "buzz", 26, "fizz", 28, 29, "fizzbuzz", 31, 32},
		},
	}

	for _, test := range tests {
		got := fizzBuzz(test.num)
		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("Num:\t%d results don't match expected\nGot:\t%v\nExpected:\t%v\n", test.num, got, test.expected)
		}
	}

}
