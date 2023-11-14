package main

import "fmt"

func main() {
	strs := []string{"flower", "flow", "flight"}
	res := longestCommonPrefix(strs)
	fmt.Println("Result from comparison:", res)
}

func longestCommonPrefix(strs []string) string {
	l := 200

	// find shortest string
	for _, s := range strs {
		if len(s) < l {
			l = len(s)
		}
	}

	longestPrefixFound := false
	lenLongestPrefix := 0
	for i := 0; i < l && !longestPrefixFound; i++ {
		c1 := strs[0][i:i+1]
		for _, str := range strs {
			cx := str[i:i+1]
			fmt.Printf("Comparing '%s' to '%s'", c1, cx)
			if (c1 != cx) {
				fmt.Println("NO MATCH, stopping comparison")
				longestPrefixFound = true
				lenLongestPrefix = i
				break
			}
		}
	}

	fmt.Println("Result: ", longestPrefixFound, lenLongestPrefix)

	return strs[0][0:lenLongestPrefix]
}