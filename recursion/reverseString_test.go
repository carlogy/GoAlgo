package recursion

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{
			s:        "hello",
			expected: "olleh",
		}, {
			s:        "abcdefg",
			expected: "gfedcba",
		}, {
			s:        "stopwatch",
			expected: "hctawpots",
		}, {
			s:        "",
			expected: "",
		},
	}

	for _, test := range tests {

		got := reverseString(test.s)

		if got != test.expected {
			t.Errorf("For:\t%s\nGot:\t%sExpected:\t%s\t", test.s, got, test.expected)
		}
	}
}
