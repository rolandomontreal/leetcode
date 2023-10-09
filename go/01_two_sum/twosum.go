package main

func twoSum(nums []int, target int) []int {
	i := 0
	k := i+1
	matchFound := false

	for !matchFound && i < len(nums) - 1 {
		for !matchFound && k < len(nums) {
			num1, num2 := nums[i], nums[k]
			// fmt.Println("num1: ", num1, ", num2: ", num2, ", sum: ", num1 + num2)
			matchFound = num1 + num2 == target
			if (!matchFound) {
				k++
			}
		}
		if (!matchFound) {
			i++
			k = i + 1
		}
	}

	return []int{i, k}
}