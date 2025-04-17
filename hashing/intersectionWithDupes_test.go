package hashing

import (
	"reflect"
	"slices"
	"testing"
)

func TestIntersectionWithDupes(t *testing.T) {
	tests := []struct {
		a        []string
		b        []string
		expected []string
	}{
		{
			a:        []string{"a", "b", "c", "b"},
			b:        []string{"x", "y", "b", "b"},
			expected: []string{"b", "b"},
		}, {
			a:        []string{"q", "b", "m", "s", "s", "s"},
			b:        []string{"s", "m", "s"},
			expected: []string{"m", "s", "s"},
		}, {
			a:        []string{"p", "r", "r", "r"},
			b:        []string{"r"},
			expected: []string{"r"},
		}, {
			a:        []string{"r"},
			b:        []string{"p", "r", "r", "r"},
			expected: []string{"r"},
		}, {
			a:        []string{"t", "v", "u"},
			b:        []string{"g", "e", "d", "f"},
			expected: []string{},
		}, {
			a:        []string{"a", "a", "a", "a", "a", "a"},
			b:        []string{"a", "a", "a", "a"},
			expected: []string{"a", "a", "a", "a"},
		},
	}

	for _, test := range tests {
		got := intersectionWithDupes(test.a, test.b)

		slices.Sort(got)
		slices.Sort(test.expected)

		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("For:\t%v\t%v\nGot:\t%v\tExpected:\t%v\t", test.a, test.b, got, test.expected)
		}
	}
}
