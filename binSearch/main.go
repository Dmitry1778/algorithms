package main

import "fmt"

func main() {
	nums := []int{1, 4, 7, 9, 12}
	x := 9
	index := binarySearch(nums, x)
	fmt.Println(index)
}

func binarySearch(nums []int, x int) int {
	var firstIndex = 0
	var lastIndex = len(nums) - 1
	for firstIndex < lastIndex {
		var midIndex = (firstIndex + lastIndex) / 2
		var midValue = nums[midIndex]
		if midValue == x {
			return midIndex
		} else if midValue < x {
			firstIndex = midIndex + 1
		} else {
			lastIndex = midIndex - 1
		}
	}
	return -1
}
