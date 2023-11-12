package main

import "fmt"

func main() {
	n := romanToInt("D")
	fmt.Println(n)
}

func romanToInt(s string) int {
	if (len(s) == 1) {
		return singleRomanToInt(s)
	}

	sum := 0;

	for i := 0; i < len(s); i++ {
		c1 := s[i:i+1]
		n1 := singleRomanToInt(c1)
		if (i == len(s) - 1) {
			sum += n1
		} else {
			c2 := s[i+1:i+2]
			n2 := singleRomanToInt(c2)
			if (n1 < n2) {
				sum = sum + (n2 - n1)
				i++
			} else {
				sum = sum + n1
			}
		}
	}

	return sum
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