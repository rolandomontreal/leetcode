package main

import (
	"fmt"
	"testing"
)

func baseTestExecutor(t *testing.T, nums []int, target int, expectedValue []int) {
	actualValue := twoSum(nums, target)

	if expectedValue[0] != actualValue[0] || expectedValue[1] != actualValue[1] {
		fmt.Println("expected: ", expectedValue)
		fmt.Println("actual: ", actualValue)
		t.Fatalf("Expected value: and actual: does not match")
	}
}

func TestTwoSumCase1(t *testing.T) {
	nums := []int{2,7,11,15}
	target := 9
	expectedValue := []int{0, 1}

	baseTestExecutor(t, nums, target, expectedValue)	
}

func TestTwoSumCase2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	expectedValue := []int{1, 2}

	baseTestExecutor(t, nums, target, expectedValue)	
}

func TestTwoSumCase3(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	expectedValue := []int{0, 1}

	baseTestExecutor(t, nums, target, expectedValue)	
}

func TestTwoSumCase4(t *testing.T) {
	nums := []int{3, 2, 3}
	target := 6
	expectedValue := []int{0, 2}

	baseTestExecutor(t, nums, target, expectedValue)	
}