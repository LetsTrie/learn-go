package main

import (
	"log"
	"sort"
)

// threeSum finds all unique triplets in the array that sum to zero.
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	result := [][]int{}

	for i := 0; i < n-2; i++ {
		v := nums[i]
		if v > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1
		for left < right {
			total := nums[left] + nums[right] + v
			if total == 0 {
				result = append(result, []int{v, nums[left], nums[right]})
				for left+1 < right && nums[left] == nums[left+1] {
					left++
				}
				left++
				for left < right-1 && nums[right] == nums[right-1] {
					right--
				}
				right--
			} else if total < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

func main() {
	testCases := [][]int{
		{-1, 0, 1, 2, -1, -4},
		{0, 0, 0},
		{},
		{1, 2, -2, -1},
		{-2, 0, 1, 1, 2},
		{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6},
		{-5, 1, 10, -1, -2, 3, 4, -3, 0},
		{-10, 5, 2, 4, -4, -5, 0, 0},
	}

	for i, testCase := range testCases {
		log.Printf("Test Case %d: Input: %v, Output: %v", i+1, testCase, threeSum(testCase))
	}
}
