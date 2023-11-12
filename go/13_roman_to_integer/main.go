package main

import "fmt"

func main() {
	n := romanToInt("D")
	fmt.Println(n)
}

func romanToInt(s string) int {
	sum := 0;

	for i := 0; i < len(s); i++ {
		if (i == len(s) - 1) {
			sum += romanStringToInt(s[i:])
		} else {
			substr := s[i:i+2]
			num := romanStringToInt(substr)
			if num == -1 {
				num = romanStringToInt(s[i:i+1])
			} else {
				i++
			}
			sum += num
		}
	}

	return sum
}

func romanStringToInt(s string) int {
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
	// The 6 special cases
	case s == "IV": {
	 	return 4
	}
	case s == "IX": {
		return 9
	}
	case s == "XL": {
		return 40
	}
	case s == "XC": {
		return 90
	}
	case s == "CD": {
		return 400
	}
	case s == "CM": {
		return 900
	}
	}
	return -1
}