package main

import "fmt"

func main() {
	n := romanToInt("D")
	fmt.Println(n)
}

func romanToInt(s string) int {
	// sum := 0;

	if (len(s) == 1) {
		return singleRomanToInt(s)
	}

	// for i := 0; i < len(s) + 1; i++ {
	// 	c1 := str[i:i+1]
	// 	c2 := str[i+1:i+2]
	// 	n1 := singleRomanToInt()
	// }

	return -1
}

func singleRomanToInt(s string) int {
	switch {
	case s == "I": {
		return 1
	}
	case s == "V": {
		return 5			
	}
	case s == "X": {
		return 10
	}
	case s == "L": {
		return 50	
	}
	case s == "C": {
		return 100
	}
	case s == "D": {
		return 500	
	}
	case s == "M": {
		return 1000		
	}
	}
	return -1
}