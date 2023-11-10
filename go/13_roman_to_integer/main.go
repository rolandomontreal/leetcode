package main

import "fmt"

func main() {
	n := romanToInt("D")
	fmt.Println(n)
}

func romanToInt(s string) int {
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