package main

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	testCases := []struct{
		name string
		strs []string
		expected string
	}{
		{
			name: "Simple example test case",
			strs: []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			name: "Simple example test case 2",
			strs: []string{"dog", "racecar", "car"},
			expected: "",
		},
		{
			name: "Edge case 1",
			strs: []string{"a"},
			expected: "a",
		},
	}

	for _, tc := range testCases {
		result := longestCommonPrefix(tc.strs)
		if (result != tc.expected) {
			t.Errorf("'%s' failed, expected: '%s', got: '%s'", tc.name, tc.expected, result)
		}
	}
}