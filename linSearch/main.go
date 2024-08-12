package main

import "fmt"

func main() {
	nums := []int{1, 4, 7, 9, 12}
	x := 9
	index := linearSearch(nums, x)
	fmt.Println(index)
}

func linearSearch(nums []int, x int) int {
	for i, num := range nums {
		if num == x {
			return i
		}
	}
	return -1
}
