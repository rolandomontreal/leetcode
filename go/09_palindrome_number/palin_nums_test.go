package main

import "testing"

// Example 1:

// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.
// Example 2:

// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
// Example 3:

// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

func Test1IsPalindrome(t *testing.T) {
	x := 121
	expected := true

	actual := isPalindrome(x)
	
	if (actual != expected) {
		t.Errorf("isPalindrome(%d) failed, expected: %t, actual: %t", x, expected, actual)
	}
}

func Test2IsPalindrome(t *testing.T) {
	x := -121
	expected := false

	actual := isPalindrome(x)
	
	if (actual != expected) {
		t.Errorf("isPalindrome(%d) failed, expected: %t, actual: %t", x, expected, actual)
	}
}

func Test3IsPalindrome(t *testing.T) {
	x := 10
	expected := false

	actual := isPalindrome(x)
	
	if (actual != expected) {
		t.Errorf("isPalindrome(%d) failed, expected: %t, actual: %t", x, expected, actual)
	}
}