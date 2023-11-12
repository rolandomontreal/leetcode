package main

import "testing"

func Test_romanToInt(t *testing.T) {
	romanToIntTests := []struct {
		name string
		s string
		expected int
	}{
		{ name: "One", s: "I", expected: 1, },
		{ name: "Five", s: "V", expected: 5, },
		{ name: "Ten", s: "X", expected: 10, },
		{ name: "Fifty", s: "L", expected: 50, },
		{ name: "1 hundred", s: "C", expected: 100, },
		{ name: "5 hundred", s: "D", expected: 500, },
		{ name: "1 thousand", s: "M", expected: 1000, },
	}
	for _, tc := range romanToIntTests {
		result := romanToInt(tc.s)
		if result != tc.expected {
			t.Errorf(
				"Test case '%s' failed for input '%s', expected value: %d, actual value: %d",
				tc.name,
				tc.s,
				tc.expected,
				result,
			)
		}
	}
}