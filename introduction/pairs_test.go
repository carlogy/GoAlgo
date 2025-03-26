package introduction

import (
	"reflect"
	"testing"
)

func TestPairs(t *testing.T) {
	tests := []struct {
		arr      []string
		expected [][]string
	}{
		{
			arr: []string{"a", "b", "c"},
			expected: [][]string{
				{"a", "b"},
				{"a", "c"},
				{"b", "c"},
			},
		}, {
			arr: []string{"a", "b", "c", "d"},
			expected: [][]string{
				{"a", "b"},
				{"a", "c"},
				{"a", "d"},
				{"b", "c"},
				{"b", "d"},
				{"c", "d"},
			},
		},
		{
			arr: []string{"cherry", "cranberry", "banana", "blueberry", "lime", "papaya"},
			expected: [][]string{
				{"cherry", "cranberry"},
				{"cherry", "banana"},
				{"cherry", "blueberry"},
				{"cherry", "lime"},
				{"cherry", "papaya"},
				{"cranberry", "banana"},
				{"cranberry", "blueberry"},
				{"cranberry", "lime"},
				{"cranberry", "papaya"},
				{"banana", "blueberry"},
				{"banana", "lime"},
				{"banana", "papaya"},
				{"blueberry", "lime"},
				{"blueberry", "papaya"},
				{"lime", "papaya"},
			},
		},
	}

	for _, test := range tests {
		got := pairs(test.arr)
		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("Got:\t%v\nExpected:\t%v", got, test.expected)
		}
	}
}
