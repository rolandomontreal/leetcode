package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	x := 752257
	result := isPalindrome(x)
	fmt.Println("Number is a palindrome?: ", result)
}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	isPalindrome := true
	halfL := int(math.Ceil(float64(len(s)) / 2))
	
	for i := 0; i < halfL && isPalindrome; i++ {
		if (s[i:i+1] != s[len(s)-i-1:len(s)-i]) {
			isPalindrome = false
		}
	}

	return isPalindrome
}
